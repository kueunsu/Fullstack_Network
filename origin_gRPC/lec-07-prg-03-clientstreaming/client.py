# Bidirectional streaming gRPC client
# Reference: https://velotio.medium.com/implementing-grpc-in-python-a-step-by-step-guide-e9733871acb0

import grpc

import clientstreaming_pb2_grpc
import clientstreaming_pb2

def make_message(message):
    return clientstreaming_pb2.Message(
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
    response = stub.GetServerResponse(generate_messages())
    print("[server to client] %d" % response.value)

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = clientstreaming_pb2_grpc.ClientStreamingStub(channel)
        send_message(stub)

if __name__ == '__main__':
    run()