package main

import (
	"flag"
	"log"
	"net"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"google.golang.org/grpc"
)

var gRPCAddr string

type Replica struct {
	NodeID      string
	gRPCAddress string
}

type server struct {
}

func Init() {
	flag.StringVar(&gRPCAddr, "gRPCAddr", "", "Address (with port) at which gRPC server is started")
}

func main() {
	listener, err := net.Listen("tcp", gRPCAddr)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	grpcServer := grpc.NewServer()
	log.Printf("Starting TCP Server on %v for gRPC\n", gRPCAddr)
	RavelClusterAdminPB.RegisterRavelClusterAdminServer(grpcServer, &server{})
	grpcServer.Serve(listener)

}
