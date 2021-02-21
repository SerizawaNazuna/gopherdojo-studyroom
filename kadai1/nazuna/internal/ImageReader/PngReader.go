package ImageReader

import (
	"errors"
	"fmt"
	"image"
	"io/fs"
	"os"
	"strings"
)

type PngReader struct {
}

func NewPngReader() *PngReader {
	return &PngReader{}
}

func (r *PngReader) Read(path string, f fs.FileInfo) (image.Image, error) {
	if f.IsDir() {
		return nil, nil
	}
	if !strings.HasSuffix(f.Name(), "png") {
		return nil, nil
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer file.Close()

	img, fo, err := image.Decode(file)
	if fo != "png" {
		return nil, errors.New("not png")
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return img, nil
}
