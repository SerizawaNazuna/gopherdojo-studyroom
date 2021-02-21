package ImageConverter

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

type JpegConverter struct {
	Image    image.Image
	OutDir   string
	FileName string
}

func NewJpegConverter(image image.Image, outDir string, name string) *JpegConverter {
	return &JpegConverter{
		Image:    image,
		OutDir:   outDir,
		FileName: name,
	}
}

func (c *JpegConverter) Convert() error {
	out, err := os.Create(filepath.Join(c.OutDir, c.FileName+".jpg"))
	if err != nil {
		return err
	}
	defer out.Close()
	errConv := jpeg.Encode(out, c.Image, &jpeg.Options{Quality: 100})
	if errConv != nil {
		return errConv
	}
	return nil
}
