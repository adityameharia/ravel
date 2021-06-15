# Ravel Node

The `ravel_node` package implements the replicas which eventually form a cluster. It is responsible for storing the
Key-Value pairs in a persistent manner and communicates with other servers in the cluster for fault tolerance.

Since we use the [Raft consensus algorithm](https://raft.github.io/) for fault tolerance each replica server can be the
leader or follower i.e. if a leader is down the follower can become the leader and handle requests from the admin
servers. This server communicates only with the admin servers and other replicas in its cluster.

- It has a gRPC server to talk the other replicas and the admin server.
    - If it's a follower
        + It coordinates with the leader for log replication and snapshots.
    - If it's a leader
        + Send heartbeats to all its followers to prevent re-election.
        + Handles post, get and delete request from the admin server.
        + Handles new replicas joining and leaving the cluster.

- It is also a gRPC client
    + It sends heartbeats to all its followers to prevent re-election.
    + It is responsible for sending out append entries to its followers.
    + It informs the admin in case there is a change in leadership.

## Data flow for a POST/DELETE request from admin

- The leader receives the request from the admin and sends out append entries to its followers.
- On receiving a response from majority of its followers it sends out an "apply" message to its followers to actually
  execute the request.
- If the majority of the followers respond without any errors, the leader sends an "ok" message to the admin

## Data flow for a GET request from admin

- A get request is pretty simple to handle and doesn't involve any followers
- The leader receives the request from admin and find the key in the disc and responds with the Key-Value.

## Data flow for adding new replicas

- When a new replica is bootstrapped, it sends a request to the admin,which adds the server to its list of all servers in
  the system and responds with the cluster the replica has to join and the address of its leader node.
- The replica then sends a join request to the leader node.
- On successfully joining the cluster the leader sends its latest snapshots and log entries to the node for replication.

## Data flow for removing a replica

- When a replica is taken down, it sends a request to the admin to get the gRPC address of its current leader.
- The admin removes it from the list of all servers in the system and responds with the address of the leader.
- The replica then sends a leave request to the leader for successful removal.
- If the replica suffers a failure it abruptly leaves the cluster and can easily join back the cluster


