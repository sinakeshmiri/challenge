package Image

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

type ImageRepository interface {
	AddImage(image Image) error
	GetImage(sha256 string) *Image
}
