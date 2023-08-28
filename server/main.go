package main

import (
	services_proto "grpcserver/services-proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//1. init server grpc
	serv := grpc.NewServer()

	// create listener for run server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	//call grpc register with service implement interface proto
	services_proto.RegisterCalculatorServer(serv, services_proto.NewCalculatorServer())

	// run server
	log.Println("gRPC server listening on port 50051")
	if err := serv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
