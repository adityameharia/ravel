package main

import (
	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"log"
	"net"
	"sync"
)

const RavelClusterAdminGRPCAddr = "localhost:42000"
const RavelClusterAdminHTTPAddr = "localhost:3142"
const RavelClusterAdminBackupPath = "/tmp/badger/cluster_admin"

var consistentHash RavelConsistentHash
var clusterAdminGRPCServer *ClusterAdminGRPCServer
var clusterAdminHTTPServer *ClusterAdminHTTPServer

func startAdminGRPCServer() {
	log.Println("Starting Ravel Cluster Admin gRPC Server on", RavelClusterAdminGRPCAddr)
	listener, err := net.Listen("tcp", RavelClusterAdminGRPCAddr)
	if err != nil {
		log.Fatalf("Error in starting Ravel Cluster Admin TCP Listener: %v\n", err)
	}

	clusterAdminGRPCServer = NewClusterAdminGRPCServer()
	RavelClusterAdminPB.RegisterRavelClusterAdminServer(clusterAdminGRPCServer.Server, clusterAdminGRPCServer)
	err = clusterAdminGRPCServer.Server.Serve(listener)
}

func startAdminHTTPServer() {
	log.Println("Starting Ravel Cluster Admin HTTP Server on", RavelClusterAdminHTTPAddr)
	clusterAdminHTTPServer = NewClusterAdminHTTPServer()
	clusterAdminHTTPServer.Router.Run(RavelClusterAdminHTTPAddr)
}

func main() {
	consistentHash.Init(271, 40, 1.2)
	go startAdminGRPCServer()
	go startAdminHTTPServer()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
