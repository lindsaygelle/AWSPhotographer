package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"

	"golang.org/x/image/draw"
)

const (
	drawScalerApproxBiLinear  = "APPROX_BILINEAR"
	drawScalerBiLinear        = "BI_LINEAR"
	drawScalerCatmullRom      = "CATMULL_ROM"
	drawScalerNearestNeighbor = "NEAREST_NEIGHBOR"
)

func compressImageJPG(fileDirectory string, fileName string, drawInterpolator draw.Interpolator) {
	source, err := openImage(fileDirectory, fileName)
	if err != nil {
		log.Fatalln(err)
	}
	destinationRectangle := image.Rect(0, 0, source.Bounds().Max.X/2, source.Bounds().Max.Y/2)
	res := scaleImage(source, destinationRectangle, drawInterpolator)
	err = saveImageJPG(fileDirectory, fmt.Sprintf("%s.JPG", "X"), res)
	if err != nil {
		log.Fatalln(err)
	}
}

// getDrawScaler returns a draw scaler implementation.
func getDrawScaler(drawScalerName string) draw.Scaler {
	switch strings.ToUpper(drawScalerName) {
	case drawScalerApproxBiLinear:
		return draw.ApproxBiLinear
	case drawScalerBiLinear:
		return draw.BiLinear
	case drawScalerCatmullRom:
		return draw.CatmullRom
	default:
		return draw.NearestNeighbor
	}
}

// openImage opens an image file from the OS and returns an image interface.
func openImage(fileDirectory string, fileName string) (image.Image, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s", fileDirectory, fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func scaleImage(source image.Image, rectangle image.Rectangle, scale draw.Scaler) image.Image {
	destination := image.NewRGBA(rectangle)
	scale.Scale(destination, rectangle, source, source.Bounds(), draw.Over, nil)
	return destination
}

func saveImageJPG(fileDirectory string, fileName string, img image.Image) error {
	file, err := os.Open(fmt.Sprintf("%s/%s", fileDirectory, fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}
	return nil
}
