package pqueue

import (
	"encoding/json"
	"fmt"
	"github.com/kris-nova/logger"
	"sync"
)

type PQueueConfig struct {
	//
	Name string
}

type PQueue struct {
	Config *PQueueConfig
	Queue []*Message // TODO we are putting a LOT in memory Nova
	Count int
	mutex sync.Mutex
}

func NewPQueue(cfg *PQueueConfig) *PQueue {
	logger.Info("Initalizing a new PQueue [%s]", cfg.Name)
	return &PQueue{
		Config: cfg,
	}
}

func (q *PQueue) Length() int {
	return q.Count
}

func (q *PQueue) Enqueue(message *Message) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	// Procedure is critical here
	q.Queue = append(q.Queue, message)
	// TODO Flush to disk
	q.Count++
	return nil
}

func (q *PQueue) Dequeue() (*Message, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Count < 1 {
		return nil, fmt.Errorf("no messages in queue")
	}
	var message *Message
	message, q.Queue = q.Queue[0], q.Queue[1:]
	q.Count--
	return message, nil
}

func (q *PQueue) String() (string, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	jbytes, err := json.Marshal(&q)
	if err != nil {
		return "", fmt.Errorf("unable to json marshal: %v", err)
	}
	return string(jbytes), nil
}

func (q *PQueue) Empty() error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	var empty []*Message
	q.Count = 0
	q.Queue = empty
	return nil
}

// Recover will return the number of messages
// recovered from disk in the event of an outage.
func (q *PQueue) Recover() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return -1
}


func (q *PQueue) ListenAndServe() error {
	logger.Info("ListenAndServe...")
	for {
	}
	return nil
}
