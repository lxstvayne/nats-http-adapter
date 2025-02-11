package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type App struct {
	nc *nats.Conn
}

func New(nc *nats.Conn) *App {
	return &App{nc: nc}
}

func (app *App) Serve() {
	r := gin.Default()
	r.POST("/", app.handler)
	if err := r.Run(":8080"); err != nil {
		logrus.Error(err)
	}
}
