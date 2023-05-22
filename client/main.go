package main

import (
	"github.com/ozzyozbourne/unaryone/pbout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const addr string = "localhost:8080"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to conect %s\n", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Error on closing the connection %s\n", err)
		}
	}()

	c := pbout.NewSaveInGRPCClient(conn)
	saveInGrpc(c)
	readFromGrpc(c)
}
