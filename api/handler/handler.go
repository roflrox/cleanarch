package handler

import (
	"cleanarch/usecases"
	"net/http"
)
import "github.com/gin-gonic/gin"

type GinHanlder struct {
	Usecase usecases.UseCase
}

func NewGinHandler(usc usecases.UseCase) *gin.Engine {
	h := &GinHanlder{Usecase: usc}

	router := gin.Default()
	router.GET("/hello", h.sayhello)
	return router
}

func (h GinHanlder) sayhello(context *gin.Context) {
	msg, err := h.Usecase.SayHello(context)
	if err != nil {
		context.Error(err)
		return
	}
	context.String(http.StatusOK, msg)
}
