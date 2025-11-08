from generated import customer_pb2_grpc, customer_pb2
import grpc

class HelloWorld(customer_pb2_grpc.HelloWorldServicer):
    async def SayHello(self, request, context):
        return customer_pb2.HelloReply(message=f"Hello, {request.name}!")