package data

import (
	"context"

	"Freebird/app/datastore/data"
	"Freebird/app/services/notify"
)

type server struct {
	UnimplementedDataServer
	n notify.NotifySrvc
}

func (s *server) TweetStream(req *TweetStreamRequest, server Data_TweetStreamServer) error {
	userChan := make(chan *data.Bird)
	s.n.Subscribe(userChan)

	return nil
}

func (s *server) EntityData(ctx context.Context, req *EntityRequest) (*EntityResponse, error) {
	return nil, nil
}
