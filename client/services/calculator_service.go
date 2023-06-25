package services

import (
	"context"
	"fmt"
)

type ICalculatorService interface {
	Hello(name string) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) ICalculatorService {
	return calculatorService{calculatorClient}

}

func (service calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}

	res, err := service.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service : Hello \n")
	fmt.Printf("Request: %++v\n ", req.Name)
	fmt.Printf("Response: %++v \n", res.Result)

	return nil
}
