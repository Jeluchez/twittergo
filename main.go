package main

import (
	"log"

	"github.com/jeluchez/twittergo/bd"
	"github.com/jeluchez/twittergo/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("no conection to database")
		return
	}
	handlers.HandlerConnect()
}
