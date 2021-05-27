package main

import (
	"github.com/adityameharia/ravel/datastore"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	badgerPath := os.Getenv("BADGER_PATH")
	err = datastore.Init(badgerPath)
	err = datastore.Close()
}
