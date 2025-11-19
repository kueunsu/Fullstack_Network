# Bidirectional streaming gRPC client
# Reference: https://velotio.medium.com/implementing-grpc-in-python-a-step-by-step-guide-e9733871acb0

import grpc

import bidirectional_pb2_grpc
import bidirectional_pb2

def make_message(message):
    return bidirectional_pb2.Message(
        message=message
    )

def generate_messages():
    messages = [
        make_message("message #1"),
        make_message("message #2"),
        make_message("message #3"),
        make_message("message #4"),
        make_message("message #5"),
    ]
    for msg in messages:
        print("[client to server] %s" % msg.message)
        yield msg

def send_message(stub):
    responses = stub.GetServerResponse(generate_messages())
    for response in responses:
        print("[server to client] %s" % response.message)

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = bidirectional_pb2_grpc.BidirectionalStub(channel)
        send_message(stub)

if __name__ == '__main__':
    run()