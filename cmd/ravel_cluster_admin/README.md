# Ravel_cluster_admin

The ravel_cluster_admin package implements the admin server which is responsible for handling the http requests from the client and gRPC server to talk to the clusters.

- It has a gRPC server to talk to nodes
  - It handles new nodes joining in the cluster
  - It manages the list of the all the nodes in our system and leaders of each clusters.

- It has a http server to talk to the client and can take string,int,float,bool and json as values.

- It has the logic to decide which cluster a new replica will join.

- It has the logic to shard the keys and values coming from the client
  -It keeps a track of partition id's assigned to the various members
  -It keeps track of which keys go into which partition
<br/><br/>

## Data Flow for POST request from client
  + The http server recieves the key value from the client through a post request.
  + Then it checks whether any clusters are available.
  + The key is then passed through hash functions to get the partition on which it will be stored.
  + The value is then converted into a byte array and a gRPC request is sent to the cluster leader.
  + The key along with the data type of the value is then stored in a map for future reference.
<br/><br/>

## Data Flow for GET request from client
  + The http server recieves the key from the client through a get request
  + It passes the key through the hash functions to get the cluster on which the key is stored
  + A gRPC request is sent to the cluster leader to retrieve the value as byte array.
  + The value is then converted to its original data type using the map which stored the key and its data type and then the response is sent to the client.

## Data flow on adding new clusters
  +
 

  

```shell
protoc --go_out=. \
    --go-grpc_out=require_unimplemented_servers=false:. \
    cmd/ravel_cluster_admin/cluster_admin.proto
```