package uploadchunks

import (
	"github.com/FaridehGhani/ompfinex_challenge/infra/mongodb"
	"github.com/FaridehGhani/ompfinex_challenge/middle"
)

type Application struct {
	middle.ASCIIArtService
	images mongodb.Images
}

func (app Application) RegisterImage(image middle.RegisterImageRequest) error {
	return app.images.AddImage(RegisterImageRequestToImage(image))
}

func (app Application) UploadImageChunk(sha256 string, chunk middle.UploadImageChunk) error {
	image := app.images.GetImage(sha256)
	if image == nil {
		return middle.ErrImageNotFound
	}
	if err := image.AppendChunk(UploadImageChunkToChunk(chunk)); err != nil {
		return err
	}

	return app.images.AddChunk(*image)
}
