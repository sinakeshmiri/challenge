package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"

	img "github.com/FaridehGhani/ompfinex_challenge/asciiart/Image"
	"github.com/FaridehGhani/ompfinex_challenge/middle"
)

type Images struct {
	db Storage
}

func (r Images) AddImage(image img.Image) error {
	entity := ImageToDb(image)

	if count := r.db.count(bson.M{"sha256": entity.SHA256}); count > 0 {
		return middle.ErrImageAlreadyExists
	}
	return r.db.insert(entity)
}

func (r Images) AddChunk(image img.Image) error {
	entity := ImageToDb(image)

	return r.db.replace(bson.M{"sha256": entity.SHA256}, image)
}

func (r Images) GetImage(sha256 string) *img.Image {
	entity := r.db.get(bson.M{"sha256": sha256})

	var image Image
	bytes, err := bson.Marshal(entity)
	if err != nil {
		return nil
	}
	err = bson.Unmarshal(bytes, &image)
	if err != nil {
		return nil
	}

	return DbToImageValPtr(image)
}

type Image struct {
	SHA256    string `bson:"sha256" json:"sha256"`
	Size      int    `bson:"size" json:"size"`
	ChunkSize int    `bson:"chunk_size" json:"chunk_size"`

	Chunks []Chunk `bson:"chunks"`
}

type Chunk struct {
	ID   int    `bson:"id" json:"id"`
	Size int    `bson:"size" json:"size"`
	Data string `bson:"data" json:"data"`
}

func ImageToDb(src img.Image) Image {
	return Image{
		SHA256:    src.SHA256,
		Size:      src.Size,
		ChunkSize: src.ChunkSize,

		Chunks: ChunkToDbList(src.Chunks),
	}
}

func DbToImageValPtr(src Image) *img.Image {
	return &img.Image{
		SHA256:    src.SHA256,
		Size:      src.Size,
		ChunkSize: src.ChunkSize,

		Chunks: DbToChunkList(src.Chunks),
	}
}

func ChunkToDb(src img.Chunk) Chunk {
	return Chunk{
		ID:   src.ID,
		Size: src.Size,
		Data: src.Data,
	}
}

func DbToChunk(src Chunk) img.Chunk {
	return img.Chunk{
		ID:   src.ID,
		Size: src.Size,
		Data: src.Data,
	}
}

func ChunkToDbList(src []img.Chunk) []Chunk {
	dtos := make([]Chunk, len(src))

	for k, v := range src {
		dtos[k] = ChunkToDb(v)
	}

	return dtos
}

func DbToChunkList(src []Chunk) []img.Chunk {
	dtos := make([]img.Chunk, len(src))

	for k, v := range src {
		dtos[k] = DbToChunk(v)
	}

	return dtos
}
