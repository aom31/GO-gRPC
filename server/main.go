package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	servers := grpc.NewServer()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	//1.register for use gRPC service
	services.RegisterCalculatorServer(servers, services.NewCalculatorServer())

	fmt.Println("gRPC server listening on port :50051")
	//2.run server
	err = servers.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
