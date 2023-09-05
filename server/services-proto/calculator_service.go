package services_proto

import (
	context "context"
	"fmt"
	"io"
	"time"
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

func (server calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {
	//return หลายค่า จาก request ที่ได้แค่ 1
	for n := uint32(0); n <= req.N; n++ {
		result := fib(n)
		res := FibonacciResponse{
			Result: result,
		}
		//ส่ง response ไปให้ stream
		//1.
		stream.Send(&res)
		//2. ติด sleep ไว้หน่อย เดี๋ยวเร็วเกิน
		time.Sleep(time.Second)
	}

	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func (server calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0

	//วนรูป จาก stream ที่ client ส่่งมา
	//ใช้ for{} เพราะไม่รู้ว่า client จะส่งมามากเท่าไหร่
	for {
		//2. รับ stream
		reqStream, err := stream.Recv()
		//3. ดูว่า stream มาหมดแล้วใช่่ไหม End Of File
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		sum += reqStream.Number
		count++

	}

	//4. ตอบกลับ แค่ก้อนเดียว ก็ปั้น response
	res := AverageResponse{
		Result: sum / count,
	}

	//5. ส่งกลับหา stream
	return stream.SendAndClose(&res)
}
