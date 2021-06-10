package main

import (
	"log"
	"net"
	"sync"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
)

var RavelClusterAdminGRPCAddr string
var RavelClusterAdminHTTPAddr string
var RavelClusterAdminBackupPath string

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

func init() {
	// dirname, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Println(err)
	// }
	RavelClusterAdminHTTPAddr = "localhost:5000"
	RavelClusterAdminGRPCAddr = "localhost:42000"
	RavelClusterAdminBackupPath = "/ravel_admin"

	// flag.StringVar(&RavelClusterAdminBackupPath, "backupPath", dirname, "Path where the Cluster Admin should persist its state on disk")
	// flag.StringVar(&RavelClusterAdminHTTPAddr, "http", "", "Address (with port) on which the HTTP server should listen")
	// flag.StringVar(&RavelClusterAdminGRPCAddr, "grpc", "", "Address (with port) on which the gRPC server should listen")
}

func main() {
	log.Println("hopelly")
	consistentHash.Init(271, 40, 1.2)
	go startAdminGRPCServer()
	go startAdminHTTPServer()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}
