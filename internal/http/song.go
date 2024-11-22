package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	arg := db.CreateSongParams{
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

type songUpdateReq struct {
	ID          string `json:"id" binding:"required"`
	Group       string `json:"group,omitempty"`
	Name        string `json:"name,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}

func (h *Handler) updateSong(ctx *gin.Context) {
	var reqBody songUpdateReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	id, err := uuid.Parse(reqBody.ID)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	song, err := h.service.Song.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(ctx, http.StatusNotFound, err.Error())
			slog.Error(err.Error())
			return
		}
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}
	songBytes, err := json.Marshal(song)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}
	patchedJSON, err := jsonpatch.MergePatch(songBytes, reqBodyBytes)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}

	var updatedSong db.Song
	_ = json.Unmarshal(patchedJSON, &updatedSong)

	arg := db.UpdateSongParams{
		ID:          id,
		Group:       updatedSong.Group,
		Name:        updatedSong.Name,
		ReleaseDate: updatedSong.ReleaseDate,
		Text:        updatedSong.Text,
		Link:        updatedSong.Link,
	}

	song, err = h.service.Song.Update(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, song)
}

type songDeleteReq struct {
	ID string `json:"id" binding:"required"`
}

func (h *Handler) deleteSong(ctx *gin.Context) {
	var reqBody songDeleteReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}
	id, err := uuid.Parse(reqBody.ID)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return
	}

	err = h.service.Song.Delete(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}
