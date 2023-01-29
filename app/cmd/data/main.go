package main

import (
	"net"

	"Freebird/app/api/data"
	"Freebird/app/system/log"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	serv := data.NewServer()
	data.RegisterDataServer(s, serv)
	log.Info("listening on port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
