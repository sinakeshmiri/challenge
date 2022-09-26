package middle

import "errors"

var (
	ErrImageAlreadyExists   = errors.New("image already exists")
	ErrChunkAlreadyExists   = errors.New("chunk already exists")
	ErrImageNotFound        = errors.New("image not found")
	ErrMalformedRequest     = errors.New("malformed request")
	ErrUnsupportedMediaType = errors.New("unsupported payload format")
)
