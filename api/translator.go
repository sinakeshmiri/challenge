package api

import (
	"errors"
	"strings"

	"github.com/FaridehGhani/ompfinex_challenge/middle"
)

var (
	SuccessImageRegistered = "image successfully registered"
	SuccessChunkUploaded   = "chunk successfully uploaded"
	SuccessImageDownloaded = "image successfully downloaded"

	ErrImageAlreadyExists   = "image already exists"
	ErrChunkAlreadyExists   = "chunk already exists"
	ErrImageNotFound        = "image not found"
	ErrMalformedRequest     = "malformed request"
	ErrUnsupportedMediaType = "unsupported payload format"
)

type RegisterImageRequest struct {
	SHA256    string `json:"sha256"`
	Size      int    `json:"size"`
	ChunkSize int    `json:"chunk_size"`
}

func (r RegisterImageRequest) validate() error {
	if len(strings.TrimSpace(r.SHA256)) == 0 {
		return errors.New(ErrMalformedRequest)
	}
	if r.Size < 0 || r.ChunkSize < 0 {
		return errors.New(ErrMalformedRequest)
	}

	return nil
}

func RegisterImageRequestToImage(req RegisterImageRequest) middle.RegisterImageRequest {
	return middle.RegisterImageRequest{
		SHA256:    req.SHA256,
		Size:      req.Size,
		ChunkSize: req.ChunkSize,
	}
}

type UploadImageChunkRequest struct {
	ID   int    `json:"id"`
	Size int    `json:"size"`
	Data string `json:"data"`
}

func (u UploadImageChunkRequest) validate() error {
	if u.ID < 0 || u.Size < 0 {
		return errors.New(ErrMalformedRequest)
	}

	return nil
}

func UploadImageChunkToChunk(req UploadImageChunkRequest) middle.UploadImageChunk {
	return middle.UploadImageChunk{
		ID:   req.ID,
		Size: req.Size,
		Data: req.Data,
	}
}
