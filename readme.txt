cd proto
protoc -I . --go_out=plugins=grpc:. ./user.proto

