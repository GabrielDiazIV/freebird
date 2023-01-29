package data

import (
	"context"

	"Freebird/app/datastore/data"
	"Freebird/app/services/notify"
)

type server struct {
	UnimplementedDataServer
}

func (s *server) TweetStream(req *TweetStreamRequest, server Data_TweetStreamServer) error {
	userChan := make(chan *data.Bird)
	notify.GetNotify().Subscribe(userChan)

	defer func() {
		if err := recover(); err != nil {
			notify.GetNotify().Remove(userChan)
		}
	}()

	for bird := range userChan {
		server.Send(&TweetStreamResponse{
			EntityType: FreebirdEntity(bird.EntityType),
			EntityName: bird.Name,
			Score:      float32(bird.Score),
			Positive:   bird.NPositive,
			Negative:   bird.NNegative,
			ImageUrl:   bird.ImgUrl,
		})
	}

	return nil
}

func NewServer() *server {
	return &server{}
}

func (s *server) EntityData(ctx context.Context, req *EntityRequest) (*EntityResponse, error) {
	return nil, nil
}
