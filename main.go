package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adityameharia/ravel/db"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	badgerPath := os.Getenv("BADGER_PATH")
	err = db.Init(badgerPath)
	defer db.Close()

	fmt.Println(c)
}
