package pqueue

import "github.com/kris-nova/logger"

type ClientConfig struct {
	//
	Name string
}

type PQueueClient struct {
	//
	Config *ClientConfig
}

func NewClient(cfg *ClientConfig) *PQueueClient {
	logger.Info("Initalizing a new PQueue [%s]", cfg.Name)
	return &PQueueClient{
		Config: cfg,
	}
}

func (c *PQueueClient) Connect() error {
	logger.Info("Connecting...")
	for {
	}
	return nil
}
