package main

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/urfave/cli"
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

func main() {
	app := cli.NewApp()
	app.Name = "Ravel Cluster Admin"
	app.Usage = "Start a Ravel Cluster Admin server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "http",
			Required:    true,
			Usage:       "Address (with port) on which the HTTP server should listen",
			Destination: &RavelClusterAdminHTTPAddr,
		},
		cli.StringFlag{
			Name:        "grpc",
			Required:    true,
			Usage:       "Address (with port) on which the gRPC server should listen",
			Destination: &RavelClusterAdminGRPCAddr,
		},
		cli.StringFlag{
			Name:        "backupPath",
			Required:    true,
			Usage:       "Path where the Cluster Admin should persist its state on disk",
			Destination: &RavelClusterAdminBackupPath,
		},
	}

	app.Action = func(c *cli.Context) {
		consistentHash.Init(271, 40, 1.2)
		go startAdminGRPCServer()
		go startAdminHTTPServer()

		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
