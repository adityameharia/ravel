package main

import (
	"flag"
	"github.com/adityameharia/ravel/node"
	"log"
	"os"
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
	err := r.Open(true, "1", "/tmp/badger/run", c.raftAddr, "localhost:5000")
	if err != nil {
		log.Println(err)
	}
}
