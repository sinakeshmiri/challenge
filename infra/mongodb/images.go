package mongodb

import (
	Image2 "github.com/FaridehGhani/ompfinex_challenge/ascii_art/Image"
)

type Images struct {
	db Storage
}

func (r Images) AddImage(image Image2.Image) error {
	img := ImageToDb(image)

	return r.db.insert(img)
}

func (r Images) GetImage(sha256 string) *Image2.Image {
	//TODO implement me
	panic("implement me")
}

type Image struct {
	SHA256    string `bson:"sha256"`
	Size      int    `bson:"size"`
	ChunkSize int    `bson:"chunk_size"`

	Chunks []Chunk `bson:"chunks"`
}

type Chunk struct {
	ID   int    `bson:"id"`
	Size int    `bson:"size"`
	Data string `bson:"data"`
}

func ImageToDb(src Image2.Image) Image {
	return Image{
		SHA256:    src.SHA256,
		Size:      src.Size,
		ChunkSize: src.ChunkSize,

		Chunks: ChunkToDbList(src.Chunks),
	}
}

func ChunkToDb(src Image2.Chunk) Chunk {
	return Chunk{
		ID:   src.ID,
		Size: src.Size,
		Data: src.Data,
	}
}

func ChunkToDbList(src []Image2.Chunk) []Chunk {
	dtos := make([]Chunk, len(src))

	for k, v := range src {
		dtos[k] = ChunkToDb(v)
	}

	return dtos
}
