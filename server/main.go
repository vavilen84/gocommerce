package main

import (
	"github.com/vavilen84/gocommerce/handlers"
	"github.com/vavilen84/gocommerce/store"
	"log"
)

func main() {
	store.InitDB()
	handler := handlers.MakeHandler()
	httpServer := handlers.InitHttpServer(handler)
	log.Fatal(httpServer.ListenAndServe())
}
