package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "bidirectional_streaming/bidipb"
)

// Python의 BidirectionalService(bidirectional_pb2_grpc.BidirectionalServicer)에 해당
type BidirectionalService struct {
	pb.UnimplementedBidirectionalServer
}

// Python의 GetServerResponse: 받은 메시지를 그대로 다시 클라이언트로 보내줌
func (s *BidirectionalService) GetServerResponse(
	stream pb.Bidirectional_GetServerResponseServer,
) error {
	log.Println("Server processing gRPC bidirectional streaming.")

	for {
		// 클라이언트 → 서버로부터 메시지 받기
		msg, err := stream.Recv()
		if err != nil {
			// 스트림 닫힘(io.EOF 포함) → 종료
			log.Printf("stream.Recv finished: %v", err)
			return nil
		}

		log.Printf("[client to server] %s", msg.Message)

		// 받은 메시지를 그대로 다시 클라이언트로 전송
		if err := stream.Send(msg); err != nil {
			log.Printf("stream.Send error: %v", err)
			return err
		}
	}
}

func main() {
	// gRPC 서버 생성
	server := grpc.NewServer()

	// 서비스 등록 (Python: add_BidirectionalServicer_to_server)
	pb.RegisterBidirectionalServer(server, &BidirectionalService{})

	// 50051 포트 리슨
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
