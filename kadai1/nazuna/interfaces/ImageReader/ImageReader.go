package ImageReader

import (
	"image"
	"io/fs"
)

type ImageReader interface {
	Read(string, fs.FileInfo) (image.Image, error)
}
