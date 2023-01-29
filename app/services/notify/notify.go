package notify

import (
	"context"
	"encoding/json"
	"sync"

	"Freebird/app/datastore/data"
	"Freebird/app/domain"
	"Freebird/app/system/log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	notifyOnce sync.Once
	notifyInst *NotifySrvc
)

type NotifySrvc struct {
	pool *domain.UserPool
}

func GetNotify() *NotifySrvc {
	notifyInst = &NotifySrvc{
		pool: domain.NewPool(),
	}

	notifyOnce.Do(func() {
		c, err := data.RwInstancePool().Acquire(context.Background())
		defer c.Release()

		if err != nil {
			panic(err)
		}

		_, err = c.Exec(context.Background(), "listen tweet_stream")
		if err != nil {
			panic(err)
		}

		go notifyLoop(c)
	})

	return notifyInst
}

func (n *NotifySrvc) Subscribe(userChan chan<- *data.Bird) {
	n.pool.Subscibe(userChan)
}

func (n *NotifySrvc) Remove(userChan chan<- *data.Bird) {
	n.pool.Remove(userChan)
}

func notifyLoop(c *pgxpool.Conn) {
	for {
		msg, err := c.Conn().WaitForNotification(context.Background())
		if err != nil {
			log.Error("err with notif", err)
			panic(err)
		}

		bird := data.Bird{}
		if err := json.Unmarshal([]byte(msg.Payload), &bird); err != nil {
			log.Error("%v", err)
			continue
		}
		log.Info("ok")

		for _, userChan := range notifyInst.pool.UserChans() {
			userChan <- &bird
		}
	}
}
