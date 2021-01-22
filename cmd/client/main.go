package main

import (
	"fmt"
	"os"

	"github.com/kris-nova/pqueue"

	"github.com/kris-nova/logger"
	"github.com/urfave/cli"
)

// main is the main entry point of the message queue service
func main() {
	logger.Debug("Starting service...")
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "verbose",
				Value:       4,
				Usage:       "Verbosity level (0 off, 1 Critical, 2 Warning, 3 Info, 4 Debug)",
				Destination: &logger.Level,
			},
			//&cli.StringFlag{
			//	Name:        "name",
			//	Value:       namer.RandomName(),
			//	Usage:       "A name for your queue, the name is how a queue is recovered from disk in the event of a crash.",
			//	Destination: &RunConfig.Name,
			//},
		},
		Action: func(c *cli.Context) error {
			return RunService(RunConfig)
		},
	}

	// Run the application
	err := app.Run(os.Args)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
	logger.Always("Bye...")
	os.Exit(0)
}

var (
	ClientConfig = &pqueue.PQueueConfig{}
)

func RunClient(cfg *pqueue.ClientConfig) error {
	logger.Debug("RunClient for [%s]", cfg.Name)
	c := pqueue.NewClient(cfg)
	err := c.Connect()
	if err != nil {
		return fmt.Errorf("unable to connect to service %v", err)
	}
	// TODO Implement client commands here...
	return nil
}
