package services_proto

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

func (server calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	result := fmt.Sprintf("Hello %s", req.Name)
	response := HelloResponse{

		Result: result,
	}
	return &response, nil
}
