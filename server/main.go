package main

import (
	"github.com/ozzyozbourne/unaryone/pbout"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pbout.SaveInGRPCServer
}

const addr string = "0.0.0.0:8080"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error failed to listen port %s [ERROR] -> %s\n", addr, err)
	}

	log.Printf("Listening on port %s\n", addr)
	s := grpc.NewServer()
	pbout.RegisterSaveInGRPCServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s\n", err)
	}

}
