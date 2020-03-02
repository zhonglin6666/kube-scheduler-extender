package app

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"kube-scheduler-extender/pkg/api"
)

func DaemonCmd() *cli.Command {
	return &cli.Command{
		Name:  "daemon",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			if err := startDaemon(c); err != nil {
				logrus.Fatalf("Failed to start daemon , err: %v", err)
			}
			return nil
		},
	}
}

func startDaemon(c *cli.Context) error {
	server := &api.Server{
		Debug: c.Bool("debug"),
	}
	server.Start()

	return nil
}
