package middle

type ASCIIArtService interface {
	RegisterImage(image RegisterImageRequest) error
	UploadImageChunk(sha256 string, chunk UploadImageChunk) error

	DownloadImage(sha256 string) (*string, error)
}
