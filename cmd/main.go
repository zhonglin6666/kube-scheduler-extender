package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"kube-scheduler-extender/cmd/app"
)

var (
	version = "v1.0.0"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	app := cli.App{
		Name:            "kube-scheduler-extender",
		Usage:           "",
		Version:         version,
		CommandNotFound: cmdNotFound,
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("debug") {
				logrus.SetLevel(logrus.DebugLevel)
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug, d",
				Usage: "enable debug log level",
			},
		},
		Commands: []*cli.Command{
			app.DaemonCmd(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("Run error: %v", err)
	}

	// listening OS shutdown signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	logrus.Infof("Shutdown signal, shutting down http server ...")
}

func cmdNotFound(ctx *cli.Context, cmd string) {
	panic(fmt.Errorf("invalid command: %s", cmd))
}
