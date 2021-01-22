package main

import (
	"fmt"
	"github.com/kris-nova/logger"
	"github.com/kris-nova/pqueue"
	"github.com/kubicorn/kubicorn/pkg/namer"
	"os"
)

// This example will show how to use the queue
// within the context of a single process in
// a systems process queue
//
// There will be no client/server interaction
// however there will be client and service
// implementations

func main() {

	cfg := &pqueue.PQueueConfig{
		Name: "ExampleQueue",
	}
	q := pqueue.NewPQueue(cfg)
	i := q.Recover()
	logger.Always("Recovered [%d] messages", i)
	m1 := &pqueue.Message{
		Key: namer.RandomName(),
		Value : struct {
			Data string
		}{
				Data: namer.RandomName(),
		},
	}
	err := q.Enqueue(m1)
	if err != nil {
		logger.Critical("Error: %v", err)
		os.Exit(1)
	}
	newM, err := q.Dequeue()
	if err != nil {
		logger.Critical("Error: %v", err)
		os.Exit(1)
	}
	fmt.Println(newM)
	m2 := &pqueue.Message{
		Key: namer.RandomName(),
		Value : struct {
			Data string
		}{
			Data: namer.RandomName(),
		},
	}
	err = q.Enqueue(m2)
	if err != nil {
		logger.Critical("Error: %v", err)
		os.Exit(1)
	}
}













