// client.go
package main

// (1) gRPC 모듈을 import 함
import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// (3) gRPC 통신 채널을 생성함
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// (4) protoc가 생성한 *_pb2_grpc.go 화일의 stub 함수를, (3)의 채널을 사용하여 실행하여 stub를 생성함
	// 같은 폴더 안에 hello_grpc_pb2_grpc.go 가 있어서, 그냥 NewMyServiceClient 를 바로 쓸 수 있음
	stub := NewMyServiceClient(conn)

	// (5) protoc가 생성한 *_pb2.go 화일의 메세지 타입에 맞춰서, 원격 함수에 전달할 메시지를 만들고, 전달할 값을 저장함
	// hello_grpc_pb2.go 안에 정의된 MyNumber 타입을 그대로 사용
	request := &MyNumber{Value: 4}

	// (6) 원격 함수를 stub을 사용하여 호출함
	response, err := stub.MyFunction(context.Background(), request)
	if err != nil {
		log.Fatalf("MyFunction error: %v", err)
	}

	// (7) 결과를 활용하는 작업을 수행함
	log.Println("gRPC result:", response.GetValue())
}
