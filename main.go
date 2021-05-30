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
	joinAddr string
	raftAddr string
}

var c Config

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&c.dir, "raftDir", dirname, "raft data directory")
	flag.StringVar(&c.gRPCAddr, "addr", "", "server listen address")
	flag.StringVar(&c.id, "id", "", "replica id")
	flag.StringVar(&c.joinAddr, "join", "", "join to already exist cluster")
	flag.StringVar(&c.raftAddr, "raftAddr", "", "Set Raft internal communication address")
}

func main() {
	flag.Parse()
	address := c.gRPCAddr
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	var r node.RavelNode
	err = r.Open(c.joinAddr == "", c.id, c.dir, c.raftAddr)
	if err != nil {
		log.Println(err)
	}

	s := grpc.NewServer()
	RavelClusterPB.RegisterRavelClusterServer(s, &server.Server{Node: &r})
	s.Serve(lis)

	if c.joinAddr != "" {
		if err := server.RequestJoin(c.joinAddr, c.raftAddr, c.id); err != nil {
			log.Fatalf("failed to join node at %s: %s", c.joinAddr, err.Error())
		}
	}
	defer server.RequestLeave(c.id, c.joinAddr)

}
