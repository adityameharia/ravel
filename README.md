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
* [Setup and Usage](#setup-and-usage)
* [Examples](#examples)
* [Killing A Ravel Instance](#killing-a-ravel-instance)
* [Uninstalling Ravel](#unistalling-ravel)
* [Documentation and Further Reading](#documentation-and-further-reading)
* [Contributing](#contributing)
* [Contact](#contact)
* [License](#license)

## Installation

Ravel has two functional components. A cluster admin server and a replica node, both of them have their separate binary
files. To setup Ravel correctly, you'll need to start one cluster admin server and many replica nodes as per requirement.

### Using Curl 

This will download the `ravel_node` and `ravel_cluster_admin` binary files and move it to `/usr/local/bin`. Make sure you
have it in your `$PATH`

```bash
curl https://raw.githubusercontent.com/nvm-sh/nvm/master/install.sh | bash
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
and `cmd/ravel_cluster_admin` respectively and move them to /usr/local/bin

You can copy them to your `$PATH` or run them from those directories

## Setup and Usage

The most simple configuration in the ravel system would be to have 2 cluster with 3 replicas each.

1. Setup the admin server

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

Once the replicas and admin are set up,we can start sending http req to our admin.

The admin exposes 3 routes for us to use:

- /put: Send a POST request to this route with attributes `key` and `val` in body to store the data in one of the clusters.
- /get: Send a POST request to this route with the `key` attribute in body to get the key-value pair 
- /delete: Send a POST request to this route with the `key` attribute in body to delete the key-value pair from the system.

---
**NOTE**
- -l=true sets up a new cluster,defaults to false
- Dont forget the storage directory as you will need it to delete the replica
- All the commands and flag can be viewed using the -h or --help flag  
- The admin will automatically decide which replica goes to which cluster
- Adding and removing clusters from the system automatically relocates all the keys in the cluster.Removing the last cluster deletes all the keys in that cluster.
---

## Examples

## Killing A Ravel Instance

Stopping a ravel instance niethers delete the data or configuration nor removes it from the system, it just replicates a crash.

In order to delete all the data and configuration and remove the instance from the system you need to kill it.

```shell
ravel_node kill -s="the storage directory you specified while starting the node"
```
Stopping the ravel_admin breaks the entire system and renders it useless.It is recommended not to stop/kill the admin unless all the replicas have been properly killed.

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
* In order to read about the data flow of the system refer to [data flow in admin](https://github.com/adityameharia/ravel/blob/main/cmd/ravel_cluster_admin/README.md) and [data flow in replica](https://github.com/adityameharia/ravel/blob/main/cmd/ravel_node/README.md)
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