![](header.png)

Ravel is a sharded, fault-tolerant key-value store built using [BadgerDB](https://github.com/dgraph-io/badger)
and [hashicorp/raft](https://github.com/hashicorp/raft). You can shard your data across multiple clusters with multiple
replicas, the data is persisted on disk using BadgerDB for high throughput in reads and writes. Replication and
fault-tolerance is done using [Raft](https://raft.github.io/).

Ravel exposes a simple HTTP API for the user to read and write data and Ravel handles the sharding and the replication
of data across clusters.

## Table of Contents

* [Installation](#installation)
    * [Using Curl](#using-curl)
    * [From Source](#from-source)
* [Usage](#usage)
* [Setup a Cluster](#setup-a-cluster)
* [Reading and Writing Data](#reading-and-writing-data)
* [Killing A Ravel Instance](#killing-a-ravel-instance)
* [Uninstalling Ravel](#unistalling-ravel)
* [Documentation and Further Reading](#documentation-and-further-reading)
* [Contributing](#contributing)
* [Contact](#contact)
* [License](#license)

## Installation

Ravel has two functional components. A cluster admin server and a replica node, both of them have their separate binary
files. To setup Ravel correctly, you'll need to start one cluster admin server and many replica nodes as per
requirement.

### Using `curl`

This will download the `ravel_node` and `ravel_cluster_admin` binary files and move it to `/usr/local/bin`, make sure
you have it in your `$PATH`

```sh
curl https://raw.githubusercontent.com/adityameharia/ravel/main/install.sh | bash
```

### From Source

- `cmd/ravel_node` directory has the implementation of `ravel_node` which is the replica node
- `cmd/ravel_cluster_admin` directory has the implementation of `ravel_cluster_admin` which is the cluster admin server

1. Clone this repository

```shell
git clone https://github.com/adityameharia/ravel
cd ravel
git checkout master
```

2. Build `ravel_node` and `ravel_cluster_admin`

```shell
cd cmd/ravel_node
go build 
sudo mv ./ravel_node /usr/local/bin
cd ../ravel_cluster_admin
go build
sudo mv ./ravel_cluster_admin /usr/local/bin
```

This will build the `ravel_node` and `ravel_cluster_admin` binaries in `cmd/ravel_node`
and `cmd/ravel_cluster_admin` respectively and move them to `/usr/local/bin`

## Usage

Usage info for `ravel_cluster_admin`

```shell
$ ravel_cluster_admin --help
NAME:
   Ravel Cluster Admin - Start a Ravel Cluster Admin server

USAGE:
   ravel_cluster_admin [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --http value        Address (with port) on which the HTTP server should listen
   --grpc value        Address (with port) on which the gRPC server should listen
   --backupPath value  Path where the Cluster Admin should persist its state on disk
   --help, -h          show help
```

Usage info for `ravel_node`

```shell
$ ravel_node --help
NAME:
   Ravel Replica - Manage a Ravel replica server

USAGE:
   ravel_node [global options] command [command options] [arguments...]

COMMANDS:
   start    Starts a replica server
   kill     Removes and deletes all the data in the cluster
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

Usage info for the `start` command in `ravel_node`. Use this command to start a replica server.

```shell
$ ravel_node start --help
NAME:
   ravel_node start - Starts a replica server

USAGE:
   ravel_node start [command options] [arguments...]

OPTIONS:
   --storagedir value, -s value    Storage Dir (default: "~/ravel_replica")
   --grpcaddr value, -g value      GRPC Addr of this replica (default: "localhost:50000")
   --raftaddr value, -r value      Raft Internal address for this replica (default: "localhost:60000")
   --adminrpcaddr value, -a value  GRPC address of the cluster admin (default: "localhost:42000")
   --yaml value, -y value          yaml file containing the config
   --leader, -l                    Register this node as a new leader or not (default: false)
   --help, -h                      show help (default: false)
```

## Setup a Cluster

Executing the following instructions will setup a sample Ravel instance. The most simple configuration of a Ravel
instance would consist of 2 clusters with 3 replicas each.

The key value pairs will be sharded across the two clusters and replicated thrice on each cluster. The admin will
automatically decide which replica goes to which cluster. Adding and removing clusters from the system automatically
relocates all the keys in that cluster to some other one. Deleting the last standing cluster deletes all the keys in the
instance.

1. Setup the cluster admin server

```shell
sudo ravel_cluster_admin --http="localhost:5000" --grpc="localhost:42000" --backupPath="~/ravel_admin"
```

2. Setting up the cluster leaders

```shell
sudo ravel_node start -s="/tmp/ravel_leader1" -l=true -r="localhost:60000" -g="localhost:50000" -a="localhost:42000"
sudo ravel_node start -s="/tmp/ravel_leader2" -l=true -r="localhost:60001" -g="localhost:50001" -a="localhost:42000"
```

3. Setting up the replicas

```shell
sudo ravel_node start -s="/tmp/ravel_replica1" -r="localhost:60002" -g="localhost:50002" -a="localhost:42000"
sudo ravel_node start -s="/tmp/ravel_replica2" -r="localhost:60003" -g="localhost:50003" -a="localhost:42000"
sudo ravel_node start -s="/tmp/ravel_replica3" -r="localhost:60004" -g="localhost:50004" -a="localhost:42000"
sudo ravel_node start -s="/tmp/ravel_replica4" -r="localhost:60005" -g="localhost:50005" -a="localhost:42000"
```

**NOTE**

- `-l=true` sets up a new cluster,defaults to false
- Dont forget the storage directory as you will need it to delete the replica
- All the commands and flag can be viewed using the -h or --help flag

## Reading and Writing Data

Once the replicas and admin are set up, we can start sending HTTP requests to our cluster admin server to read, write
and delete key value pairs.

The cluster admin server exposes 3 HTTP routes:

- URL: `/put`
  - Method: `POST`
  - Description: Store a key value pair in the system 
  - Request Body: `{"key": "<your_key_here>", "val":<your_value_here>}`
    - `key = [string]`
    - `val = [string | float | JSON Object]`
  - Success Response: `200` with body `{"msg": "ok"}`

- URL: `/get`
  - Method:`POST`
  - Description: Get a key value pair from the system
  - Request Body: `{"key": "<your_key_here>"`
    - `key = [string]`
  - Success Response: `200` with body `{"key": <key>, "val":<value>}`

- URL: `/get`
  - Method:`POST`
  - Description: Delete a key value pair from the system
  - Request Body: `{"key": "<your_key_here>"`
    - `key = [string]`
  - Success Response: `200` with body `{"msg": "ok"}`
  
### Sample Requests

* Sample `/put` requests
```json
{
  "key": "the_answer",
  "value": 42
}
```
```json
{
  "key": "dogegod",
  "value": "Elon Musk"
}
```
```json
{
  "key": "hello_friend",
  "value": {
    "elliot": "Rami Malek",
    "darlene": "Carly Chaikin"
  }
}
```

* Sample `/get` request
```json
{
  "key": "dogegod"
}
```

* Sample `/delete` request
```json
{
  "key": "dogegod"
}
```

## Killing A Ravel Instance

Stopping a ravel instance niethers delete the data or configuration nor removes it from the system, it just replicates a
crash.

In order to delete all the data and configuration and remove the instance from the system you need to kill it.

```shell
ravel_node kill -s="the storage directory you specified while starting the node"
```

Stopping the ravel_admin breaks the entire system and renders it useless.It is recommended not to stop/kill the admin
unless all the replicas have been properly killed.

In order to kill the admin just delete its storage directory.

```shell
sudo rm -rf "path to storage directory"
```

## Uninstalling Ravel

Ravel can be uninstalled bye deleting the binaries from /usr/local/bin

```shell
sudo rm /usr/local/bin/ravel_node
sudo rm /usr/local/bin/ravel_cluster_admin
```

## Documentation and Further Reading

* API Reference : https://pkg.go.dev/github.com/adityameharia/ravel
* In order to read about the data flow of the system refer
  to [data flow in admin](https://github.com/adityameharia/ravel/blob/main/cmd/ravel_cluster_admin/README.md)
  and [data flow in replica](https://github.com/adityameharia/ravel/blob/main/cmd/ravel_node/README.md)
* Each package also has its own readme explainin what it does and how it does it.
* Other blogs and resources
    * https://raft.github.io/
    * https://blog.dgraph.io/post/badger/
    * [MIT 6.824: Distributed Systems](https://youtube.com/playlist?list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB)

## Contributing

If you're interested in contributing to Ravel, check out [CONTRIBUTING.md](CONTRIBUTING.md)

## Contact

Reach out to the authors with questions, concerns or ideas about improvement.

* adityameharia14@gmail.com
* junaidrahim5a@gmail.com

## License

Copyright (c) **Aditya Meharia** and **Junaid Rahim**. All rights reserved. Released under the [MIT](LICENSE) License