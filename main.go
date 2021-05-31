package main

import (
	"flag"
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

	var err error
	var r node.RavelNode
	r.Raft, r.Fsm, err = r.Open(c.joinAddr == "", c.id, c.dir, c.raftAddr)
	if err != nil {
		log.Println(err)
	}

	log.Println("sdfishfgodfnvs kdfgnsdjfgndo")

	if c.joinAddr != "" {
		log.Println("hi")
		if err := server.RequestJoin(c.id, c.joinAddr, c.raftAddr); err != nil {
			log.Fatal("pp")
			log.Fatalf("failed to join node at %s: %s", c.joinAddr, err.Error())
		}
	}
	defer server.RequestLeave(c.id, c.joinAddr)

	address := c.gRPCAddr
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Println("Server is listening")
	s := grpc.NewServer()
	RavelClusterPB.RegisterRavelClusterServer(s, &server.Server{Node: &r})
	s.Serve(lis)

}
