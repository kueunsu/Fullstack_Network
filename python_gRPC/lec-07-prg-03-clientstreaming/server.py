# Bidirectional streaming gRPC server
# Reference: https://velotio.medium.com/implementing-grpc-in-python-a-step-by-step-guide-e9733871acb0

import grpc
from concurrent import futures

import clientstreaming_pb2
import clientstreaming_pb2_grpc

class ClientStreamingServicer(clientstreaming_pb2_grpc.ClientStreamingServicer):

    def GetServerResponse(self, request_iterator, context):
        print('Server processing gRPC client-streaming.')
        count = 0
        for message in request_iterator:
            count += 1
        return clientstreaming_pb2.Number(value=count)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    clientstreaming_pb2_grpc.add_ClientStreamingServicer_to_server(ClientStreamingServicer(), server)
    print('Starting server. Listening on port 50051.')
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()