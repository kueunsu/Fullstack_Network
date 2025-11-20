# Bidirectional streaming gRPC client
# Reference: https://velotio.medium.com/implementing-grpc-in-python-a-step-by-step-guide-e9733871acb0

import grpc

import serverstreaming_pb2
import serverstreaming_pb2_grpc

def recv_message(stub):
    request = serverstreaming_pb2.Number(value=5)
    responses = stub.GetServerResponse(request)
    for response in responses:
        print("[server to client] %s" % response.message)

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = serverstreaming_pb2_grpc.ServerStreamingStub(channel)
        recv_message(stub)

if __name__ == '__main__':
    run()