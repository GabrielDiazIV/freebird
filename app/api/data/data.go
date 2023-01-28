package data

import "context"

type server struct {
	UnimplementedDataServer
}

func (s *server) TweetStream(req *TweetStreamRequest, server Data_TweetStreamServer) error {
	return nil
}

func (s *server) EntityData(ctx context.Context, req *EntityRequest) (*EntityResponse, error) {
	return nil, nil
}
