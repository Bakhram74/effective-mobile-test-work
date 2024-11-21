package http

import (
	"log/slog"
	"net/http"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
	"github.com/gin-gonic/gin"
)

type songCreateReq struct {
	Group       string `json:"group" binding:"required"`
	Name        string `json:"name" binding:"required"`
	ReleaseDate string `json:"release_date" binding:"required"`
	Text        string `json:"text" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

func (h *Handler) createSong(ctx *gin.Context) {
	var reqBody songCreateReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	arg := db.CreateSongsParams{
		Group:       reqBody.Group,
		Name:        reqBody.Name,
		ReleaseDate: reqBody.ReleaseDate,
		Text:        reqBody.Text,
		Link:        reqBody.Link,
	}

	song, err := h.service.Song.Create(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, song)
}
