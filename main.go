package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adityameharia/ravel/datastore"
	"github.com/joho/godotenv"
)

var dir string
var gRPCaddr string
var id string
var join string
var raftAddr string

type Config struct {
	dir      string
	gRPCaddr string
	id       string
	join     string
	raftAddr string
}

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&dir, "raftDir", dirname, "raft data directory")
	flag.StringVar(&gRPCaddr, "addr", ":5000", "server listen address")
	flag.StringVar(&id, "id", "", "replica id")
	flag.StringVar(&join, "join", "", "join to already exist cluster")
	flag.StringVar(&raftAddr, "raftAddr", "", "Set Raft internal communication address")

}

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	badgerPath := os.Getenv("BADGER_PATH")
	err = datastore.Init(badgerPath)
	defer datastore.Close()

	c := &Config{
		dir:      dir,
		gRPCaddr: gRPCaddr,
		id:       id,
		join:     join,
		raftAddr: raftAddr,
	}
	fmt.Println(c)

	err = datastore.Close()
}
