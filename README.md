![](header.png)

Ravel is a sharded, fault-tolerant key-value store built using [BadgerDB](https://github.com/dgraph-io/badger)
and [hashicorp/raft](https://github.com/hashicorp/raft). You can shard your data across multiple clusters with multiple
replicas, the data is persisted on disk using BadgerDB for high throughput in reads and writes. Replication and
fault-tolerance is done using [Raft](https://raft.github.io/).

Ravel exposes a simple HTTP API for the user to read and write data and Ravel handles the sharding and the replication
of data across clusters.

## Table of Contents

* [Installation](#installation)
    * [Using Docker](#using-docker)
    * [From Source](#from-source)
* [Setup and Usage](#setup-and-usage)
* [Examples](#examples)
* [Documentation and Further Reading](#documentation-and-further-reading)
* [Contributing](#contributing)
* [Contact](#contact)
* [License](#license)

## Installation

### Using Docker

### From Source

Ravel has two functional components. A cluster admin server and a replica node, both of them have their separate binary
files.

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
cd ../ravel_cluster_admin
go build
```

This will build the `ravel_node` and `ravel_cluster_admin` binaries in `cmd/ravel_node`
and `cmd/ravel_cluster_admin` respectively. 

You can copy them to your `$PATH` or run them from those directories

## Setup and Usage

## Examples

## Documentation and Further Reading

* API Reference : https://pkg.go.dev/github.com/adityameharia/ravel
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