package main

import (
	"context"
	"log"

	pb "github.com/Reuben-Stephen-John/grpc-unary/calculator/proto"
)

func doAdd(c pb.CalculatorServiceClient){
	log.Println("doAdd was invoked")
	res, err:= c.Add(context.Background(),&pb.SumRequest{
		Num1: 1,
		Num2: 1,
	})

	if err!=nil{
		log.Fatalf("Could not add: %v",err)
	}

	log.Printf("Sum: %d\n",res.Sum)
}