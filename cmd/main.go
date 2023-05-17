package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dgyurics/p2p/internal/config"
	"github.com/dgyurics/p2p/internal/server"
)

var (
	port = flag.Int("port", config.Port, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	config.Port = *port

	s := server.NewGRPCServer()
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
