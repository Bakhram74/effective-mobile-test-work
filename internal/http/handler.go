package http

import (
	"net/http"

	"github.com/Bakhram74/effective-mobile-test-work.git/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	config *config.Config
}

func NewHandler(config *config.Config) *Handler {

	return &Handler{

		config: config,
	}

}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {

	{
		router.GET("/status", h.getStatus)
	}
}

func (h *Handler) getStatus(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "ok")
}
