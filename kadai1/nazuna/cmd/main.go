package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	"io/fs"
	"path/filepath"

	inc "github.com/SerizawaNazuna/gopherdojo-studyroom/kadai1/nazuna/interfaces/ImageConverter"
	in "github.com/SerizawaNazuna/gopherdojo-studyroom/kadai1/nazuna/interfaces/ImageReader"
	imc "github.com/SerizawaNazuna/gopherdojo-studyroom/kadai1/nazuna/internal/ImageConverter"
	im "github.com/SerizawaNazuna/gopherdojo-studyroom/kadai1/nazuna/internal/ImageReader"
)

var srcDir = flag.String("src", "", "path of srcDir")
var outDir = flag.String("out", "", "path of outDir")
var inImageType = flag.String("inImage", "jpg", "convert target image type - e.g. jpg")
var outImageType = flag.String("outImage", "png", "convert dist image type - e.g. png")

func main() {
	flag.Parse()
	filepath.Walk(*srcDir, convert)
}

func convert(path string, info fs.FileInfo, err error) error {
	var r in.ImageReader
	switch *inImageType {
	case "jpg":
		r = im.NewJpegReader()
	case "png":
		r = im.NewPngReader()
	default:
		fmt.Printf("not readable data type: %s", *inImageType)
		panic("format wrong.")
	}
	image, err := r.Read(path, info)
	if err != nil {
		return err
	}
	if image == nil {
		fmt.Printf("not convert target file, skipping: %s", path)
		return nil
	}
	inName := info.Name()
	extension := filepath.Ext(inName)
	outName := inName[0 : len(inName)-len(extension)]
	var c inc.ImageConverter

	switch *outImageType {
	case "png":
		c = imc.NewPngConverter(image, *outDir, outName)
	case "jpg":
		c = imc.NewJpegConverter(image, *outDir, outName)
	default:
		// end.
		fmt.Printf("not convertible format: %s", *outImageType)
		panic("format wrong.")
	}
	errConv := c.Convert()
	if errConv != nil {
		return errConv
	}
	return nil
}
