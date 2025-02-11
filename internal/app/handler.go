package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var (
	typeRequest = "request"
	typePublish = "publish"
)

type Request struct {
	Subject   string          `json:"subject"`
	Type      string          `json:"type"`
	Data      json.RawMessage `json:"data"`
	TimeoutMs uint64          `json:"timeout"`
}

func (app *App) handler(ctx *gin.Context) {
	request := Request{
		TimeoutMs: 5_000,
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResp{Error: err.Error()})
		return
	}

	switch request.Type {
	case typeRequest:
		msg, err := app.nc.Request(request.Subject, request.Data, time.Duration(request.TimeoutMs)*time.Millisecond)
		if err != nil {
			ctx.JSON(http.StatusRequestTimeout, errorResp{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, natsMsgToResponse(msg))
		return
	case typePublish:
		err := app.nc.Publish(request.Subject, request.Data)
		if err != nil {
			ctx.JSON(http.StatusRequestTimeout, errorResp{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, nil)
		return
	default:
		ctx.JSON(http.StatusBadRequest, errorResp{Error: fmt.Sprintf("unknown type '%s'", request.Type)})
		return
	}
}

type errorResp struct {
	Error string `json:"error"`
}

func natsMsgToResponse(msg *nats.Msg) interface{} {
	response := make(map[string]interface{})

	var data interface{}

	err := json.Unmarshal(msg.Data, &data)

	if err != nil {
		logrus.Error(err)
	}

	response["data"] = data
	response["headers"] = msg.Header
	response["subject"] = msg.Subject
	response["reply"] = msg.Reply

	return response
}
