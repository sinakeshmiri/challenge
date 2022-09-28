package uploadchunks

import (
	"github.com/FaridehGhani/ompfinex_challenge/middle"
	img "github.com/FaridehGhani/ompfinex_challenge/uploadchunks/Image"
)

func RegisterImageRequestToImage(src middle.RegisterImageRequest) img.Image {
	return img.Image{
		SHA256:    src.SHA256,
		Size:      src.Size,
		ChunkSize: src.ChunkSize,

		Chunks: nil,
	}
}

func UploadImageChunkToChunk(src middle.UploadImageChunk) img.Chunk {
	return img.Chunk{
		ID:   src.ID,
		Size: src.Size,
		Data: src.Data,
	}
}
