package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FaridehGhani/ompfinex_challenge/ascii_art"
)

type apiHandler struct {
	service ascii_art.AsciiArtApp
}

func (api apiHandler) RegisterImage(ctx *gin.Context) {
	var req RegisterImageRequest

	if err := ctx.BindJSON(&req); err != nil {
		response(ctx, http.StatusBadRequest, ErrMalformedRequest)
		return
	}
	if err := req.validate(); err != nil {
		response(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := api.service.RegisterImage(RegisterImageRequestToImage(req)); err != nil {
		// TODO: response
		return
	}

	response(ctx, http.StatusCreated, SuccessImageRegistered)
}

func (api apiHandler) UploadImageChunk(ctx *gin.Context) {
	var req UploadImageChunkToChunk

	if err := ctx.BindJSON(&req); err != nil {
		response(ctx, http.StatusBadRequest, ErrMalformedRequest)
		return
	}

}

func (api apiHandler) DownloadImage(ctx *gin.Context) {

}

func response(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{"message": msg})
}
