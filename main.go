package main

import (
	"net/http"

	"github.com/thudsonbu/srvr/business"
	"github.com/thudsonbu/srvr/controller"
	"github.com/thudsonbu/srvr/log"
	"github.com/thudsonbu/srvr/store"
)

func main() {
	logger := business.LoggerAdapter(log.Log)
	store := store.NewKeyValStore()
	logic := business.NewBusinessLogic(logger, store)
	control := controller.NewController(logger, logic)

	http.HandleFunc("/hello", control.SayHello)

	port := ":7777"

	log.Log("Starting server on port: " + port)

	http.ListenAndServe(port, nil)
}
