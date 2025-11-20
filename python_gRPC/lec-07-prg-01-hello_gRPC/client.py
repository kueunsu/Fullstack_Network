# (1) grpc 모듈을 import 함
import grpc

# (2) protoc가 생성한 클래스를 import 함
import hello_grpc_pb2
import hello_grpc_pb2_grpc

# (3) gRPC 통신 채널을 생성함
channel = grpc.insecure_channel('localhost:50051')

# (4) protoc가 생성한 _pb2_grpc 화일의 stub 함수를, (3)의 채널을, 사용하여 실행하여 stub를 생성함
stub = hello_grpc_pb2_grpc.MyServiceStub(channel)

# (5) protoc가 생성한 _pb2 화일의 메세지 타입에 맞춰서, 원격 함수에 전달할 메시지를 만들고, 전달할 값을 저장함
request = hello_grpc_pb2.MyNumber(value=4)

# (6) 원격 함수를 stub을 사용하여 호출함
response = stub.MyFunction(request)

# (7) 결과를 활용하는 작업을 수행함 [optioanal]
print("gRPC result:", response.value)
