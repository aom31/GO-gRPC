package services_proto

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonnacci(n uint32) error
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

func (service calculatorService) Fibonnacci(n uint32) error {
	//1. ปั้น request ใส่ message
	req := FibonacciRequest{
		N: n,
	}

	//note เราสามารถกำหนด timeout ได้ ถ้า stream มันวิ่งมาหาเรานานเกินไป
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	//2. เรียก method เพื่อส่ง req เข้าไป และรับ response จาก server มา
	streamResponse, err := service.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}

	//3. วนloop เพื่อรับ response แบบ stream
	fmt.Println("Service : Fibonacci")
	fmt.Printf("Request : %v \n", req.N)
	for {

		//4.รับ stream
		res, err := streamResponse.Recv()
		// note ถ้าทางฝั่ง server stream เสร็จแล้ว จะให้ err == เพื่อดูว่า stream หมดแล้วรึยัง
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Printf("Response from stream : %v \n", res.Result)
	}

	return nil
}
