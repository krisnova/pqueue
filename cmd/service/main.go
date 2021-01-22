package main

import (
	"os"

	"github.com/kubicorn/kubicorn/pkg/namer"

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
			&cli.StringFlag{
				Name:        "name",
				Value:       namer.RandomName(),
				Usage:       "A name for your queue, the name is how a queue is recovered from disk in the event of a crash.",
				Destination: &ServiceConfig.Name,
			},
		},
		Action: func(c *cli.Context) error {
			return RunService(ServiceConfig)
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
	// RunConfig is the state in memory for our
	// runtime configuration of a PQueue
	ServiceConfig = &pqueue.PQueueConfig{
		// Optionally configure other sane defaults
	}
)

// RunService is used to run a PQueue given a prepoulated
// PQueueueConfig
func RunService(cfg *pqueue.PQueueConfig) error {
	logger.Debug("RunService for [%s]", cfg.Name)
	q := pqueue.NewPQueue(cfg)
	return q.ListenAndServe() // Todo we can have concurrent messages/errors return here
}
