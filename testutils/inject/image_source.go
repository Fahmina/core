package inject

import (
	"context"
	"image"

	"github.com/edaniels/gostream"
)

type ImageSource struct {
	gostream.ImageSource
	NextFunc  func(ctx context.Context) (image.Image, func(), error)
	CloseFunc func() error
}

func (is *ImageSource) Next(ctx context.Context) (image.Image, func(), error) {
	if is.NextFunc == nil {
		return is.ImageSource.Next(ctx)
	}
	return is.NextFunc(ctx)
}

func (is *ImageSource) Close() error {
	if is.CloseFunc == nil {
		return is.ImageSource.Close()
	}
	return is.CloseFunc()
}