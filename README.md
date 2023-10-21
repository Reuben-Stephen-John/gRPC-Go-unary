# gRPC Go Unary Functions 
Unary operations in gRPC refer to a simple request-response interaction between a client and a server. In this type of operation, the client sends a single request to the server and waits for a single response. 
## Overview
This repository contains an implementation of a unary gRPC service using Go. The service is designed to be simple and straightforward to understand, making it a great starting point for anyone interested in learning about gRPC and Go.

## Prerequisites
- Go version 1.21.3
- Protocol Buffers v3
- gRPC Go Package

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/reuben-stephen-john/grpc-go-unary.git
2. Navigate to the project directory:
   ```sh
   cd grpc-go-unary
3. Check out the quickstart guide to install gRPC for your system https://grpc.io/docs/languages/go/quickstart/
## Usage
1. Start the server:
   ```sh
    go run server/server.go
  
2. In another terminal, start the client:
    ```sh
    go run client/client.go
