// server.go
package main

// (1) grpc/futures 모듈을 import함
//    → Go에서는 gRPC + net + context를 import
import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// (4) protoc가 생성한 Servicer 인터페이스를 구현하는
//     서버용 struct를 정의함 (파이썬의 MyServiceServicer에 해당)
type myServiceServer struct {
	UnimplementedMyServiceServer // 생성된 코드에 있는 기본 구현(옵션)
}

// (5) 서버 클래스에 원격 호출될 함수에 대한 rpc 함수를 작성함
// (5.1) proto 화일 내 정의한 rpc 함수 이름에 대응하는 메서드를 작성함
func (s *myServiceServer) MyFunction(ctx context.Context, req *MyNumber) (*MyNumber, error) {
	// (5.2) proto 화일 내 message 이름과 동일한 message struct를 생성하여 응답 전달 용도로 사용함
	resp := &MyNumber{}

	// (5.3) 원격 호출할 함수(myFunc)에 client로부터 받은 입력 파라메터를 전달하고 결과를 가져옴
	//       파이썬의 hello_grpc.my_func(request.value)에 해당
	resp.Value = myFunc(req.Value)

	// (5.4) 원격 함수 호출 결과를 client에게 돌려줌
	return resp, nil
}

// (3) 원격 호출될 함수들을 import 함
//     → Go에서는 같은 패키지 안에 "원래 로직" 함수를 두면 됨.
//     아래는 예시 구현. 실제 로직은 네가 원하는 대로 바꾸면 됨.
func myFunc(x int32) int32 {
	// TODO: 파이썬 hello_grpc.my_func와 동일한 동작으로 수정하기
	return x * x // 예시: 제곱
}

func main() {
	// (6) grpc.Server를 생성함
	grpcServer := grpc.NewServer()

	// (7) RegisterMyServiceServer()를 사용해서, grpc.Server에 (4)의 Servicer를 추가함
	RegisterMyServiceServer(grpcServer, &myServiceServer{})

	// (8) 통신 포트를 열고, 서버를 실행함
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Starting server. Listening on port 50051.")

	// (9) grpc.Server가 유지되도록 프로그램 실행을 유지함
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
