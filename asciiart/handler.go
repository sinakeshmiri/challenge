package ascii_art

import (
	"github.com/FaridehGhani/ompfinex_challenge/infra/mongodb"
	"github.com/FaridehGhani/ompfinex_challenge/middle"
)

type AsciiArtApp struct {
	middle.ASCIIArtService
	images mongodb.Images
}

func (app AsciiArtApp) RegisterImage(image middle.RegisterImageRequest) error {
	return app.images.AddImage(RegisterImageRequestToImage(image))
}

func (app AsciiArtApp) NewImageChunk(sha256 string, chunk middle.UploadImageChunk) error {
	return nil
}
