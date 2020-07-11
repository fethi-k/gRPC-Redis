package main

import (
	"context"
	"flag"
	"fmt"
	"main/database"
	"main/proto"
	"net"
	"os"

	"google.golang.org/grpc"
)

var db database.Database

type server struct{}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	databaseImplementation := os.Args[1]
	db, err = database.Factory(databaseImplementation)
	if err != nil {
		panic(err)
	}
	proto.RegisterGrpcServiceServer(grpcServer, &server{})
	fmt.Println("gRPC server 4040 portunu dinliyor...")
	if e := grpcServer.Serve(lis); e != nil {
		panic(err)
	}

}

func (s *server) RecordDB(ctx context.Context, SetPerson *proto.SetPerson) (*proto.ServerResponse, error) {
	value, err := db.RecordDB(SetPerson.GetKey(), SetPerson.GetValue())
	return generateResponse(value, err)
}

func (s *server) GetRecordDB(ctx context.Context, GetPerson *proto.GetPerson) (*proto.ServerResponse, error) {
	value, err := db.GetRecordDB(GetPerson.GetKey())
	return generateResponse(value, err)
}

func generateResponse(value string, err error) (*proto.ServerResponse, error) {
	if err != nil {
		return &proto.ServerResponse{Success: false, Value: value, Error: err.Error()}, nil
	}
	return &proto.ServerResponse{Success: true, Value: value, Error: ""}, nil
}
