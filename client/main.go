package main

import (
	services_proto "GO-gRPC/client/services-proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 1. connect to grpc server and port start
	cc, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	//2. init client of file_grpc.pb.go
	calculatorClient := services_proto.NewCalculatorClient(cc)

	//3. init service
	calculatorService := services_proto.NewCalculatorService(calculatorClient)

	//handle service method
	if err := calculatorService.Hello("mock data name"); err != nil {
		log.Fatalf("error service hello %s", err.Error())
	}

}
