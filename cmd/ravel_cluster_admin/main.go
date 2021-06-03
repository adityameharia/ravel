package main

import (
	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"log"
	"net"
)

const RavelClusterAdminGRPCAddr = "localhost:42000"

func main() {
	log.Println("Starting Ravel Cluster Admin gRPC Server on ", RavelClusterAdminGRPCAddr)
	listener, err := net.Listen("tcp", RavelClusterAdminGRPCAddr)
	if err != nil {
		log.Fatalf("Error in starting Ravel Cluster Admin TCP Listener: %v\n", err)
	}

	clusterAdminServer := NewClusterAdminGRPCServer()
	RavelClusterAdminPB.RegisterRavelClusterAdminServer(clusterAdminServer.Server, clusterAdminServer)
	err = clusterAdminServer.Server.Serve(listener)
}
