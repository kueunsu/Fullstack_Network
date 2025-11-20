# Bidirectional streaming gRPC server
# Reference: https://velotio.medium.com/implementing-grpc-in-python-a-step-by-step-guide-e9733871acb0

import grpc
from concurrent import futures

import bidirectional_pb2_grpc

class BidirectionalService(bidirectional_pb2_grpc.BidirectionalServicer):

    def GetServerResponse(self, request_iterator, context):
        print('Server processing gRPC bidirectional streaming.')
        for message in request_iterator:
            yield message

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    bidirectional_pb2_grpc.add_BidirectionalServicer_to_server(BidirectionalService(), server)
    print('Starting server. Listening on port 50051.')
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()