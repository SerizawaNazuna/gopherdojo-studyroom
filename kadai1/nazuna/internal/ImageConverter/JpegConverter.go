package ImageConverter

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
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
	out, err := os.Create(fmt.Sprintf("%s/%s.jpg", c.OutDir, c.FileName))
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
