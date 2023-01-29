package data

import (
	"context"
	"sync"

	"Freebird/app/datastore/data"
	"Freebird/app/services/notify"
	"Freebird/app/system/log"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	servOnce     sync.Once
	servInstance *server
)

type server struct {
	UnimplementedDataServer
	db *data.Queries
}

func t2t(tweet *data.Tweet) *Tweet {
	return &Tweet{
		BirdFk:         tweet.BirdFk,
		Body:           tweet.Body,
		AuthorName:     tweet.AuthorName,
		AuthorUsername: tweet.AuthorUsername,
		PostTime:       timestamppb.New(tweet.PostTime),
		Score:          float32(tweet.Score),
		Certainty:      float32(tweet.Certainty),
	}
}
func b2b(bird *data.Bird) *Bird {
	return &Bird{
		Id:         bird.ID,
		Name:       bird.Name,
		EntityType: FreebirdEntity(bird.EntityType),
		BirdFk:     bird.BirdFk,
		Score:      float32(bird.Score),
		NPositive:  bird.NPositive,
		NNegative:  bird.NNegative,
		ImgUrl:     bird.ImgUrl,
	}
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
		log.Info("sending: %+v\n", bird)
		server.Send(b2b(bird))
	}

	return nil
}

func NewServer() *server {
	c := data.RwInstancePool()
	return &server{
		db: data.New(c),
	}
}

func (s *server) GetBird(ctx context.Context, req *BirdRequest) (*Bird, error) {
	bird, err := s.db.GetBird(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return b2b(bird), err
}

func (s *server) GetBirdTweets(ctx context.Context, req *BirdRequest) (*Tweets, error) {
	tweets, err := s.db.GetTweets(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	tweets_res := &Tweets{
		Tweets: make([]*Tweet, len(tweets)),
	}

	for i, t := range tweets {
		tweets_res.Tweets[i] = t2t(t)
	}

	return tweets_res, err
}

func (s *server) GetAllBirds(ctx context.Context, req *BirdsRequest) (*Birds, error) {

	birds, err := s.db.GetBirds(ctx)
	if err != nil {
		return nil, err
	}

	birds_res := &Birds{
		Birds: make([]*Bird, len(birds)),
	}

	for i, bird := range birds {
		birds_res.Birds[i] = b2b(bird)
	}

	return birds_res, nil
}
