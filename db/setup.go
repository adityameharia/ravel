package db

import "log"

var KeyValueDB RavelDatabase
var LogDB RavelDatabase

func Setup(kvPath string, logPath string) {
	err := KeyValueDB.Init(kvPath)
	if err != nil {
		log.Fatal("Error in setting up KeyValueDB", err)
	}

	err = LogDB.Init(logPath)
	if err != nil {
		log.Fatal("Error in setting up LogDB", err)
	}
}
