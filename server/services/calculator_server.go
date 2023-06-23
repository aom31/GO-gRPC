package services

import (
	context "context"
	"fmt"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// implement service on server with func on interface in file _grpc.pb.go
func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %s", req.Name)
	res := HelloResponse{
		Result: result,
	}

	return &res, nil
}
