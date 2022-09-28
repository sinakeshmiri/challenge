package middle

import (
	"github.com/FaridehGhani/ompfinex_challenge/middle/proto"
)

func ProtoImageToImage(src *proto.Image) *Image {
	if src == nil {
		return nil
	}

	return &Image{
		SHA256:    src.Sha256,
		Size:      int(src.Size),
		ChunkSize: int(src.ChunkSize),
		Chunks:    ProtoChunkToChunkListPtrVal(src.Chunks),
	}
}

func ProtoChunkToChunkPrtVal(src *proto.Chunk) Chunk {
	if src == nil {
		return Chunk{}
	}

	return Chunk{
		ID:   int(src.ID),
		Size: int(src.Size),
		Data: src.Data,
	}
}

func ProtoChunkToChunkListPtrVal(src []*proto.Chunk) []Chunk {
	if src == nil {
		return nil
	}

	dto := make([]Chunk, len(src))
	for k, v := range src {
		dto[k] = ProtoChunkToChunkPrtVal(v)
	}

	return dto
}
