package main

import (
	"image"
	"io"
	"os"

	"golang.org/x/image/draw"
)

// openImage opens an image file specified by the filename and returns the decoded image and any error encountered.
// It first opens the file, then uses getImage to decode and return the image. The file is closed automatically when done.
func openImage(fileName string) (image.Image, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return getImage(file)
}

// getImage decodes an image from the provided io.Reader and returns the decoded image and any error encountered.
// It uses the image.Decode function to decode the image. If successful, it returns the decoded image.
func getImage(file io.Reader) (image.Image, error) {
	imageSource, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return scaleImage(imageSource, draw.ApproxBiLinear), nil
}

// scaleImage scales the provided image using the given scaler and returns the scaled image.
// It scales the image to half of its original dimensions by default.
func scaleImage(imageSource image.Image, scale draw.Scaler) image.Image {
	imageSourceRectangle := imageSource.Bounds()
	imageDestinationRectangle := image.Rect(0, 0, imageSourceRectangle.Max.X, imageSourceRectangle.Max.Y)
	imageDestination := image.NewRGBA(imageDestinationRectangle)
	scale.Scale(imageDestination, imageDestinationRectangle, imageSource, imageSourceRectangle, draw.Over, nil)
	return imageDestination
}
