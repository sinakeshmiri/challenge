package uploadchunks

import (
	"github.com/FaridehGhani/ompfinex_challenge/middle"
	"github.com/FaridehGhani/ompfinex_challenge/uploadchunks/Image"
)

func RegisterImageRequestToImage(src middle.RegisterImageRequest) Image.Image {
	return Image.Image{
		SHA256:    src.SHA256,
		Size:      src.Size,
		ChunkSize: src.ChunkSize,

		Chunks: nil,
	}
}

func UploadImageChunkToChunk(chunk middle.UploadImageChunk) Image.Chunk {
	return Image.Chunk{
		ID:   chunk.ID,
		Size: chunk.Size,
		Data: chunk.Data,
	}
}
