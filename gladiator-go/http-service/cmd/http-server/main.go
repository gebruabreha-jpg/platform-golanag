package main

import (
	"flag"
	"log"
	"http-service/internal/server"
)

func main() {
	port := flag.Int("port", 80, "Server port")
	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	srv := server.NewServer(*port, *debug)
	log.Printf("Starting HTTP server on :%d (debug=%v)", *port, *debug)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
