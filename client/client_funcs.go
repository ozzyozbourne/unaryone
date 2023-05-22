package main

import (
	"context"
	"github.com/ozzyozbourne/unaryone/pbout"
	"github.com/ozzyozbourne/unaryone/utils"
	"log"
)

func saveInGrpc(c pbout.SaveInGRPCClient) {
	log.Printf("Save in grpc function invoked\n")
	res, err := c.PersistValues(context.Background(), utils.GetXlsxValue())
	if err != nil {
		log.Fatalf("Unable to persist the request in server as xlsx %s\n", err)
	}
	log.Printf("Response -> %v\n", res)
}

func readFromGrpc(c pbout.SaveInGRPCClient) {
	log.Printf("Read from grpc function invoked\n")
	res, err := c.GetXlsxValues(context.Background(), &pbout.GetValRequest{Res: "Schools out"})
	if err != nil {
		log.Fatalf("Unable to read from the server %s\n", err)
	}
	log.Printf("Response -> %v\n", res)
}
