package main

import (
	"context"
	"log"

	pb "github.com/Reuben-Stephen-John/grpc-unary/calculator/proto"
)


func (s *Server) Add(ctx context.Context,in *pb.SumRequest) (*pb.SumResponse, error){
	log.Printf("Sum function was invoked with %v\n",in)
	return &pb.SumResponse{
		Sum : in.Num1 + in.Num2,
	},nil
}