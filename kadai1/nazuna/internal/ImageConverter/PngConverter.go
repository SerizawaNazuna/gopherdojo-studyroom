package ImageConverter

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type PngConverter struct {
	Image    image.Image
	OutDir   string
	FileName string
}

func NewPngConverter(image image.Image, outDir string, name string) *PngConverter {
	return &PngConverter{
		Image:    image,
		OutDir:   outDir,
		FileName: name,
	}
}

func (c *PngConverter) Convert() error {
	out, err := os.Create(fmt.Sprintf("%s/%s.png", c.OutDir, c.FileName))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer out.Close()
	errConv := png.Encode(out, c.Image)
	if errConv != nil {
		fmt.Println(err.Error())
		return errConv
	}
	return nil
}
