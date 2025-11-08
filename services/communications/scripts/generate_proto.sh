#!/bin/bash

PROTO_DIR=proto
OUT_DIR=src/generated

mkdir -p $OUT_DIR
python -m grpc_tools.protoc \
    -I$PROTO_DIR \
    --python_out=$OUT_DIR \
    --grpc_python_out=$OUT_DIR \
    $PROTO_DIR/communications.proto

sed -i '' 's/^import communications_pb2 as /from . import communications_pb2 as /' $OUT_DIR/*_pb2_grpc.py