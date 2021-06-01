package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/adityameharia/ravel/RavelClusterPB"
	"github.com/adityameharia/ravel/node"
	"github.com/adityameharia/ravel/server"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Config struct {
	dir      string
	gRPCAddr string
	id       string
	joinAddr string
	raftAddr string
}

var adminAddress string

var c Config

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	id := uuid.New().String()

	flag.StringVar(&c.dir, "raftDir", dirname, "raft data directory")
	flag.StringVar(&c.gRPCAddr, "addr", "", "server listen address")
	flag.StringVar(&c.id, "id", id, "replica id")
	flag.StringVar(&c.joinAddr, "join", "", "join to already exist cluster")
	flag.StringVar(&c.raftAddr, "raftAddr", "", "Set Raft internal communication address")
	flag.StringVar(&adminAddress, "admin", "", "get the address of the admin cluster")
}

func main() {
	flag.Parse()

	var err error
	var r node.RavelNode
	r.Raft, r.Fsm, err = r.Open(c.joinAddr == "", c.id, c.dir, c.raftAddr)
	if err != nil {
		log.Println(err)
	}

	if c.joinAddr != "" {

		if err := server.RequestJoin(c.id, c.joinAddr, c.raftAddr); err != nil {

			log.Fatalf("failed to join node at %s: %s", c.joinAddr, err.Error())
		}
	}
	onSigInt()
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

func onSigInt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		var temp string
		if c.joinAddr == "" {
			temp = c.gRPCAddr
		} else {
			temp = c.joinAddr
		}
		err := server.RequestLeave(c.id, temp)
		if err != nil {
			log.Println(err)
		}
		os.Exit(1)
	}()
}
