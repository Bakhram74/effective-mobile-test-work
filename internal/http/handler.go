package http

import (
	"github.com/Bakhram74/effective-mobile-test-work.git/config"
	_ "github.com/Bakhram74/effective-mobile-test-work.git/docs"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	config  *config.Config
	service service.Service
}

func NewHandler(config *config.Config, service service.Service) *Handler {

	return &Handler{
		service: service,
		config:  config,
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
	song := router.Group("/song")
	{
		song.POST("/create", h.createSong)
		song.PUT("/update", h.updateSong)
		song.DELETE("/delete", h.deleteSong)
		song.POST("/verses", h.paginatedVerses)
		song.GET("/songs", h.filteredSongs)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, msg)
}
