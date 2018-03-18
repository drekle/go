#  gRPC python example

## Setup
python -m pip install grpcio
python -m pip install grpcio-tools

## Generate API
python -m grpc_tools.protoc -I../api --python_out=. --grpc_python_out=. ../api/message.proto