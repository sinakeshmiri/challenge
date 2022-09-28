package middle

type RegisterImageRequest struct {
	SHA256    string
	Size      int
	ChunkSize int
}

type UploadImageChunk struct {
	ID   int
	Size int
	Data string
}

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
