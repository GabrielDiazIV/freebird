package main

import (
	"Freebird/app/datastore/data"
	"Freebird/app/services/notify"
	"Freebird/app/system"
	"fmt"
)

type test struct {
	id       int
	birdChan chan *data.Bird
}

func (t *test) testRun() {
	for bird := range t.birdChan {
		fmt.Printf("%d: %s", t.id, bird.Name)
	}
}

func main() {
	n := notify.Notify()
	for i := 0; i < 40; i++ {
		t := test{
			id:       i,
			birdChan: make(chan *data.Bird),
		}
		n.Subscribe(t.birdChan)
		go t.testRun()
	}

	handler := system.GetSigHandler()
	handler.SetExit(false)
	handler.Wait()
}
