package services_proto

import (
	"context"
	"log"
)

type CalculatorService interface {
	Hello(name string) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{
		calculatorClient: calculatorClient,
	}
}

func (service calculatorService) Hello(name string) error {

	reqTogrpc := HelloRequest{

		Name: name,
	}

	responseFromgrpc, err := service.calculatorClient.Hello(context.Background(), &reqTogrpc)
	if err != nil {
		return err
	}

	log.Printf("Response from service call client to grpc server : %v", responseFromgrpc)
	return nil
}
