package grpc

//go:generate protoc -I ../../../proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative message/v1/message.proto
