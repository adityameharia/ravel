# Ravel Cluster Admin

The `ravel_cluster_admin` package implements the cluster admin server which is responsible for handling the http
requests from the client for R/W and gRPC server to talk to the nodes in the clusters. It has the following components:

- A gRPC server to talk to nodes
    - Handles new nodes joining in the cluster
    - It chooses and assigns a replica to the correct cluster
    - Manages the list of the all the nodes in the system and leader of each cluster.
- An HTTP server that the client can talk to for reading and writing key value pairs.

The cluster admin server has the logic to shard the keys & values using consistent hashing and then distributing them
across clusters. It keeps a track of partition id's assigned to the various members and of which keys go into which
partition

## Data Flow for POST request from client

+ The http server receives the key value from the client through a post request.
+ Then it checks whether any clusters are available.
+ The key is then passed through hash functions to get the partition on which it will be stored.
+ The value is then converted into a byte array then a gRPC request is sent to the cluster leader.
+ The key along with the data type of the value is then stored in a map for future reference.

## Data Flow for GET request from client

+ The http server receives the key from the client through a get request
+ It passes the key through the hash functions to get the cluster on which the key is stored
+ A gRPC request is sent to the cluster leader to retrieve the value as byte array.
+ The value is then converted to its original data type using the map which stored the key and its data type and then
  the response is sent to the client.

## Data flow on adding new clusters/replicas

- When you add a new replica, it is added to the cluster with the least number of replicas. The gRPC server responds
  with the required info for the replica node to join into the designated cluster.
- When a node is started as the leader to a new cluster, the gRPC server responds the new cluster info.
    - Once the leader node for the newly created cluster has properly started up, it requests the gRPC server on the
      admin for data relocation.
    - As adding a new cluster changes the configuration of the hash ring in consistent hashing, the location at which
      the existing keys are hashed at also changes. So the keys that will now get hashed to this new cluster are
      relocated there.
- This hashing is done using [consistent hashing with bounded loads](https://arxiv.org/abs/1608.01350) to ensure that
  every cluster gets an even distribution of load with minimal relocation.
- Data relocation also takes place when clusters are removed. 

```shell
protoc --go_out=. \
    --go-grpc_out=require_unimplemented_servers=false:. \
    cmd/ravel_cluster_admin/cluster_admin.proto
```