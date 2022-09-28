package mergechunks

import (
	"bytes"
	"github.com/FaridehGhani/ompfinex_challenge/infra/mongodb"
	"github.com/FaridehGhani/ompfinex_challenge/middle"
	"github.com/FaridehGhani/ompfinex_challenge/middle/proto"
)

type Application struct {
	middle.ASCIIArtService
	grpc proto.AsciiArtServiceClient

	images mongodb.Images
}

func (app Application) DownloadImage(sha256 string) (*string, error) {
	//res, err := app.grpc.GetImage(context.Background(), &proto.ImageRequest{Sha256: sha256})
	//if err != nil {
	//	return nil, err
	//}
	//if res.Image == nil {
	//	return nil, middle.ErrImageNotFound
	//}
	//str := mergeImageChunks(middle.ProtoChunkToChunkListPtrVal(res.Image.Chunks))

	image := app.images.GetImage(sha256)
	if image == nil {
		return nil, middle.ErrImageNotFound
	}
	str := mergeImageChunks(ToMiddleChunkList(image.Chunks))

	return &str, nil
}

func mergeImageChunks(src []middle.Chunk) string {
	chunks := make([]middle.Chunk, len(src))
	for _, v := range src {
		chunks[v.ID] = v
	}

	var b bytes.Buffer
	for _, chunk := range chunks {
		b.WriteString(chunk.Data)
	}

	return b.String()
}
