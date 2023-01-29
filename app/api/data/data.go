package data

import (
	"Freebird/app/datastore/data"
	"Freebird/app/services/notify"
	"context"
)

type server struct {
	UnimplementedDataServer
	n notify.NotifySrvc
}

func (s *server) TweetStream(req *TweetStreamRequest, server Data_TweetStreamServer) error {
	userChan := make(chan *data.Tweet)
	s.n.Subscribe(userChan)

	/*
		for tweet := range userChan {
			server.Send(&TweetStreamResponse{
				EntityName:      tweet.AuthorName.String,
				EntitySentiment: tweet.,
				EntityImage:     "",
				EntityType:      0,
			})
		}
	*/
	return nil
}

func (s *server) EntityData(ctx context.Context, req *EntityRequest) (*EntityResponse, error) {
	return nil, nil
}
