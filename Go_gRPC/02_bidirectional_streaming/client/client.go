package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "bidirectional_streaming/bidipb"
)

// Python: make_message(message)
func makeMessage(message string) *pb.Message {
	return &pb.Message{Message: message}
}

// Python: generate_messages()
func generateMessages() []*pb.Message {
	return []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}
}

// Python: send_message(stub)
func sendMessage(stub pb.BidirectionalClient) error {
	// 타임아웃 컨텍스트
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 양방향 스트림 생성 (Python: stub.GetServerResponse(generate_messages()))
	stream, err := stub.GetServerResponse(ctx)
	if err != nil {
		return err
	}

	// 1) 클라이언트 → 서버로 메시지 보내기 (generate_messages + print 부분)
	go func() {
		for _, msg := range generateMessages() {
			log.Printf("[client to server] %s", msg.Message)

			if err := stream.Send(msg); err != nil {
				log.Printf("Send error: %v", err)
				return
			}
		}
		// 더 이상 보낼 메시지가 없으면 스트림 닫기
		if err := stream.CloseSend(); err != nil {
			log.Printf("CloseSend error: %v", err)
		}
	}()

	// 2) 서버 → 클라이언트 응답 읽기 (Python: for response in responses)
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Printf("Recv finished: %v", err)
			break
		}
		log.Printf("[server to client] %s", resp.Message)
	}

	return nil
}

// Python: run()
func main() {
	// Python: with grpc.insecure_channel('localhost:50051') as channel:
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Python: stub = bidirectional_pb2_grpc.BidirectionalStub(channel)
	stub := pb.NewBidirectionalClient(conn)

	// Python: send_message(stub)
	if err := sendMessage(stub); err != nil {
		log.Fatalf("sendMessage error: %v", err)
	}
}
