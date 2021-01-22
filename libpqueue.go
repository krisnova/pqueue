package pqueue

import "github.com/kris-nova/logger"

type PQueueConfig struct {
	//
	Name string
}

type PQueue struct {
	//
	Config *PQueueConfig
}

func NewPQueue(cfg *PQueueConfig) *PQueue {
	logger.Info("Initalizing a new PQueue [%s]", cfg.Name)
	return &PQueue{
		Config: cfg,
	}
}

func (q *PQueue) ListenAndServe() error {
	logger.Info("ListenAndServe...")
	for {
	}
	return nil
}
