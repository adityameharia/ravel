#!/bin/bash

go build
sudo rm -rf /tmp/badger
./ravel_node -nodeID=1 -storageDir=/tmp/badger/run4 -gRPCAddr="localhost:50000" -raftAddr="localhost:60000"
