from __future__ import print_function

import grpc
import sys

import message_pb2
import message_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = message_pb2_grpc.URLStub(channel)
    response = stub.register(message_pb2.UrlRequest(url=sys.argv[1]))
    print("URL VALID:  " + str(response.valid))


if __name__ == '__main__':
    run()
