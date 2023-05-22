package main

import (
	"context"
	"github.com/ozzyozbourne/unaryone/excelize"
	"github.com/ozzyozbourne/unaryone/pbout"
	"log"
)

func (s *Server) PersistValues(ctx context.Context, in *pbout.XlsxValues) (*pbout.XlsxValuesResponse, error) {
	log.Printf("PersistValues function invoked with request %v\n", in)
	if err := excelize.WriteToFile(in); err != nil {
		log.Printf("Error 400 in writing xlsx to disk %s\n", err)
		return &pbout.XlsxValuesResponse{Res: 400}, err
	} else {
		log.Printf("Success 200 Error nil\n")
		return &pbout.XlsxValuesResponse{Res: 200}, nil
	}
}

func (s *Server) GetXlsxValues(ctx context.Context, in *pbout.GetValRequest) (*pbout.XlsxValues, error) {
	log.Printf("GetXlsxValues function invoked with request %v\n", in)
	log.Printf("Function not implemented yet\n")
	return &pbout.XlsxValues{}, nil
}
