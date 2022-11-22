package controller

import "net/http"

type Logger interface {
	Log(message string)
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	logger Logger
	logic  Logic
}

func (controller Controller) SayHello(res http.ResponseWriter, req *http.Request) {
	controller.logger.Log("Request to Say Hello")

	userID := req.URL.Query().Get("user_id")

	message, err := controller.logic.SayHello(userID)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))

		return
	}

	res.Write([]byte(message))
}

func NewController(logger Logger, logic Logic) Controller {
	return Controller{
		logger: logger,
		logic:  logic,
	}
}
