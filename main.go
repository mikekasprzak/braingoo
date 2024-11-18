package main

import (
	"braingoo/api"
	"braingoo/wellknown"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const DEFAULT_PORT = 3000

//import "braingoo/wellknown"
//import "braingoo/api"

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		log.Println("Server stopped.")
		os.Exit(1)
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("/.well-known/nodeinfo", wellknown.NodeInfo)
	mux.HandleFunc("/api/v1/instance", api.Instance)

	log.Print("Listening...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
