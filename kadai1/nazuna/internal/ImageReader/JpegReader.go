package ImageReader

import (
	"errors"
	"fmt"
	"image"
	"io/fs"
	"os"
	"strings"
)

type JpegReader struct {
}

func NewJpegReader() *JpegReader {
	return &JpegReader{}
}

func (r *JpegReader) Read(path string, f fs.FileInfo) (image.Image, error) {
	if f.IsDir() {
		return nil, nil
	}
	if !strings.HasSuffix(f.Name(), "jpeg") && !strings.HasSuffix(f.Name(), "jpg") {
		return nil, nil
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer file.Close()

	img, fo, err := image.Decode(file)
	if fo != "jpeg" {
		return nil, errors.New("not jpeg")
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return img, nil
}
