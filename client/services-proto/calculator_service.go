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

	Average(numbers ...float64) error
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

func (service calculatorService) Average(numbers ...float64) error {

	//0. เรียกเอา stream ออกมาก่อน จาก
	streamReq, err := service.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}

	//1. ทำการ วน loop req
	for _, number := range numbers {
		//2. ปั้น request ใน message Request model ส่งไปให้ server เรืื่อยๆๆ
		req := AverageRequest{
			Number: number,
		}

		//3. client ปั้น request แบบวน loop แล้วส่งเป็น stream ไปเรื่อยๆ
		streamReq.Send(&req)
		// เพิ่ม delay ทยอยส่ง เดี๋ยวมันจะเร็วเกิน
		time.Sleep(time.Second)
	}

	//4. ตอนจบ stream แล้ว ให้ได้รับ response จาก server กลับมา
	res, err := streamReq.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("Response from server: %v \n", res)
	return nil
}
