package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FaridehGhani/ompfinex_challenge/mergechunks"
	"github.com/FaridehGhani/ompfinex_challenge/uploadchunks"
)

type apiHandler struct {
	upl uploadchunks.Application
	dwl mergechunks.Application
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
	if err := api.upl.RegisterImage(RegisterImageRequestToImage(req)); err != nil {
		response(ctx, statusCode(err.Error()), err.Error())
		return
	}

	response(ctx, http.StatusCreated, SuccessImageRegistered)
}

func (api apiHandler) UploadImageChunk(ctx *gin.Context) {
	var req UploadImageChunkRequest

	sha256 := ctx.Param("sha256")
	if err := ctx.BindJSON(&req); err != nil {
		response(ctx, http.StatusBadRequest, ErrMalformedRequest)
		return
	}
	if err := req.validate(); err != nil {
		response(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := api.upl.UploadImageChunk(sha256, UploadImageChunkToChunk(req)); err != nil {
		response(ctx, statusCode(err.Error()), err.Error())
		return
	}

	response(ctx, http.StatusCreated, SuccessChunkUploaded)
}

func (api apiHandler) DownloadImage(ctx *gin.Context) {
	sha256 := ctx.Param("sha256")
	image, err := api.dwl.DownloadImage(sha256)
	if err != nil {
		response(ctx, statusCode(err.Error()), err.Error())
		return
	}
	if image == nil {
		response(ctx, http.StatusNotFound, ErrImageNotFound)
	}

	ctx.String(http.StatusOK, *image)
}

func response(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{"message": msg})
}

func statusCode(msg string) int {
	switch msg {
	case ErrImageAlreadyExists:
		return http.StatusConflict

	case ErrChunkAlreadyExists:
		return http.StatusConflict

	case ErrImageNotFound:
		return http.StatusNotFound

	case ErrMalformedRequest:
		return http.StatusBadRequest

	case ErrUnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	}

	return -1
}
