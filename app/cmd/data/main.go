package main

import (
	"net"

	"Freebird/app/api/data"
	"Freebird/app/system/log"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	data.RegisterDataServer(s, data.NewServer())
	log.Info("listening on port 9090")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
