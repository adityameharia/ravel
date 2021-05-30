package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/adityameharia/ravel/RavelClusterPB"
	"github.com/adityameharia/ravel/node"
	"github.com/adityameharia/ravel/server"
	"google.golang.org/grpc"
)

type Config struct {
	dir      string
	gRPCAddr string
	id       string
	join     string
	raftAddr string
}

var c Config

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&c.dir, "raftDir", dirname, "raft data directory")
	flag.StringVar(&c.gRPCAddr, "addr", ":5000", "server listen address")
	flag.StringVar(&c.id, "id", "", "replica id")
	flag.StringVar(&c.join, "join", "", "join to already exist cluster")
	flag.StringVar(&c.raftAddr, "raftAddr", "", "Set Raft internal communication address")
}

func main() {
	flag.Parse()
	var r node.RavelNode
	err := r.Open(true, "1", "/tmp/badger/run", "/tmp/badger/run/snapshot", "localhost:5000")
	if err != nil {
		log.Println(err)
	}
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	RavelClusterPB.RegisterRavelClusterServer(s, &server.Server{Node: &r})
	s.Serve(lis)
}
