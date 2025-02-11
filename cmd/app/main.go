package main

import (
	"lxstvayne/nats-http-adapter/internal/app"
	"lxstvayne/nats-http-adapter/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	nc, err := nats.Connect(cfg.NatsURL)

	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)

	app := app.New(nc)

	logrus.Info("Listening...")

	app.Serve()
}
