package middle

type ASCIIArtService interface {
	RegisterImage(image RegisterImageRequest) error
	NewImageChunk(sha256 string, chunk UploadImageChunk) error
}
