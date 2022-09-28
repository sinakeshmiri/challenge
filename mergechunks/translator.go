package mergechunks

import (
	"github.com/FaridehGhani/ompfinex_challenge/middle"
	img "github.com/FaridehGhani/ompfinex_challenge/uploadchunks/Image"
)

func ToMiddleImage(src *img.Image) middle.Image {
	if src == nil {
		return middle.Image{}
	}

	return middle.Image{
		SHA256:    src.SHA256,
		Size:      src.Size,
		ChunkSize: src.ChunkSize,
		Chunks:    ToMiddleChunkList(src.Chunks),
	}
}

func ToChunk(src img.Chunk) middle.Chunk {
	return middle.Chunk{
		ID:   src.ID,
		Size: src.Size,
		Data: src.Data,
	}
}

func ToMiddleChunkList(src []img.Chunk) []middle.Chunk {
	dto := make([]middle.Chunk, len(src))

	for k, v := range src {
		dto[k] = ToChunk(v)
	}

	return dto
}
