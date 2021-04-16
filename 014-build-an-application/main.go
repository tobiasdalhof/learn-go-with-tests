package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	file, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("error opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatalf("error creating file system player store, %v ", err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
