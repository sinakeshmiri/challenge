package ascii_art

import (
	"github.com/FaridehGhani/ompfinex_challenge/ascii_art/Image"
	"github.com/FaridehGhani/ompfinex_challenge/middle"
)

func RegisterImageRequestToImage(src middle.RegisterImageRequest) Image.Image {
	return Image.Image{
		SHA256:    src.SHA256,
		Size:      src.Size,
		ChunkSize: src.ChunkSize,

		Chunks: nil,
	}
}
