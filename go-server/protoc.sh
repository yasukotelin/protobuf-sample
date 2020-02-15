#! /bin/sh

cd ./_proto
protoc *.proto --go_out=../protobuf
