package Image

import "github.com/FaridehGhani/ompfinex_challenge/middle"

type Image struct {
	SHA256    string
	Size      int
	ChunkSize int

	Chunks []Chunk
}

type Chunk struct {
	ID   int
	Size int
	Data string
}

func (img *Image) AppendChunk(chunk Chunk) error {
	for _, ch := range img.Chunks {
		if ch.ID == chunk.ID {
			return middle.ErrChunkAlreadyExists
		}
	}
	img.Chunks = append(img.Chunks, chunk)

	return nil
}

type ImageRepository interface {
	AddImage(image Image) error
	AddChunk(image Image) error

	GetImage(sha256 string) *Image
}
