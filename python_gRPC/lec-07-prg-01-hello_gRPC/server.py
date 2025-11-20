# (1) grpc/futures 모듈을 import함
import grpc
from concurrent import futures

# (2) protoc가 생성한 클래스를 import 함
import hello_grpc_pb2 # message class
import hello_grpc_pb2_grpc # client & server class

# (3) 원격 호출될 함수들을 import 함
import hello_grpc # original remotely called functions

# (4) protoc가 생성한 Servicer 클래스를 base class로 해서 원격 호출될 함수들을 멤버로 갖는 서버 클래스를 생성함
class MyServiceServicer(hello_grpc_pb2_grpc.MyServiceServicer):

# (5) 서버 클래스에 원격 호출될 함수에 대한 rpc 함수를 작성함 

    # (5.1) proto 화일내 정의한 rpc 함수 이름에 대응하는 멤버 함수를 작성함
    def MyFunction(self, request, context):
        # (5.2) proto 화일내 message 이름과 동일한 message class를 생성하여 응답 전달 용도로 사용함
        response = hello_grpc_pb2.MyNumber() 
        # (5.3) proto 화일내 message 이름과 동일한 message class의 변수에 원격 함수의 수행 결과를 저장함
        # 앞서 (3)의 원격 호출할 함수에게 client로 부터 받은 입력 파라메타를 전달하고 결과를 가져옴
        response.value = hello_grpc.my_func(request.value)
        # (5.4) 원격 함수 호출 결과를 client에게 돌려줌
        return response

# (6) grpc.server를 생성함
# ThreadPoolExecutor : a pool of threads to execute calls asynchronously
# futures : a high-level interface for asynchronously executing callables
server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

# (7) add_CalculatorServicer_to_server()를 사용해서, grpc.server에 (4)의 Servicer를 추가함
hello_grpc_pb2_grpc.add_MyServiceServicer_to_server(MyServiceServicer(), server)

# (8) grpc.server의 통신 포트를 열고, start()로 서버를 실행함
print('Starting server. Listening on port 50051.')
server.add_insecure_port('[::]:50051')
server.start()

# (9) grpc.server가 유지되도록 프로그램 실행을 유지함
try:
    server.wait_for_termination()
except KeyboardInterrupt:
    server.stop(0)

    