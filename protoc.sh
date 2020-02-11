#! /bin/sh

protoc ./protobuf/*.proto --go_out=./go-server/
