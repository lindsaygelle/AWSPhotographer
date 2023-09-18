package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

// Exif is the interface for an Exif entity. It defines methods for retrieving Exif tags.
type Exif interface {
	// Get returns the ExifTag associated with the specified field name.
	// If the field name is not found or there is an error retrieving the tag, an error is returned.
	Get(exif.FieldName) (ExifTag, error)
}

// ExifTag is the interface for Exif metadata tags. It represents a single Exif tag's value and properties.
type ExifTag interface {
	// Int returns the integer value of the Exif tag.
	// If the tag value cannot be converted to an integer or there is an error, an error is returned.
	Int(int) (int, error)

	// String returns the string representation of the Exif tag value.
	String() string

	// StringVal returns the string value of the Exif tag.
	// If the tag value cannot be converted to a string or there is an error, an error is returned.
	StringVal() (string, error)
}

// ExifMetadata contains the known exif tags that can be accessed from an image file.
type ExifMetadata struct {
	ApertureValue                    interface{} `json:"ApertureValue"`
	Artist                           interface{} `json:"Artist"`
	BitsPerSample                    interface{} `json:"BitsPerSample"`
	BrightnessValue                  *string     `json:"BrightnessValue"`
	CFAPattern                       interface{} `json:"CFAPattern"`
	ColorSpace                       *int        `json:"ColorSpace"`
	ComponentsConfiguration          *string     `json:"ComponentsConfiguration"`
	CompressedBitsPerPixel           *string     `json:"CompressedBitsPerPixel"`
	Compression                      interface{} `json:"Compression"`
	Contrast                         interface{} `json:"Contrast"`
	Copyright                        interface{} `json:"Copyright"`
	CustomRendered                   interface{} `json:"CustomRendered"`
	DateTime                         interface{} `json:"DateTime"`
	DateTimeDigitized                interface{} `json:"DateTimeDigitized"`
	DateTimeOriginal                 interface{} `json:"DateTimeOriginal"`
	DeviceSettingDescription         interface{} `json:"DeviceSettingDescription"`
	DigitalZoomRatio                 interface{} `json:"DigitalZoomRatio"`
	ExifIFDPointer                   interface{} `json:"ExifIFDPointer"`
	ExifVersion                      interface{} `json:"ExifVersion"`
	ExposureBiasValue                interface{} `json:"ExposureBiasValue"`
	ExposureIndex                    interface{} `json:"ExposureIndex"`
	ExposureMode                     interface{} `json:"ExposureMode"`
	ExposureProgram                  interface{} `json:"ExposureProgram"`
	ExposureTime                     interface{} `json:"ExposureTime"`
	FNumber                          interface{} `json:"FNumber"`
	FileSource                       interface{} `json:"FileSource"`
	Flash                            interface{} `json:"Flash"`
	FlashEnergy                      interface{} `json:"FlashEnergy"`
	FlashpixVersion                  interface{} `json:"FlashpixVersion"`
	FocalLength                      interface{} `json:"FocalLength"`
	FocalLengthIn35mmFilm            interface{} `json:"FocalLengthIn35mmFilm"`
	FocalPlaneResolutionUnit         interface{} `json:"FocalPlaneResolutionUnit"`
	FocalPlaneXResolution            interface{} `json:"FocalPlaneXResolution"`
	FocalPlaneYResolution            interface{} `json:"FocalPlaneYResolution"`
	GPSAltitude                      interface{} `json:"GPSAltitude"`
	GPSAltitudeRef                   interface{} `json:"GPSAltitudeRef"`
	GPSAreaInformation               interface{} `json:"GPSAreaInformation"`
	GPSDOP                           interface{} `json:"GPSDOP"`
	GPSDateStamp                     interface{} `json:"GPSDateStamp"`
	GPSDestBearing                   interface{} `json:"GPSDestBearing"`
	GPSDestBearingRef                interface{} `json:"GPSDestBearingRef"`
	GPSDestDistance                  interface{} `json:"GPSDestDistance"`
	GPSDestDistanceRef               interface{} `json:"GPSDestDistanceRef"`
	GPSDestLatitude                  interface{} `json:"GPSDestLatitude"`
	GPSDestLatitudeRef               interface{} `json:"GPSDestLatitudeRef"`
	GPSDestLongitude                 interface{} `json:"GPSDestLongitude"`
	GPSDestLongitudeRef              interface{} `json:"GPSDestLongitudeRef"`
	GPSDifferential                  interface{} `json:"GPSDifferential"`
	GPSImgDirection                  interface{} `json:"GPSImgDirection"`
	GPSImgDirectionRef               interface{} `json:"GPSImgDirectionRef"`
	GPSInfoIFDPointer                interface{} `json:"GPSInfoIFDPointer"`
	GPSLatitude                      interface{} `json:"GPSLatitude"`
	GPSLatitudeRef                   interface{} `json:"GPSLatitudeRef"`
	GPSLongitude                     interface{} `json:"GPSLongitude"`
	GPSLongitudeRef                  interface{} `json:"GPSLongitudeRef"`
	GPSMapDatum                      interface{} `json:"GPSMapDatum"`
	GPSMeasureMode                   interface{} `json:"GPSMeasureMode"`
	GPSProcessingMethod              interface{} `json:"GPSProcessingMethod"`
	GPSSatelites                     interface{} `json:"GPSSatelites"`
	GPSSpeed                         interface{} `json:"GPSSpeed"`
	GPSSpeedRef                      interface{} `json:"GPSSpeedRef"`
	GPSStatus                        interface{} `json:"GPSStatus"`
	GPSTimeStamp                     interface{} `json:"GPSTimeStamp"`
	GPSTrack                         interface{} `json:"GPSTrack"`
	GPSTrackRef                      interface{} `json:"GPSTrackRef"`
	GPSVersionID                     interface{} `json:"GPSVersionID"`
	GainControl                      interface{} `json:"GainControl"`
	ISOSpeedRatings                  interface{} `json:"ISOSpeedRatings"`
	ImageDescription                 interface{} `json:"ImageDescription"`
	ImageLength                      interface{} `json:"ImageLength"`
	ImageUniqueID                    interface{} `json:"ImageUniqueID"`
	ImageWidth                       interface{} `json:"ImageWidth"`
	InteroperabilityIFDPointer       interface{} `json:"InteroperabilityIFDPointer"`
	InteroperabilityIndex            interface{} `json:"InteroperabilityIndex"`
	LensMake                         interface{} `json:"LensMake"`
	LensModel                        interface{} `json:"LensModel"`
	LightSource                      interface{} `json:"LightSource"`
	Make                             interface{} `json:"Make"`
	MakerNote                        interface{} `json:"MakerNote"`
	MaxApertureValue                 interface{} `json:"MaxApertureValue"`
	MeteringMode                     interface{} `json:"MeteringMode"`
	Model                            interface{} `json:"Model"`
	OECF                             interface{} `json:"OECF"`
	Orientation                      interface{} `json:"Orientation"`
	PhotometricInterpretation        interface{} `json:"PhotometricInterpretation"`
	PixelXDimension                  interface{} `json:"PixelXDimension"`
	PixelYDimension                  interface{} `json:"PixelYDimension"`
	PlanarConfiguration              interface{} `json:"PlanarConfiguration"`
	RelatedSoundFile                 interface{} `json:"RelatedSoundFile"`
	ResolutionUnit                   interface{} `json:"ResolutionUnit"`
	SamplesPerPixel                  interface{} `json:"SamplesPerPixel"`
	Saturation                       interface{} `json:"Saturation"`
	SceneCaptureType                 interface{} `json:"SceneCaptureType"`
	SceneType                        interface{} `json:"SceneType"`
	SensingMethod                    interface{} `json:"SensingMethod"`
	Sharpness                        interface{} `json:"Sharpness"`
	ShutterSpeedValue                interface{} `json:"ShutterSpeedValue"`
	Software                         interface{} `json:"Software"`
	SpatialFrequencyResponse         interface{} `json:"SpatialFrequencyResponse"`
	SpectralSensitivity              interface{} `json:"SpectralSensitivity"`
	SubSecTime                       interface{} `json:"SubSecTime"`
	SubSecTimeDigitized              interface{} `json:"SubSecTimeDigitized"`
	SubSecTimeOriginal               interface{} `json:"SubSecTimeOriginal"`
	SubjectArea                      interface{} `json:"SubjectArea"`
	SubjectDistance                  interface{} `json:"SubjectDistance"`
	SubjectDistanceRange             interface{} `json:"SubjectDistanceRange"`
	SubjectLocation                  interface{} `json:"SubjectLocation"`
	ThumbJPEGInterchangeFormat       interface{} `json:"ThumbJPEGInterchangeFormat"`
	ThumbJPEGInterchangeFormatLength interface{} `json:"ThumbJPEGInterchangeFormatLength"`
	UserComment                      interface{} `json:"UserComment"`
	WhiteBalance                     interface{} `json:"WhiteBalance"`
	XPAuthor                         interface{} `json:"XPAuthor"`
	XPComment                        interface{} `json:"XPComment"`
	XPKeywords                       interface{} `json:"XPKeywords"`
	XPSubject                        interface{} `json:"XPSubject"`
	XPTitle                          interface{} `json:"XPTitle"`
	XResolution                      interface{} `json:"XResolution"`
	YCbCrPositioning                 interface{} `json:"YCbCrPositioning"`
	YCbCrSubSampling                 interface{} `json:"YCbCrSubSampling"`
	YResolution                      interface{} `json:"YResolution"`
}

// exifContainer implements the Exif interface.
type exifContainer struct {
	*exif.Exif
}

// Get returns the ExifTag associated with the specified field name.
// If the field name is not found or there is an error retrieving the tag, an error is returned.
func (e *exifContainer) Get(fieldName exif.FieldName) (ExifTag, error) {
	return e.Exif.Get(fieldName)
}

// getExif extracts Exif metadata from an io.Reader and returns an ExifMetadata object.
// It decodes the Exif data from the input file and populates the ExifMetadata structure.
// If an error occurs during decoding, an error is returned.
func getExif(file io.Reader) (*ExifMetadata, error) {
	// Decode the Exif data from the input file.
	e, err := exif.Decode(file)
	if err != nil {
		return nil, err
	}

	// Create an empty ExifMetadata structure.
	exifMetadata := ExifMetadata{}

	// Populate the ExifMetadata structure with data from the Exif container.
	setExif(&exifMetadata, &exifContainer{e})

	// Return the populated ExifMetadata.
	return &exifMetadata, nil
}

// getExifApertureValue retrieves the Aperture Value Exif tag from the provided Exif object.
// It returns the Aperture Value as an interface{} type, and an error if the tag retrieval fails.
func getExifApertureValue(e Exif) (any, error) {
	var apertureValue any
	tag, err := e.Get(exif.ApertureValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &apertureValue, nil
}

// getExifArtist retrieves the Artist Exif tag from the provided Exif object.
// It returns the Artist as an interface{} type, and an error if the tag retrieval fails.
func getExifArtist(e Exif) (any, error) {
	var artist any
	tag, err := e.Get(exif.Artist)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &artist, nil
}

// getExifBitsPerSample retrieves the Bits Per Sample Exif tag from the provided Exif object.
// It returns the Bits Per Sample as an interface{} type, and an error if the tag retrieval fails.
func getExifBitsPerSample(e Exif) (any, error) {
	var bitsPerSample any
	tag, err := e.Get(exif.BitsPerSample)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &bitsPerSample, nil
}

// getExifBrightnessValue retrieves the Brightness Value Exif tag from the provided Exif object.
// It returns the Brightness Value as a *string, and an error if the tag retrieval or conversion fails.
func getExifBrightnessValue(e Exif) (*string, error) {
	var brightnessValue string
	tag, err := e.Get(exif.BrightnessValue)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	brightnessValue = strings.Trim(tag.String(), "\"")
	return &brightnessValue, nil
}

// getExifCFAPattern retrieves the Color Filter Array (CFA) Pattern Exif tag from the provided Exif object.
// It returns the CFA Pattern as an interface{} type, and an error if the tag retrieval fails.
func getExifCFAPattern(e Exif) (any, error) {
	var cFAPattern any
	tag, err := e.Get(exif.CFAPattern)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &cFAPattern, nil
}

// getExifColorSpace retrieves the Color Space Exif tag from the provided Exif object.
// It returns the Color Space as an *int, and an error if the tag retrieval or conversion fails.
func getExifColorSpace(e Exif) (*int, error) {
	var colorSpace int
	tag, err := e.Get(exif.ColorSpace)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	colorSpace, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &colorSpace, nil
}

// getExifComponentsConfiguration retrieves the Components Configuration Exif tag from the provided Exif object.
// It returns the Components Configuration as a *string, and an error if the tag retrieval fails.
func getExifComponentsConfiguration(e Exif) (*string, error) {
	var componentsConfiguration string
	tag, err := e.Get(exif.ComponentsConfiguration)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	componentsConfiguration = strings.Trim(tag.String(), "\"")
	return &componentsConfiguration, nil
}

// getExifCompressedBitsPerPixel retrieves the Compressed Bits Per Pixel Exif tag from the provided Exif object.
// It returns the Compressed Bits Per Pixel as a *string, and an error if the tag retrieval fails.
func getExifCompressedBitsPerPixel(e Exif) (*string, error) {
	var compressedBitsPerPixel string
	tag, err := e.Get(exif.CompressedBitsPerPixel)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	compressedBitsPerPixel = strings.Trim(tag.String(), "\"")
	return &compressedBitsPerPixel, nil
}

// getExifCompression retrieves the Compression Exif tag from the provided Exif object.
// It returns the Compression as an interface{} type, and an error if the tag retrieval fails.
func getExifCompression(e Exif) (interface{}, error) {
	var compression interface{}
	tag, err := e.Get(exif.Compression)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return compression, nil
}

// getExifContrast retrieves the Contrast Exif tag from the provided Exif object.
// It returns the Contrast as an *int, and an error if the tag retrieval or conversion fails.
func getExifContrast(e Exif) (*int, error) {
	var contrast int
	tag, err := e.Get(exif.Contrast)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	contrast, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &contrast, nil
}

// getExifCopyright retrieves the Copyright Exif tag from the provided Exif object.
// It returns the Copyright as an interface{} type, and an error if the tag retrieval fails.
func getExifCopyright(e Exif) (interface{}, error) {
	var copyright interface{}
	tag, err := e.Get(exif.Copyright)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return copyright, nil
}

// getExifCustomRendered retrieves the Custom Rendered Exif tag from the provided Exif object.
// It returns the Custom Rendered as an *int, and an error if the tag retrieval or conversion fails.
func getExifCustomRendered(e Exif) (*int, error) {
	var customRendered int
	tag, err := e.Get(exif.CustomRendered)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	customRendered, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &customRendered, nil
}

// getExifDateTime retrieves the Date and Time Exif tag from the provided Exif object.
// It returns the Date and Time as a *string, and an error if the tag retrieval or conversion fails.
func getExifDateTime(e Exif) (*string, error) {
	var dateTime string
	tag, err := e.Get(exif.DateTime)
	if err != nil {
		return nil, err
	}
	// Get the string value of the tag.
	dateTime, err = tag.StringVal()
	if err != nil {
		return nil, err
	}
	return &dateTime, nil
}

// getExifDateTimeDigitized retrieves the Date and Time Digitized Exif tag from the provided Exif object.
// It returns the Date and Time Digitized as a *string, and an error if the tag retrieval or conversion fails.
func getExifDateTimeDigitized(e Exif) (*string, error) {
	var dateTimeDigitized string
	tag, err := e.Get(exif.DateTimeDigitized)
	if err != nil {
		return nil, err
	}
	// Get the string value of the tag.
	dateTimeDigitized, err = tag.StringVal()
	if err != nil {
		return nil, err
	}
	return &dateTimeDigitized, nil
}

// getExifDateTimeOriginal retrieves the Original Date and Time Exif tag from the provided Exif object.
// It returns the Original Date and Time as a *string, and an error if the tag retrieval or conversion fails.
func getExifDateTimeOriginal(e Exif) (*string, error) {
	var dateTimeOriginal string
	tag, err := e.Get(exif.DateTimeOriginal)
	if err != nil {
		return nil, err
	}
	// Get the string value of the tag.
	dateTimeOriginal, err = tag.StringVal()
	if err != nil {
		return nil, err
	}
	return &dateTimeOriginal, nil
}

// getExifDeviceSettingDescription retrieves the Device Setting Description Exif tag from the provided Exif object.
// It returns the Device Setting Description as an interface{} type, and an error if the tag retrieval fails.
func getExifDeviceSettingDescription(e Exif) (interface{}, error) {
	var deviceSettingDescription interface{}
	tag, err := e.Get(exif.DeviceSettingDescription)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return deviceSettingDescription, nil
}

// getExifDigitalZoomRatio retrieves the Digital Zoom Ratio Exif tag from the provided Exif object.
// It returns the Digital Zoom Ratio as a *string, and an error if the tag retrieval fails.
func getExifDigitalZoomRatio(e Exif) (*string, error) {
	var digitalZoomRatio string
	tag, err := e.Get(exif.DigitalZoomRatio)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	digitalZoomRatio = strings.Trim(tag.String(), "\"")
	return &digitalZoomRatio, nil
}

// getExifExifIFDPointer retrieves the Exif IFD Pointer Exif tag from the provided Exif object.
// It returns the Exif IFD Pointer as an *int, and an error if the tag retrieval or conversion fails.
func getExifExifIFDPointer(e Exif) (*int, error) {
	var exifIFDPointer int
	tag, err := e.Get(exif.ExifIFDPointer)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	exifIFDPointer, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &exifIFDPointer, nil
}

// getExifExifVersion retrieves the Exif Version Exif tag from the provided Exif object.
// It returns the Exif Version as a *string, and an error if the tag retrieval fails.
func getExifExifVersion(e Exif) (*string, error) {
	var exifVersion string
	tag, err := e.Get(exif.ExifVersion)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	exifVersion = strings.Trim(tag.String(), "\"")
	return &exifVersion, nil
}

// getExifExposureBiasValue retrieves the Exposure Bias Value Exif tag from the provided Exif object.
// It returns the Exposure Bias Value as a *string, and an error if the tag retrieval fails.
func getExifExposureBiasValue(e Exif) (*string, error) {
	var exposureBiasValue string
	tag, err := e.Get(exif.ExposureBiasValue)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	exposureBiasValue = strings.Trim(tag.String(), "\"")
	return &exposureBiasValue, nil
}

// getExifExposureIndex retrieves the Exposure Index Exif tag from the provided Exif object.
// It returns the Exposure Index as an interface{} type, and an error if the tag retrieval fails.
func getExifExposureIndex(e Exif) (interface{}, error) {
	var exposureIndex interface{}
	tag, err := e.Get(exif.ExposureIndex)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return exposureIndex, nil
}

// getExifExposureMode retrieves the Exposure Mode Exif tag from the provided Exif object.
// It returns the Exposure Mode as an *int, and an error if the tag retrieval or conversion fails.
func getExifExposureMode(e Exif) (*int, error) {
	var exposureMode int
	tag, err := e.Get(exif.ExposureMode)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	exposureMode, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &exposureMode, nil
}

// getExifExposureProgram retrieves the Exposure Program Exif tag from the provided Exif object.
// It returns the Exposure Program as an *int, and an error if the tag retrieval or conversion fails.
func getExifExposureProgram(e Exif) (*int, error) {
	var exposureProgram int
	tag, err := e.Get(exif.ExposureProgram)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	exposureProgram, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &exposureProgram, nil
}

// getExifExposureTime retrieves the Exposure Time Exif tag from the provided Exif object.
// It returns the Exposure Time as a *string, and an error if the tag retrieval fails.
func getExifExposureTime(e Exif) (*string, error) {
	var exposureTime string
	tag, err := e.Get(exif.ExposureTime)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	exposureTime = strings.Trim(tag.String(), "\"")
	return &exposureTime, nil
}

// getExifFNumber retrieves the FNumber Exif tag from the provided Exif object.
// It returns the FNumber as a *string, and an error if the tag retrieval fails.
func getExifFNumber(e Exif) (*string, error) {
	var fNumber string
	tag, err := e.Get(exif.FNumber)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	fNumber = strings.Trim(tag.String(), "\"")
	return &fNumber, nil
}

// getExifFileSource retrieves the File Source Exif tag from the provided Exif object.
// It returns the File Source as a *string, and an error if the tag retrieval fails.
func getExifFileSource(e Exif) (*string, error) {
	var fileSource string
	tag, err := e.Get(exif.FileSource)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	fileSource = strings.Trim(tag.String(), "\"")
	return &fileSource, nil
}

// getExifFlash retrieves the Flash Exif tag from the provided Exif object.
// It returns the Flash as an *int, and an error if the tag retrieval or conversion fails.
func getExifFlash(e Exif) (*int, error) {
	var flash int
	tag, err := e.Get(exif.Flash)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	flash, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &flash, nil
}

// getExifFlashEnergy retrieves the Flash Energy Exif tag from the provided Exif object.
// It returns the Flash Energy as an interface{} type, and an error if the tag retrieval fails.
func getExifFlashEnergy(e Exif) (interface{}, error) {
	var flashEnergy interface{}
	tag, err := e.Get(exif.FlashEnergy)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return flashEnergy, nil
}

// getExifFlashpixVersion retrieves the Flashpix Version Exif tag from the provided Exif object.
// It returns the Flashpix Version as a *string, and an error if the tag retrieval fails.
func getExifFlashpixVersion(e Exif) (*string, error) {
	var flashpixVersion string
	tag, err := e.Get(exif.FlashpixVersion)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	flashpixVersion = strings.Trim(tag.String(), "\"")
	return &flashpixVersion, nil
}

// getExifFocalLength retrieves the Focal Length Exif tag from the provided Exif object.
// It returns the Focal Length as a *string, and an error if the tag retrieval fails.
func getExifFocalLength(e Exif) (*string, error) {
	var focalLength string
	tag, err := e.Get(exif.FocalLength)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	focalLength = strings.Trim(tag.String(), "\"")
	return &focalLength, nil
}

// getExifFocalLengthIn35mmFilm retrieves the Focal Length in 35mm Film Exif tag from the provided Exif object.
// It returns the Focal Length in 35mm Film as an *int, and an error if the tag retrieval or conversion fails.
func getExifFocalLengthIn35mmFilm(e Exif) (*int, error) {
	var focalLengthIn35mmFilm int
	tag, err := e.Get(exif.FocalLengthIn35mmFilm)
	if err != nil {
		return nil, err
	}
	// Convert the tag value to an integer.
	focalLengthIn35mmFilm, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &focalLengthIn35mmFilm, nil
}

// getExifFocalPlaneResolutionUnit retrieves the Focal Plane Resolution Unit Exif tag from the provided Exif object.
// It returns the Focal Plane Resolution Unit as an interface{} type, and an error if the tag retrieval fails.
func getExifFocalPlaneResolutionUnit(e Exif) (interface{}, error) {
	var focalPlaneResolutionUnit interface{}
	tag, err := e.Get(exif.FocalPlaneResolutionUnit)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return focalPlaneResolutionUnit, nil
}

// getExifFocalPlaneXResolution retrieves the Focal Plane X Resolution Exif tag from the provided Exif object.
// It returns the Focal Plane X Resolution as an interface{} type, and an error if the tag retrieval fails.
func getExifFocalPlaneXResolution(e Exif) (interface{}, error) {
	var focalPlaneXResolution interface{}
	tag, err := e.Get(exif.FocalPlaneXResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return focalPlaneXResolution, nil
}

// getExifFocalPlaneYResolution retrieves the Focal Plane Y Resolution Exif tag from the provided Exif object.
// It returns the Focal Plane Y Resolution as an interface{} type, and an error if the tag retrieval fails.
func getExifFocalPlaneYResolution(e Exif) (interface{}, error) {
	var focalPlaneYResolution interface{}
	tag, err := e.Get(exif.FocalPlaneYResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return focalPlaneYResolution, nil
}

// getExifGPSAltitude retrieves the GPS Altitude Exif tag from the provided Exif object.
// It returns the GPS Altitude as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSAltitude(e Exif) (interface{}, error) {
	var gpsAltitude interface{}
	tag, err := e.Get(exif.GPSAltitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsAltitude, nil
}

// getExifGPSAltitudeRef retrieves the GPS Altitude Ref Exif tag from the provided Exif object.
// It returns the GPS Altitude Ref as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSAltitudeRef(e Exif) (interface{}, error) {
	var gpsAltitudeRef interface{}
	tag, err := e.Get(exif.GPSAltitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsAltitudeRef, nil
}

// getExifGPSAreaInformation retrieves the GPS Area Information Exif tag from the provided Exif object.
// It returns the GPS Area Information as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSAreaInformation(e Exif) (interface{}, error) {
	var gpsAreaInformation interface{}
	tag, err := e.Get(exif.GPSAreaInformation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsAreaInformation, nil
}

// getExifGPSDOP retrieves the GPS Dilution of Precision (DOP) Exif tag from the provided Exif object.
// It returns the GPS DOP as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDOP(e Exif) (interface{}, error) {
	var gpsDOP interface{}
	tag, err := e.Get(exif.GPSDOP)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDOP, nil
}

// getExifGPSDateStamp retrieves the GPS Date Stamp Exif tag from the provided Exif object.
// It returns the GPS Date Stamp as a *string, and an error if the tag retrieval fails.
func getExifGPSDateStamp(e Exif) (*string, error) {
	var gpsDateStamp string
	tag, err := e.Get(exif.GPSDateStamp)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	gpsDateStamp = strings.Trim(tag.String(), "\"")
	return &gpsDateStamp, nil
}

// getExifGPSDestBearing retrieves the GPS Destination Bearing Exif tag from the provided Exif object.
// It returns the GPS Destination Bearing as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestBearing(e Exif) (interface{}, error) {
	var gpsDestBearing interface{}
	tag, err := e.Get(exif.GPSDestBearing)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestBearing, nil
}

// getExifGPSDestBearingRef retrieves the GPS Destination Bearing Reference Exif tag from the provided Exif object.
// It returns the GPS Destination Bearing Reference as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestBearingRef(e Exif) (interface{}, error) {
	var gpsDestBearingRef interface{}
	tag, err := e.Get(exif.GPSDestBearingRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestBearingRef, nil
}

// getExifGPSDestDistance retrieves the GPS Destination Distance Exif tag from the provided Exif object.
// It returns the GPS Destination Distance as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestDistance(e Exif) (interface{}, error) {
	var gpsDestDistance interface{}
	tag, err := e.Get(exif.GPSDestDistance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestDistance, nil
}

// getExifGPSDestDistanceRef retrieves the GPS Destination Distance Reference Exif tag from the provided Exif object.
// It returns the GPS Destination Distance Reference as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestDistanceRef(e Exif) (interface{}, error) {
	var gpsDestDistanceRef interface{}
	tag, err := e.Get(exif.GPSDestDistanceRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestDistanceRef, nil
}

// getExifGPSDestLatitude retrieves the GPS Destination Latitude Exif tag from the provided Exif object.
// It returns the GPS Destination Latitude as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestLatitude(e Exif) (interface{}, error) {
	var gpsDestLatitude interface{}
	tag, err := e.Get(exif.GPSDestLatitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestLatitude, nil
}

// getExifGPSDestLatitudeRef retrieves the GPS Destination Latitude Reference Exif tag from the provided Exif object.
// It returns the GPS Destination Latitude Reference as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestLatitudeRef(e Exif) (interface{}, error) {
	var gpsDestLatitudeRef interface{}
	tag, err := e.Get(exif.GPSDestLatitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestLatitudeRef, nil
}

// getExifGPSDestLongitude retrieves the GPS Destination Longitude Exif tag from the provided Exif object.
// It returns the GPS Destination Longitude as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestLongitude(e Exif) (interface{}, error) {
	var gpsDestLongitude interface{}
	tag, err := e.Get(exif.GPSDestLongitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestLongitude, nil
}

// getExifGPSDestLongitudeRef retrieves the GPS Destination Longitude Reference Exif tag from the provided Exif object.
// It returns the GPS Destination Longitude Reference as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSDestLongitudeRef(e Exif) (interface{}, error) {
	var gpsDestLongitudeRef interface{}
	tag, err := e.Get(exif.GPSDestLongitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsDestLongitudeRef, nil
}

// getExifGPSDifferential retrieves the GPS Differential Correction Exif tag from the provided Exif object.
// It returns the GPS Differential Correction as an *int, and an error if the tag retrieval fails.
func getExifGPSDifferential(e Exif) (*int, error) {
	var gpsDifferential int
	tag, err := e.Get(exif.GPSDifferential)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDifferential, nil
}

// getExifGPSImgDirection retrieves the GPS Image Direction Exif tag from the provided Exif object.
// It returns the GPS Image Direction as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSImgDirection(e Exif) (interface{}, error) {
	var gpsImgDirection interface{}
	tag, err := e.Get(exif.GPSImgDirection)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsImgDirection, nil
}

// getExifGPSImgDirectionRef retrieves the GPS Image Direction Reference Exif tag from the provided Exif object.
// It returns the GPS Image Direction Reference as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSImgDirectionRef(e Exif) (interface{}, error) {
	var gpsImgDirectionRef interface{}
	tag, err := e.Get(exif.GPSImgDirectionRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsImgDirectionRef, nil
}

// getExifGPSInfoIFDPointer retrieves the GPS Information IFD Pointer Exif tag from the provided Exif object.
// It returns the GPS Information IFD Pointer as an *int, and an error if the tag retrieval fails.
func getExifGPSInfoIFDPointer(e Exif) (*int, error) {
	var gpsInfoIFDPointer int
	tag, err := e.Get(exif.GPSInfoIFDPointer)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSInfoIFDPointer, nil
}

// getExifGPSLatitude retrieves the GPS Latitude Exif tag from the provided Exif object.
// It returns the GPS Latitude as a []string, and an error if the tag retrieval or conversion fails.
func getExifGPSLatitude(e Exif) ([]string, error) {
	var gpsLatitude []string
	tag, err := e.Get(exif.GPSLatitude)
	if err != nil {
		return nil, err
	}
	// Unmarshal the tag's string representation into a []string.
	err = json.Unmarshal([]byte(tag.String()), &gpsLatitude)
	if err != nil {
		return nil, err
	}
	return gpsLatitude, nil
}

// getExifGPSLatitudeRef retrieves the GPS Latitude Reference Exif tag from the provided Exif object.
// It returns the GPS Latitude Reference as a *string, and an error if the tag retrieval fails.
func getExifGPSLatitudeRef(e Exif) (*string, error) {
	var gpsLatitudeRef string
	tag, err := e.Get(exif.GPSLatitudeRef)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	gpsLatitudeRef = strings.Trim(tag.String(), "\"")
	return &gpsLatitudeRef, nil
}

// getExifGPSLongitude retrieves the GPS Longitude Exif tag from the provided Exif object.
// It returns the GPS Longitude as a []string, and an error if the tag retrieval or conversion fails.
func getExifGPSLongitude(e Exif) ([]string, error) {
	var gpsLongitude []string
	tag, err := e.Get(exif.GPSLongitude)
	if err != nil {
		return nil, err
	}
	// Unmarshal the tag's string representation into a []string.
	err = json.Unmarshal([]byte(tag.String()), &gpsLongitude)
	if err != nil {
		return nil, err
	}
	return gpsLongitude, nil
}

// getExifGPSLongitudeRef retrieves the GPS Longitude Reference Exif tag from the provided Exif object.
// It returns the GPS Longitude Reference as a *string, and an error if the tag retrieval fails.
func getExifGPSLongitudeRef(e Exif) (*string, error) {
	var gpsLongitudeRef string
	tag, err := e.Get(exif.GPSLongitudeRef)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	gpsLongitudeRef = strings.Trim(tag.String(), "\"")
	return &gpsLongitudeRef, nil
}

// getExifGPSMapDatum retrieves the GPS Map Datum Exif tag from the provided Exif object.
// It returns the GPS Map Datum as a *string, and an error if the tag retrieval fails.
func getExifGPSMapDatum(e Exif) (*string, error) {
	var gpsMapDatum string
	tag, err := e.Get(exif.GPSMapDatum)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	gpsMapDatum = strings.Trim(tag.String(), "\"")
	return &gpsMapDatum, nil
}

// getExifGPSMeasureMode retrieves the GPS Measure Mode Exif tag from the provided Exif object.
// It returns the GPS Measure Mode as an *int, and an error if the tag retrieval or conversion fails.
func getExifGPSMeasureMode(e Exif) (*int, error) {
	var gpsMeasureMode int
	tag, err := e.Get(exif.GPSMeasureMode)
	if err != nil {
		return nil, err
	}
	// Convert the tag's string representation to an integer.
	gpsMeasureMode, err = strconv.Atoi(strings.Trim(tag.String(), "\""))
	if err != nil {
		return nil, err
	}
	return &gpsMeasureMode, nil
}

// getExifGPSProcessingMethod retrieves the GPS Processing Method Exif tag from the provided Exif object.
// It returns the GPS Processing Method as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSProcessingMethod(e Exif) (interface{}, error) {
	var gpsProcessingMethod interface{}
	tag, err := e.Get(exif.GPSProcessingMethod)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsProcessingMethod, nil
}

// getExifGPSSatelites retrieves the GPS Satellites Exif tag from the provided Exif object.
// It returns the GPS Satellites as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSSatelites(e Exif) (interface{}, error) {
	var gpsSatellites interface{}
	tag, err := e.Get(exif.GPSSatelites)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsSatellites, nil
}

// getExifGPSSpeed retrieves the GPS Speed Exif tag from the provided Exif object.
// It returns the GPS Speed as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSSpeed(e Exif) (interface{}, error) {
	var gpsSpeed interface{}
	tag, err := e.Get(exif.GPSSpeed)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsSpeed, nil
}

// getExifGPSSpeedRef retrieves the GPS Speed Reference Exif tag from the provided Exif object.
// It returns the GPS Speed Reference as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSSpeedRef(e Exif) (interface{}, error) {
	var gpsSpeedRef interface{}
	tag, err := e.Get(exif.GPSSpeedRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsSpeedRef, nil
}

// getExifGPSStatus retrieves the GPS Status Exif tag from the provided Exif object.
// It returns the GPS Status as a *string, and an error if the tag retrieval fails.
func getExifGPSStatus(e Exif) (*string, error) {
	var gpsStatus string
	tag, err := e.Get(exif.GPSStatus)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	gpsStatus = strings.Trim(tag.String(), "\"")
	return &gpsStatus, nil
}

// getExifGPSTimeStamp retrieves the GPS Time Stamp Exif tag from the provided Exif object.
// It returns the GPS Time Stamp as a []string, and an error if the tag retrieval or conversion fails.
func getExifGPSTimeStamp(e Exif) ([]string, error) {
	var gpsTimeStamp []string
	tag, err := e.Get(exif.GPSTimeStamp)
	if err != nil {
		return nil, err
	}
	// Unmarshal the tag's string representation into a []string.
	err = json.Unmarshal([]byte(tag.String()), &gpsTimeStamp)
	if err != nil {
		return nil, err
	}
	return gpsTimeStamp, nil
}

// getExifGPSTrack retrieves the GPS Track Exif tag from the provided Exif object.
// It returns the GPS Track as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSTrack(e Exif) (interface{}, error) {
	var gpsTrack interface{}
	tag, err := e.Get(exif.GPSTrack)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return gpsTrack, nil
}

// getExifGPSTrackRef retrieves the GPS Track Reference Exif tag from the provided Exif object.
// It returns the GPS Track Reference as an interface{} type, and an error if the tag retrieval fails.
func getExifGPSTrackRef(e Exif) (any, error) {
	var gpsTrackRef any
	tag, err := e.Get(exif.GPSTrackRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSTrackRef, nil
}

// getExifGPSVersionID retrieves the GPS Version ID Exif tag from the provided Exif object.
// It returns the GPS Version ID as a []string, and an error if the tag retrieval or conversion fails.
func getExifGPSVersionID(e Exif) ([]string, error) {
	var gpsVersionID []string
	tag, err := e.Get(exif.GPSVersionID)
	if err != nil {
		return nil, err
	}
	// Unmarshal the tag's string representation into a []string.
	err = json.Unmarshal([]byte(tag.String()), &gpsVersionID)
	if err != nil {
		return nil, err
	}
	return gpsVersionID, nil
}

// getExifGainControl retrieves the Gain Control Exif tag from the provided Exif object.
// It returns the Gain Control as an interface{} type, and an error if the tag retrieval fails.
func getExifGainControl(e Exif) (interface{}, error) {
	var gainControl interface{}
	tag, err := e.Get(exif.GainControl)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gainControl, nil
}

// getExifISOSpeedRatings retrieves the ISO Speed Ratings Exif tag from the provided Exif object.
// It returns the ISO Speed Ratings as an *int, and an error if the tag retrieval or conversion fails.
func getExifISOSpeedRatings(e Exif) (*int, error) {
	var iSOSpeedRatings int
	tag, err := e.Get(exif.ISOSpeedRatings)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &iSOSpeedRatings, nil
}

// getExifImageDescription retrieves the Image Description Exif tag from the provided Exif object.
// It returns the Image Description as an interface{} type, and an error if the tag retrieval fails.
func getExifImageDescription(e Exif) (interface{}, error) {
	var imageDescription interface{}
	tag, err := e.Get(exif.ImageDescription)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageDescription, nil
}

// getExifImageLength retrieves the Image Length Exif tag from the provided Exif object.
// It returns the Image Length as an interface{} type, and an error if the tag retrieval fails.
func getExifImageLength(e Exif) (interface{}, error) {
	var imageLength interface{}
	tag, err := e.Get(exif.ImageLength)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageLength, nil
}

// getExifImageUniqueID retrieves the Image Unique ID Exif tag from the provided Exif object.
// It returns the Image Unique ID as an interface{} type, and an error if the tag retrieval fails.
func getExifImageUniqueID(e Exif) (interface{}, error) {
	var imageUniqueID interface{}
	tag, err := e.Get(exif.ImageUniqueID)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageUniqueID, nil
}

// getExifImageWidth retrieves the Image Width Exif tag from the provided Exif object.
// It returns the Image Width as an interface{} type, and an error if the tag retrieval fails.
func getExifImageWidth(e Exif) (interface{}, error) {
	var imageWidth interface{}
	tag, err := e.Get(exif.ImageWidth)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageWidth, nil
}

// getExifInteroperabilityIFDPointer retrieves the Interoperability IFD Pointer Exif tag from the provided Exif object.
// It returns the Interoperability IFD Pointer as an *int, and an error if the tag retrieval or conversion fails.
func getExifInteroperabilityIFDPointer(e Exif) (*int, error) {
	var interoperabilityIFDPointer int
	tag, err := e.Get(exif.InteroperabilityIFDPointer)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &interoperabilityIFDPointer, nil
}

// getExifInteroperabilityIndex retrieves the Interoperability Index Exif tag from the provided Exif object.
// It returns the Interoperability Index as a *string, and an error if the tag retrieval fails.
func getExifInteroperabilityIndex(e Exif) (*string, error) {
	var interoperabilityIndex string
	tag, err := e.Get(exif.InteroperabilityIndex)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	interoperabilityIndex = strings.Trim(tag.String(), "\"")
	return &interoperabilityIndex, nil
}

// getExifLensMake retrieves the Lens Make Exif tag from the provided Exif object.
// It returns the Lens Make as an interface{} type, and an error if the tag retrieval fails.
func getExifLensMake(e Exif) (interface{}, error) {
	var lensMake interface{}
	tag, err := e.Get(exif.LensMake)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &lensMake, nil
}

// getExifLensModel retrieves the Lens Model Exif tag from the provided Exif object.
// It returns the Lens Model as a *string, and an error if the tag retrieval fails.
func getExifLensModel(e Exif) (*string, error) {
	var lensModel string
	tag, err := e.Get(exif.LensModel)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	lensModel = strings.Trim(tag.String(), "\"")
	return &lensModel, nil
}

// getExifLightSource retrieves the Light Source Exif tag from the provided Exif object.
// It returns the Light Source as an *int, and an error if the tag retrieval or conversion fails.
func getExifLightSource(e Exif) (*int, error) {
	var lightSource int
	tag, err := e.Get(exif.LightSource)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &lightSource, nil
}

// getExifMake retrieves the Make Exif tag from the provided Exif object.
// It returns the Make as a *string, and an error if the tag retrieval fails.
func getExifMake(e Exif) (*string, error) {
	var make string
	tag, err := e.Get(exif.Make)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	make = strings.Trim(tag.String(), "\"")
	return &make, nil
}

// getExifMakerNote retrieves the Maker Note Exif tag from the provided Exif object.
// It returns the Maker Note as a *string, and an error if the tag retrieval fails.
func getExifMakerNote(e Exif) (*string, error) {
	var makerNote string
	tag, err := e.Get(exif.MakerNote)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	makerNote = strings.Trim(tag.String(), "\"")
	return &makerNote, nil
}

// getExifMaxApertureValue retrieves the Max Aperture Value Exif tag from the provided Exif object.
// It returns the Max Aperture Value as a *string, and an error if the tag retrieval fails.
func getExifMaxApertureValue(e Exif) (*string, error) {
	var maxApertureValue string
	tag, err := e.Get(exif.MaxApertureValue)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	maxApertureValue = strings.Trim(tag.String(), "\"")
	return &maxApertureValue, nil
}

// getExifMeteringMode retrieves the Metering Mode Exif tag from the provided Exif object.
// It returns the Metering Mode as an *int, and an error if the tag retrieval or conversion fails.
func getExifMeteringMode(e Exif) (*int, error) {
	var meteringMode int
	tag, err := e.Get(exif.MeteringMode)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &meteringMode, nil
}

// getExifModel retrieves the Model Exif tag from the provided Exif object.
// It returns the Model as a *string, and an error if the tag retrieval fails.
func getExifModel(e Exif) (*string, error) {
	var model string
	tag, err := e.Get(exif.Model)
	if err != nil {
		return nil, err
	}
	// Trim surrounding quotes from the tag's string representation.
	model = strings.Trim(tag.String(), "\"")
	return &model, nil
}

// getExifOECF retrieves the OECF Exif tag from the provided Exif object.
// It returns the OECF as an interface{} type, and an error if the tag retrieval fails.
func getExifOECF(e Exif) (interface{}, error) {
	var oECF interface{}
	tag, err := e.Get(exif.OECF)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &oECF, nil
}

// getExifOrientation retrieves the Orientation Exif tag from the provided Exif object.
// It returns the Orientation as an *int, and an error if the tag retrieval or conversion fails.
func getExifOrientation(e Exif) (*int, error) {
	var orientation int
	tag, err := e.Get(exif.Orientation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &orientation, nil
}

// getExifPhotometricInterpretation retrieves the Photometric Interpretation Exif tag from the provided Exif object.
// It returns the Photometric Interpretation as an interface{} type, and an error if the tag retrieval fails.
func getExifPhotometricInterpretation(e Exif) (any, error) {
	var photometricInterpretation any
	tag, err := e.Get(exif.PhotometricInterpretation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &photometricInterpretation, nil
}

// getExifPixelXDimension retrieves the Pixel X Dimension Exif tag from the provided Exif object.
// It returns the Pixel X Dimension as an *int, and an error if the tag retrieval or conversion fails.
func getExifPixelXDimension(e Exif) (*int, error) {
	var pixelXDimension int
	tag, err := e.Get(exif.PixelXDimension)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &pixelXDimension, nil
}

// getExifPixelYDimension retrieves the Pixel Y Dimension Exif tag from the provided Exif object.
// It returns the Pixel Y Dimension as an *int, and an error if the tag retrieval or conversion fails.
func getExifPixelYDimension(e Exif) (*int, error) {
	var pixelYDimension int
	tag, err := e.Get(exif.PixelYDimension)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &pixelYDimension, nil
}

// getExifPlanarConfiguration retrieves the Planar Configuration Exif tag from the provided Exif object.
// It returns the Planar Configuration as an interface{} type, and an error if the tag retrieval fails.
func getExifPlanarConfiguration(e Exif) (interface{}, error) {
	var planarConfiguration interface{}
	tag, err := e.Get(exif.PlanarConfiguration)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &planarConfiguration, nil
}

// getExifRelatedSoundFile retrieves the Related Sound File Exif tag from the provided Exif object.
// It returns the Related Sound File as an interface{} type, and an error if the tag retrieval fails.
func getExifRelatedSoundFile(e Exif) (interface{}, error) {
	var relatedSoundFile interface{}
	tag, err := e.Get(exif.RelatedSoundFile)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &relatedSoundFile, nil
}

// getExifResolutionUnit retrieves the Resolution Unit Exif tag from the provided Exif object.
// It returns the Resolution Unit as an *int, and an error if the tag retrieval or conversion fails.
func getExifResolutionUnit(e Exif) (*int, error) {
	var resolutionUnit int
	tag, err := e.Get(exif.ResolutionUnit)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &resolutionUnit, nil
}

// getExifSamplesPerPixel retrieves the Samples Per Pixel Exif tag from the provided Exif object.
// It returns the Samples Per Pixel as an *int, and an error if the tag retrieval or conversion fails.
func getExifSamplesPerPixel(e Exif) (*int, error) {
	var samplesPerPixel int
	tag, err := e.Get(exif.SamplesPerPixel)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &samplesPerPixel, nil
}

// getExifSaturation retrieves the Saturation Exif tag from the provided Exif object.
// It returns the Saturation as an *int, and an error if the tag retrieval or conversion fails.
func getExifSaturation(e Exif) (*int, error) {
	var saturation int
	tag, err := e.Get(exif.Saturation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &saturation, nil
}

// getExifSceneCaptureType retrieves the Scene Capture Type Exif tag from the provided Exif object.
// It returns the Scene Capture Type as an *int, and an error if the tag retrieval or conversion fails.
func getExifSceneCaptureType(e Exif) (*int, error) {
	var sceneCaptureType int
	tag, err := e.Get(exif.SceneCaptureType)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sceneCaptureType, nil
}

// getExifSceneType retrieves the Scene Type Exif tag from the provided Exif object.
// It returns the Scene Type as an *string, and an error if the tag retrieval fails.
func getExifSceneType(e Exif) (*string, error) {
	var sceneType string
	tag, err := e.Get(exif.SceneType)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sceneType, nil
}

// getExifSensingMethod retrieves the Sensing Method Exif tag from the provided Exif object.
// It returns the Sensing Method as an interface{} type, and an error if the tag retrieval fails.
func getExifSensingMethod(e Exif) (interface{}, error) {
	var sensingMethod interface{}
	tag, err := e.Get(exif.SensingMethod)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sensingMethod, nil
}

// getExifSharpness retrieves the Sharpness Exif tag from the provided Exif object.
// It returns the Sharpness as an *int, and an error if the tag retrieval or conversion fails.
func getExifSharpness(e Exif) (*int, error) {
	var sharpness int
	tag, err := e.Get(exif.Sharpness)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sharpness, nil
}

// getExifShutterSpeedValue retrieves the Shutter Speed Value Exif tag from the provided Exif object.
// It returns the Shutter Speed Value as an interface{} type, and an error if the tag retrieval fails.
func getExifShutterSpeedValue(e Exif) (interface{}, error) {
	var shutterSpeedValue interface{}
	tag, err := e.Get(exif.ShutterSpeedValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &shutterSpeedValue, nil
}

// getExifSoftware retrieves the Software Exif tag from the provided Exif object.
// It returns the Software as an *string, and an error if the tag retrieval fails.
func getExifSoftware(e Exif) (*string, error) {
	var software string
	tag, err := e.Get(exif.Software)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &software, nil
}

// getExifSpatialFrequencyResponse retrieves the Spatial Frequency Response Exif tag from the provided Exif object.
// It returns the Spatial Frequency Response as an interface{} type, and an error if the tag retrieval fails.
func getExifSpatialFrequencyResponse(e Exif) (interface{}, error) {
	var spatialFrequencyResponse interface{}
	tag, err := e.Get(exif.SpatialFrequencyResponse)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &spatialFrequencyResponse, nil
}

// getExifSpectralSensitivity retrieves the Spectral Sensitivity Exif tag from the provided Exif object.
// It returns the Spectral Sensitivity as an interface{} type, and an error if the tag retrieval fails.
func getExifSpectralSensitivity(e Exif) (interface{}, error) {
	var spectralSensitivity interface{}
	tag, err := e.Get(exif.SpectralSensitivity)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &spectralSensitivity, nil
}

// getExifSubSecTime retrieves the SubSecTime Exif tag from the provided Exif object.
// It returns the SubSecTime as an *int, and an error if the tag retrieval or conversion fails.
func getExifSubSecTime(e Exif) (*int, error) {
	var subSecTime int
	tag, err := e.Get(exif.SubSecTime)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subSecTime, nil
}

// getExifSubSecTimeDigitized retrieves the SubSecTimeDigitized Exif tag from the provided Exif object.
// It returns the SubSecTimeDigitized as an *int, and an error if the tag retrieval or conversion fails.
func getExifSubSecTimeDigitized(e Exif) (*int, error) {
	var subSecTimeDigitized int
	tag, err := e.Get(exif.SubSecTimeDigitized)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subSecTimeDigitized, nil
}

// getExifSubSecTimeOriginal retrieves the SubSecTimeOriginal Exif tag from the provided Exif object.
// It returns the SubSecTimeOriginal as an *int, and an error if the tag retrieval or conversion fails.
func getExifSubSecTimeOriginal(e Exif) (*int, error) {
	var subSecTimeOriginal int
	tag, err := e.Get(exif.SubSecTimeOriginal)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subSecTimeOriginal, nil
}

// getExifSubjectArea retrieves the Subject Area Exif tag from the provided Exif object.
// It returns the Subject Area as an interface{} type, and an error if the tag retrieval fails.
func getExifSubjectArea(e Exif) (interface{}, error) {
	var subjectArea interface{}
	tag, err := e.Get(exif.SubjectArea)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectArea, nil
}

// getExifSubjectDistance retrieves the Subject Distance Exif tag from the provided Exif object.
// It returns the Subject Distance as an interface{} type, and an error if the tag retrieval fails.
func getExifSubjectDistance(e Exif) (interface{}, error) {
	var subjectDistance interface{}
	tag, err := e.Get(exif.SubjectDistance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectDistance, nil
}

// getExifSubjectDistanceRange retrieves the Subject Distance Range Exif tag from the provided Exif object.
// It returns the Subject Distance Range as an interface{} type, and an error if the tag retrieval fails.
func getExifSubjectDistanceRange(e Exif) (interface{}, error) {
	var subjectDistanceRange interface{}
	tag, err := e.Get(exif.SubjectDistanceRange)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectDistanceRange, nil
}

// getExifSubjectLocation retrieves the Subject Location Exif tag from the provided Exif object.
// It returns the Subject Location as an interface{} type, and an error if the tag retrieval fails.
func getExifSubjectLocation(e Exif) (interface{}, error) {
	var subjectLocation interface{}
	tag, err := e.Get(exif.SubjectLocation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectLocation, nil
}

// getExifThumbJPEGInterchangeFormat retrieves the ThumbJPEGInterchangeFormat Exif tag from the provided Exif object.
// It returns the ThumbJPEGInterchangeFormat as an *int, and an error if the tag retrieval or conversion fails.
func getExifThumbJPEGInterchangeFormat(e Exif) (*int, error) {
	var thumbJPEGInterchangeFormat int
	tag, err := e.Get(exif.ThumbJPEGInterchangeFormat)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &thumbJPEGInterchangeFormat, nil
}

// getExifThumbJPEGInterchangeFormatLength retrieves the ThumbJPEGInterchangeFormatLength Exif tag from the provided Exif object.
// It returns the ThumbJPEGInterchangeFormatLength as an *int, and an error if the tag retrieval or conversion fails.
func getExifThumbJPEGInterchangeFormatLength(e Exif) (*int, error) {
	var thumbJPEGInterchangeFormatLength int
	tag, err := e.Get(exif.ThumbJPEGInterchangeFormatLength)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &thumbJPEGInterchangeFormatLength, nil
}

// getExifUserComment retrieves the UserComment Exif tag from the provided Exif object.
// It returns the UserComment as a *string, and an error if the tag retrieval fails.
func getExifUserComment(e Exif) (*string, error) {
	var userComment string
	tag, err := e.Get(exif.UserComment)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &userComment, nil
}

// getExifWhiteBalance retrieves the WhiteBalance Exif tag from the provided Exif object.
// It returns the WhiteBalance as an *int, and an error if the tag retrieval or conversion fails.
func getExifWhiteBalance(e Exif) (*int, error) {
	var whiteBalance int
	tag, err := e.Get(exif.WhiteBalance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &whiteBalance, nil
}

// getExifXPAuthor retrieves the XPAuthor Exif tag from the provided Exif object.
// It returns the XPAuthor as an interface{}, and an error if the tag retrieval fails.
func getExifXPAuthor(e Exif) (interface{}, error) {
	var xPAuthor interface{}
	tag, err := e.Get(exif.XPAuthor)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPAuthor, nil
}

// getExifXPComment retrieves the XPComment Exif tag from the provided Exif object.
// It returns the XPComment as an interface{}, and an error if the tag retrieval fails.
func getExifXPComment(e Exif) (interface{}, error) {
	var xPComment interface{}
	tag, err := e.Get(exif.XPComment)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPComment, nil
}

// getExifXPKeywords retrieves the XPKeywords Exif tag from the provided Exif object.
// It returns the XPKeywords as an interface{}, and an error if the tag retrieval fails.
func getExifXPKeywords(e Exif) (interface{}, error) {
	var xPKeywords interface{}
	tag, err := e.Get(exif.XPKeywords)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPKeywords, nil
}

// getExifXPSubject retrieves the XPSubject Exif tag from the provided Exif object.
// It returns the XPSubject as an interface{}, and an error if the tag retrieval fails.
func getExifXPSubject(e Exif) (interface{}, error) {
	var xPSubject interface{}
	tag, err := e.Get(exif.XPSubject)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPSubject, nil
}

// getExifXPTitle retrieves the XPTitle Exif tag from the provided Exif object.
// It returns the XPTitle as an interface{}, and an error if the tag retrieval fails.
func getExifXPTitle(e Exif) (interface{}, error) {
	var xPTitle interface{}
	tag, err := e.Get(exif.XPTitle)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPTitle, nil
}

// getExifXResolution retrieves the XResolution Exif tag from the provided Exif object.
// It returns the XResolution as a *string, and an error if the tag retrieval fails.
func getExifXResolution(e Exif) (*string, error) {
	var xResolution string
	tag, err := e.Get(exif.XResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xResolution, nil
}

// getExifYCbCrPositioning retrieves the YCbCrPositioning Exif tag from the provided Exif object.
// It returns the YCbCrPositioning as an *int, and an error if the tag retrieval or conversion fails.
func getExifYCbCrPositioning(e Exif) (*int, error) {
	var yCbCrPositioning int
	tag, err := e.Get(exif.YCbCrPositioning)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &yCbCrPositioning, nil
}

// getExifYCbCrSubSampling retrieves the YCbCrSubSampling Exif tag from the provided Exif object.
// It returns the YCbCrSubSampling as an interface{}, and an error if the tag retrieval fails.
func getExifYCbCrSubSampling(e Exif) (interface{}, error) {
	var yCbCrSubSampling interface{}
	tag, err := e.Get(exif.YCbCrSubSampling)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &yCbCrSubSampling, nil
}

// getExifYResolution retrieves the YResolution Exif tag from the provided Exif object.
// It returns the YResolution as a *string, and an error if the tag retrieval fails.
func getExifYResolution(e Exif) (*string, error) {
	var yResolution string
	tag, err := e.Get(exif.YResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &yResolution, nil
}

// openExif opens an image file at the given filename and extracts Exif metadata.
// It returns a pointer to ExifMetadata and any error encountered during the process.
func openExif(filename string) (*ExifMetadata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return getExif(file)
}

// setExif sets various Exif metadata fields using the Exif struct and logs any encountered errors.
func setExif(exifMetadata *ExifMetadata, e Exif) {
	exifSetters := map[exif.FieldName]func(*ExifMetadata, Exif) error{
		exif.ApertureValue:                    setExifMetadataApertureValue,
		exif.Artist:                           setExifMetadataArtist,
		exif.BitsPerSample:                    setExifMetadataBitsPerSample,
		exif.BrightnessValue:                  setExifMetadataBrightnessValue,
		exif.CFAPattern:                       setExifMetadataCFAPattern,
		exif.ColorSpace:                       setExifMetadataColorSpace,
		exif.ComponentsConfiguration:          setExifMetadataComponentsConfiguration,
		exif.CompressedBitsPerPixel:           setExifMetadataCompressedBitsPerPixel,
		exif.Compression:                      setExifMetadataCompression,
		exif.Contrast:                         setExifMetadataContrast,
		exif.Copyright:                        setExifMetadataCopyright,
		exif.CustomRendered:                   setExifMetadataCustomRendered,
		exif.DateTime:                         setExifMetadataDateTime,
		exif.DateTimeDigitized:                setExifMetadataDateTimeDigitized,
		exif.DateTimeOriginal:                 setExifMetadataDateTimeOriginal,
		exif.DeviceSettingDescription:         setExifMetadataDeviceSettingDescription,
		exif.DigitalZoomRatio:                 setExifMetadataDigitalZoomRatio,
		exif.ExifIFDPointer:                   setExifMetadataExifIFDPointer,
		exif.ExifVersion:                      setExifMetadataExifVersion,
		exif.ExposureBiasValue:                setExifMetadataExposureBiasValue,
		exif.ExposureIndex:                    setExifMetadataExposureIndex,
		exif.ExposureMode:                     setExifMetadataExposureMode,
		exif.ExposureProgram:                  setExifMetadataExposureProgram,
		exif.ExposureTime:                     setExifMetadataExposureTime,
		exif.FNumber:                          setExifMetadataFNumber,
		exif.FileSource:                       setExifMetadataFileSource,
		exif.Flash:                            setExifMetadataFlash,
		exif.FlashEnergy:                      setExifMetadataFlashEnergy,
		exif.FlashpixVersion:                  setExifMetadataFlashpixVersion,
		exif.FocalLength:                      setExifMetadataFocalLength,
		exif.FocalLengthIn35mmFilm:            setExifMetadataFocalLengthIn35mmFilm,
		exif.FocalPlaneResolutionUnit:         setExifMetadataFocalPlaneResolutionUnit,
		exif.FocalPlaneXResolution:            setExifMetadataFocalPlaneXResolution,
		exif.FocalPlaneYResolution:            setExifMetadataFocalPlaneYResolution,
		exif.GPSAltitude:                      setExifMetadataGPSAltitude,
		exif.GPSAltitudeRef:                   setExifMetadataGPSAltitudeRef,
		exif.GPSAreaInformation:               setExifMetadataGPSAreaInformation,
		exif.GPSDOP:                           setExifMetadataGPSDOP,
		exif.GPSDateStamp:                     setExifMetadataGPSDateStamp,
		exif.GPSDestBearing:                   setExifMetadataGPSDestBearing,
		exif.GPSDestBearingRef:                setExifMetadataGPSDestBearingRef,
		exif.GPSDestDistance:                  setExifMetadataGPSDestDistance,
		exif.GPSDestDistanceRef:               setExifMetadataGPSDestDistanceRef,
		exif.GPSDestLatitude:                  setExifMetadataGPSDestLatitude,
		exif.GPSDestLatitudeRef:               setExifMetadataGPSDestLatitudeRef,
		exif.GPSDestLongitude:                 setExifMetadataGPSDestLongitude,
		exif.GPSDestLongitudeRef:              setExifMetadataGPSDestLongitudeRef,
		exif.GPSDifferential:                  setExifMetadataGPSDifferential,
		exif.GPSImgDirection:                  setExifMetadataGPSImgDirection,
		exif.GPSImgDirectionRef:               setExifMetadataGPSImgDirectionRef,
		exif.GPSInfoIFDPointer:                setExifMetadataGPSInfoIFDPointer,
		exif.GPSLatitude:                      setExifMetadataGPSLatitude,
		exif.GPSLatitudeRef:                   setExifMetadataGPSLatitudeRef,
		exif.GPSLongitude:                     setExifMetadataGPSLongitude,
		exif.GPSLongitudeRef:                  setExifMetadataGPSLongitudeRef,
		exif.GPSMapDatum:                      setExifMetadataGPSMapDatum,
		exif.GPSMeasureMode:                   setExifMetadataGPSMeasureMode,
		exif.GPSProcessingMethod:              setExifMetadataGPSProcessingMethod,
		exif.GPSSatelites:                     setExifMetadataGPSSatelites,
		exif.GPSSpeed:                         setExifMetadataGPSSpeed,
		exif.GPSSpeedRef:                      setExifMetadataGPSSpeedRef,
		exif.GPSStatus:                        setExifMetadataGPSStatus,
		exif.GPSTimeStamp:                     setExifMetadataGPSTimeStamp,
		exif.GPSTrack:                         setExifMetadataGPSTrack,
		exif.GPSTrackRef:                      setExifMetadataGPSTrackRef,
		exif.GPSVersionID:                     setExifMetadataGPSVersionID,
		exif.GainControl:                      setExifMetadataGainControl,
		exif.ISOSpeedRatings:                  setExifMetadataISOSpeedRatings,
		exif.ImageDescription:                 setExifMetadataImageDescription,
		exif.ImageLength:                      setExifMetadataImageLength,
		exif.ImageUniqueID:                    setExifMetadataImageUniqueID,
		exif.ImageWidth:                       setExifMetadataImageWidth,
		exif.InteroperabilityIFDPointer:       setExifMetadataInteroperabilityIFDPointer,
		exif.InteroperabilityIndex:            setExifMetadataInteroperabilityIndex,
		exif.LensMake:                         setExifMetadataLensMake,
		exif.LensModel:                        setExifMetadataLensModel,
		exif.LightSource:                      setExifMetadataLightSource,
		exif.Make:                             setExifMetadataMake,
		exif.MakerNote:                        setExifMetadataMakerNote,
		exif.MaxApertureValue:                 setExifMetadataMaxApertureValue,
		exif.MeteringMode:                     setExifMetadataMeteringMode,
		exif.Model:                            setExifMetadataModel,
		exif.OECF:                             setExifMetadataOECF,
		exif.Orientation:                      setExifMetadataOrientation,
		exif.PhotometricInterpretation:        setExifMetadataPhotometricInterpretation,
		exif.PixelXDimension:                  setExifMetadataPixelXDimension,
		exif.PixelYDimension:                  setExifMetadataPixelYDimension,
		exif.PlanarConfiguration:              setExifMetadataPlanarConfiguration,
		exif.RelatedSoundFile:                 setExifMetadataRelatedSoundFile,
		exif.ResolutionUnit:                   setExifMetadataResolutionUnit,
		exif.SamplesPerPixel:                  setExifMetadataSamplesPerPixel,
		exif.Saturation:                       setExifMetadataSaturation,
		exif.SceneCaptureType:                 setExifMetadataSceneCaptureType,
		exif.SceneType:                        setExifMetadataSceneType,
		exif.SensingMethod:                    setExifMetadataSensingMethod,
		exif.Sharpness:                        setExifMetadataSharpness,
		exif.ShutterSpeedValue:                setExifMetadataShutterSpeedValue,
		exif.Software:                         setExifMetadataSoftware,
		exif.SpatialFrequencyResponse:         setExifMetadataSpatialFrequencyResponse,
		exif.SpectralSensitivity:              setExifMetadataSpectralSensitivity,
		exif.SubSecTime:                       setExifMetadataSubSecTime,
		exif.SubSecTimeDigitized:              setExifMetadataSubSecTimeDigitized,
		exif.SubSecTimeOriginal:               setExifMetadataSubSecTimeOriginal,
		exif.SubjectArea:                      setExifMetadataSubjectArea,
		exif.SubjectDistance:                  setExifMetadataSubjectDistance,
		exif.SubjectDistanceRange:             setExifMetadataSubjectDistanceRange,
		exif.SubjectLocation:                  setExifMetadataSubjectLocation,
		exif.ThumbJPEGInterchangeFormat:       setExifMetadataThumbJPEGInterchangeFormat,
		exif.ThumbJPEGInterchangeFormatLength: setExifMetadataThumbJPEGInterchangeFormatLength,
		exif.UserComment:                      setExifMetadataUserComment,
		exif.WhiteBalance:                     setExifMetadataWhiteBalance,
		exif.XPAuthor:                         setExifMetadataXPAuthor,
		exif.XPComment:                        setExifMetadataXPComment,
		exif.XPKeywords:                       setExifMetadataXPKeywords,
		exif.XPSubject:                        setExifMetadataXPSubject,
		exif.XPTitle:                          setExifMetadataXPTitle,
		exif.XResolution:                      setExifMetadataXResolution,
		exif.YCbCrPositioning:                 setExifMetadataYCbCrPositioning,
		exif.YCbCrSubSampling:                 setExifMetadataYCbCrSubSampling,
		exif.YResolution:                      setExifMetadataYResolution,
	}
	for exifFieldName, exifSetter := range exifSetters {
		if err := exifSetter(exifMetadata, e); err != nil {
			log.Printf("ExifMetadata.%s: Error=%s", exifFieldName, err)
		}
	}
}

func setExifMetadataApertureValue(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ApertureValue, err = getExifApertureValue(e)
	if err != nil {
		logSetExifMetadataError(exif.ApertureValue, err)
	}
}

func setExifMetadataArtist(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Artist, err = getExifArtist(e)
	if err != nil {
		logSetExifMetadataError(exif.Artist, err)
	}
}

func setExifMetadataBitsPerSample(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.BitsPerSample, err = getExifBitsPerSample(e)
	if err != nil {
		logSetExifMetadataError(exif.BitsPerSample, err)
	}
}

func setExifMetadataBrightnessValue(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.BrightnessValue, err = getExifBrightnessValue(e)
	if err != nil {
		logSetExifMetadataError(exif.BrightnessValue, err)
	}
}

func setExifMetadataCFAPattern(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.CFAPattern, err = getExifCFAPattern(e)
	if err != nil {
		logSetExifMetadataError(exif.CFAPattern, err)
	}
}

func setExifMetadataColorSpace(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ColorSpace, err = getExifColorSpace(e)
	if err != nil {
		logSetExifMetadataError(exif.ColorSpace, err)
	}
}

func setExifMetadataComponentsConfiguration(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ComponentsConfiguration, err = getExifComponentsConfiguration(e)
	if err != nil {
		logSetExifMetadataError(exif.ComponentsConfiguration, err)
	}
}

func setExifMetadataCompressedBitsPerPixel(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.CompressedBitsPerPixel, err = getExifCompressedBitsPerPixel(e)
	if err != nil {
		logSetExifMetadataError(exif.CompressedBitsPerPixel, err)
	}
}

func setExifMetadataCompression(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Compression, err = getExifCompression(e)
	if err != nil {
		logSetExifMetadataError(exif.Compression, err)
	}
}

func setExifMetadataContrast(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Contrast, err = getExifContrast(e)
	if err != nil {
		logSetExifMetadataError(exif.Contrast, err)
	}
}

func setExifMetadataCopyright(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Copyright, err = getExifCopyright(e)
	if err != nil {
		logSetExifMetadataError(exif.Copyright, err)
	}
}

func setExifMetadataCustomRendered(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.CustomRendered, err = getExifCustomRendered(e)
	if err != nil {
		logSetExifMetadataError(exif.CustomRendered, err)
	}
}

func setExifMetadataDateTime(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.DateTime, err = getExifDateTime(e)
	if err != nil {
		logSetExifMetadataError(exif.DateTime, err)
	}
}

func setExifMetadataDateTimeDigitized(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.DateTimeDigitized, err = getExifDateTimeDigitized(e)
	if err != nil {
		logSetExifMetadataError(exif.DateTimeDigitized, err)
	}
}

func setExifMetadataDateTimeOriginal(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.DateTimeOriginal, err = getExifDateTimeOriginal(e)
	if err != nil {
		logSetExifMetadataError(exif.DateTimeOriginal, err)
	}
}

func setExifMetadataDeviceSettingDescription(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.DeviceSettingDescription, err = getExifDeviceSettingDescription(e)
	if err != nil {
		logSetExifMetadataError(exif.DeviceSettingDescription, err)
	}
}

func setExifMetadataDigitalZoomRatio(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.DigitalZoomRatio, err = getExifDigitalZoomRatio(e)
	if err != nil {
		logSetExifMetadataError(exif.DigitalZoomRatio, err)
	}
}

func setExifMetadataExifIFDPointer(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ExifIFDPointer, err = getExifExifIFDPointer(e)
	if err != nil {
		logSetExifMetadataError(exif.ExifIFDPointer, err)
	}
}

func setExifMetadataExifVersion(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ExifVersion, err = getExifExifVersion(e)
	if err != nil {
		logSetExifMetadataError(exif.ExifVersion, err)
	}
}

func setExifMetadataExposureBiasValue(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ExposureBiasValue, err = getExifExposureBiasValue(e)
	if err != nil {
		logSetExifMetadataError(exif.ExposureBiasValue, err)
	}
}

func setExifMetadataExposureIndex(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ExposureIndex, err = getExifExposureIndex(e)
	if err != nil {
		logSetExifMetadataError(exif.ExposureIndex, err)
	}
}

func setExifMetadataExposureMode(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ExposureMode, err = getExifExposureMode(e)
	if err != nil {
		logSetExifMetadataError(exif.ExposureMode, err)
	}
}

func setExifMetadataExposureProgram(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ExposureProgram, err = getExifExposureProgram(e)
	if err != nil {
		logSetExifMetadataError(exif.ExposureProgram, err)
	}
}

func setExifMetadataExposureTime(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ExposureTime, err = getExifExposureTime(e)
	if err != nil {
		logSetExifMetadataError(exif.ExposureTime, err)
	}
}

func setExifMetadataFNumber(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FNumber, err = getExifFNumber(e)
	if err != nil {
		logSetExifMetadataError(exif.FNumber, err)
	}
}

func setExifMetadataFileSource(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FileSource, err = getExifFileSource(e)
	if err != nil {
		logSetExifMetadataError(exif.FileSource, err)
	}
}

func setExifMetadataFlash(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Flash, err = getExifFlash(e)
	if err != nil {
		logSetExifMetadataError(exif.Flash, err)
	}
}

func setExifMetadataFlashEnergy(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FlashEnergy, err = getExifFlashEnergy(e)
	if err != nil {
		logSetExifMetadataError(exif.FlashEnergy, err)
	}
}

func setExifMetadataFlashpixVersion(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FlashpixVersion, err = getExifFlashpixVersion(e)
	if err != nil {
		logSetExifMetadataError(exif.FlashpixVersion, err)
	}
}

func setExifMetadataFocalLength(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FocalLength, err = getExifFocalLength(e)
	if err != nil {
		logSetExifMetadataError(exif.FocalLength, err)
	}
}

func setExifMetadataFocalLengthIn35mmFilm(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FocalLengthIn35mmFilm, err = getExifFocalLengthIn35mmFilm(e)
	if err != nil {
		logSetExifMetadataError(exif.FocalLengthIn35mmFilm, err)
	}
}

func setExifMetadataFocalPlaneResolutionUnit(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FocalPlaneResolutionUnit, err = getExifFocalPlaneResolutionUnit(e)
	if err != nil {
		logSetExifMetadataError(exif.FocalPlaneResolutionUnit, err)
	}
}

func setExifMetadataFocalPlaneXResolution(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FocalPlaneXResolution, err = getExifFocalPlaneXResolution(e)
	if err != nil {
		logSetExifMetadataError(exif.FocalPlaneXResolution, err)
	}
}

func setExifMetadataFocalPlaneYResolution(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.FocalPlaneYResolution, err = getExifFocalPlaneYResolution(e)
	if err != nil {
		logSetExifMetadataError(exif.FocalPlaneYResolution, err)
	}
}

func setExifMetadataGPSAltitude(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSAltitude, err = getExifGPSAltitude(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSAltitude, err)
	}
}

func setExifMetadataGPSAltitudeRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSAltitudeRef, err = getExifGPSAltitudeRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSAltitudeRef, err)
	}
}

func setExifMetadataGPSAreaInformation(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSAreaInformation, err = getExifGPSAreaInformation(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSAreaInformation, err)
	}
}

func setExifMetadataGPSDOP(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDOP, err = getExifGPSDOP(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDOP, err)
	}
}

func setExifMetadataGPSDateStamp(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDateStamp, err = getExifGPSDateStamp(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDateStamp, err)
	}
}

func setExifMetadataGPSDestBearing(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestBearing, err = getExifGPSDestBearing(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestBearing, err)
	}
}

func setExifMetadataGPSDestBearingRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestBearingRef, err = getExifGPSDestBearingRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestBearingRef, err)
	}
}

func setExifMetadataGPSDestDistance(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestDistance, err = getExifGPSDestDistance(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestDistance, err)
	}
}

func setExifMetadataGPSDestDistanceRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestDistanceRef, err = getExifGPSDestDistanceRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestDistanceRef, err)
	}
}

func setExifMetadataGPSDestLatitude(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestLatitude, err = getExifGPSDestLatitude(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestLatitude, err)
	}
}

func setExifMetadataGPSDestLatitudeRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestLatitudeRef, err = getExifGPSDestLatitudeRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestLatitudeRef, err)
	}
}

func setExifMetadataGPSDestLongitude(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestLongitude, err = getExifGPSDestLongitude(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestLongitude, err)
	}
}

func setExifMetadataGPSDestLongitudeRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDestLongitudeRef, err = getExifGPSDestLongitudeRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDestLongitudeRef, err)
	}
}

func setExifMetadataGPSDifferential(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSDifferential, err = getExifGPSDifferential(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSDifferential, err)
	}
}

func setExifMetadataGPSImgDirection(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSImgDirection, err = getExifGPSImgDirection(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSImgDirection, err)
	}
}

func setExifMetadataGPSImgDirectionRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSImgDirectionRef, err = getExifGPSImgDirectionRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSImgDirectionRef, err)
	}
}

func setExifMetadataGPSInfoIFDPointer(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSInfoIFDPointer, err = getExifGPSInfoIFDPointer(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSInfoIFDPointer, err)
	}
}

func setExifMetadataGPSLatitude(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSLatitude, err = getExifGPSLatitude(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSLatitude, err)
	}
}

func setExifMetadataGPSLatitudeRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSLatitudeRef, err = getExifGPSLatitudeRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSLatitudeRef, err)
	}
}

func setExifMetadataGPSLongitude(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSLongitude, err = getExifGPSLongitude(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSLongitude, err)
	}
}

func setExifMetadataGPSLongitudeRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSLongitudeRef, err = getExifGPSLongitudeRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSLongitudeRef, err)
	}
}

func setExifMetadataGPSMapDatum(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSMapDatum, err = getExifGPSMapDatum(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSMapDatum, err)
	}
}

func setExifMetadataGPSMeasureMode(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSMeasureMode, err = getExifGPSMeasureMode(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSMeasureMode, err)
	}
}

func setExifMetadataGPSProcessingMethod(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSProcessingMethod, err = getExifGPSProcessingMethod(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSProcessingMethod, err)
	}
}

func setExifMetadataGPSSatelites(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSSatelites, err = getExifGPSSatelites(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSSatelites, err)
	}
}

func setExifMetadataGPSSpeed(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSSpeed, err = getExifGPSSpeed(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSSpeed, err)
	}
}

func setExifMetadataGPSSpeedRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSSpeedRef, err = getExifGPSSpeedRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSSpeedRef, err)
	}
}

func setExifMetadataGPSStatus(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSStatus, err = getExifGPSStatus(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSStatus, err)
	}
}

func setExifMetadataGPSTimeStamp(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSTimeStamp, err = getExifGPSTimeStamp(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSTimeStamp, err)
	}
}

func setExifMetadataGPSTrack(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSTrack, err = getExifGPSTrack(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSTrack, err)
	}
}

func setExifMetadataGPSTrackRef(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSTrackRef, err = getExifGPSTrackRef(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSTrackRef, err)
	}
}

func setExifMetadataGPSVersionID(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GPSVersionID, err = getExifGPSVersionID(e)
	if err != nil {
		logSetExifMetadataError(exif.GPSVersionID, err)
	}
}

func setExifMetadataGainControl(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.GainControl, err = getExifGainControl(e)
	if err != nil {
		logSetExifMetadataError(exif.GainControl, err)
	}
}

func setExifMetadataISOSpeedRatings(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ISOSpeedRatings, err = getExifISOSpeedRatings(e)
	if err != nil {
		logSetExifMetadataError(exif.ISOSpeedRatings, err)
	}
}

func setExifMetadataImageDescription(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ImageDescription, err = getExifImageDescription(e)
	if err != nil {
		logSetExifMetadataError(exif.ImageDescription, err)
	}
}

func setExifMetadataImageLength(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ImageLength, err = getExifImageLength(e)
	if err != nil {
		logSetExifMetadataError(exif.ImageLength, err)
	}
}

func setExifMetadataImageUniqueID(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ImageUniqueID, err = getExifImageUniqueID(e)
	if err != nil {
		logSetExifMetadataError(exif.ImageUniqueID, err)
	}
}

func setExifMetadataImageWidth(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ImageWidth, err = getExifImageWidth(e)
	if err != nil {
		logSetExifMetadataError(exif.ImageWidth, err)
	}
}

func setExifMetadataInteroperabilityIFDPointer(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.InteroperabilityIFDPointer, err = getExifInteroperabilityIFDPointer(e)
	if err != nil {
		logSetExifMetadataError(exif.InteroperabilityIFDPointer, err)
	}
}

func setExifMetadataInteroperabilityIndex(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.InteroperabilityIndex, err = getExifInteroperabilityIndex(e)
	if err != nil {
		logSetExifMetadataError(exif.InteroperabilityIndex, err)
	}
}

func setExifMetadataLensMake(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.LensMake, err = getExifLensMake(e)
	if err != nil {
		logSetExifMetadataError(exif.LensMake, err)
	}
}

func setExifMetadataLensModel(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.LensModel, err = getExifLensModel(e)
	if err != nil {
		logSetExifMetadataError(exif.LensModel, err)
	}
}

func setExifMetadataLightSource(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.LightSource, err = getExifLightSource(e)
	if err != nil {
		logSetExifMetadataError(exif.LightSource, err)
	}
}

func setExifMetadataMake(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Make, err = getExifMake(e)
	if err != nil {
		logSetExifMetadataError(exif.Make, err)
	}
}

func setExifMetadataMakerNote(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.MakerNote, err = getExifMakerNote(e)
	if err != nil {
		logSetExifMetadataError(exif.MakerNote, err)
	}
}

func setExifMetadataMaxApertureValue(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.MaxApertureValue, err = getExifMaxApertureValue(e)
	if err != nil {
		logSetExifMetadataError(exif.MaxApertureValue, err)
	}
}

func setExifMetadataMeteringMode(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.MeteringMode, err = getExifMeteringMode(e)
	if err != nil {
		logSetExifMetadataError(exif.MeteringMode, err)
	}
}

func setExifMetadataModel(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Model, err = getExifModel(e)
	if err != nil {
		logSetExifMetadataError(exif.Model, err)
	}
}

func setExifMetadataOECF(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.OECF, err = getExifOECF(e)
	if err != nil {
		logSetExifMetadataError(exif.OECF, err)
	}
}

func setExifMetadataOrientation(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Orientation, err = getExifOrientation(e)
	if err != nil {
		logSetExifMetadataError(exif.Orientation, err)
	}
}

func setExifMetadataPhotometricInterpretation(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.PhotometricInterpretation, err = getExifPhotometricInterpretation(e)
	if err != nil {
		logSetExifMetadataError(exif.PhotometricInterpretation, err)
	}
}

func setExifMetadataPixelXDimension(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.PixelXDimension, err = getExifPixelXDimension(e)
	if err != nil {
		logSetExifMetadataError(exif.PixelXDimension, err)
	}
}

func setExifMetadataPixelYDimension(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.PixelYDimension, err = getExifPixelYDimension(e)
	if err != nil {
		logSetExifMetadataError(exif.PixelYDimension, err)
	}
}

func setExifMetadataPlanarConfiguration(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.PlanarConfiguration, err = getExifPlanarConfiguration(e)
	if err != nil {
		logSetExifMetadataError(exif.PlanarConfiguration, err)
	}
}

func setExifMetadataRelatedSoundFile(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.RelatedSoundFile, err = getExifRelatedSoundFile(e)
	if err != nil {
		logSetExifMetadataError(exif.RelatedSoundFile, err)
	}
}

func setExifMetadataResolutionUnit(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ResolutionUnit, err = getExifResolutionUnit(e)
	if err != nil {
		logSetExifMetadataError(exif.ResolutionUnit, err)
	}
}

func setExifMetadataSamplesPerPixel(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SamplesPerPixel, err = getExifSamplesPerPixel(e)
	if err != nil {
		logSetExifMetadataError(exif.SamplesPerPixel, err)
	}
}

func setExifMetadataSaturation(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Saturation, err = getExifSaturation(e)
	if err != nil {
		logSetExifMetadataError(exif.Saturation, err)
	}
}

func setExifMetadataSceneCaptureType(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SceneCaptureType, err = getExifSceneCaptureType(e)
	if err != nil {
		logSetExifMetadataError(exif.SceneCaptureType, err)
	}
}

func setExifMetadataSceneType(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SceneType, err = getExifSceneType(e)
	if err != nil {
		logSetExifMetadataError(exif.SceneType, err)
	}
}

func setExifMetadataSensingMethod(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SensingMethod, err = getExifSensingMethod(e)
	if err != nil {
		logSetExifMetadataError(exif.SensingMethod, err)
	}
}

func setExifMetadataSharpness(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Sharpness, err = getExifSharpness(e)
	if err != nil {
		logSetExifMetadataError(exif.Sharpness, err)
	}
}

func setExifMetadataShutterSpeedValue(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ShutterSpeedValue, err = getExifShutterSpeedValue(e)
	if err != nil {
		logSetExifMetadataError(exif.ShutterSpeedValue, err)
	}
}

func setExifMetadataSoftware(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.Software, err = getExifSoftware(e)
	if err != nil {
		logSetExifMetadataError(exif.Software, err)
	}
}

func setExifMetadataSpatialFrequencyResponse(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SpatialFrequencyResponse, err = getExifSpatialFrequencyResponse(e)
	if err != nil {
		logSetExifMetadataError(exif.SpatialFrequencyResponse, err)
	}
}

func setExifMetadataSpectralSensitivity(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SpectralSensitivity, err = getExifSpectralSensitivity(e)
	if err != nil {
		logSetExifMetadataError(exif.SpectralSensitivity, err)
	}
}

func setExifMetadataSubSecTime(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SubSecTime, err = getExifSubSecTime(e)
	if err != nil {
		logSetExifMetadataError(exif.SubSecTime, err)
	}
}

func setExifMetadataSubSecTimeDigitized(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SubSecTimeDigitized, err = getExifSubSecTimeDigitized(e)
	if err != nil {
		logSetExifMetadataError(exif.SubSecTimeDigitized, err)
	}
}

func setExifMetadataSubSecTimeOriginal(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SubSecTimeOriginal, err = getExifSubSecTimeOriginal(e)
	if err != nil {
		logSetExifMetadataError(exif.SubSecTimeOriginal, err)
	}
}

func setExifMetadataSubjectArea(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SubjectArea, err = getExifSubjectArea(e)
	if err != nil {
		logSetExifMetadataError(exif.SubjectArea, err)
	}
}

func setExifMetadataSubjectDistance(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SubjectDistance, err = getExifSubjectDistance(e)
	if err != nil {
		logSetExifMetadataError(exif.SubjectDistance, err)
	}
}

func setExifMetadataSubjectDistanceRange(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SubjectDistanceRange, err = getExifSubjectDistanceRange(e)
	if err != nil {
		logSetExifMetadataError(exif.SubjectDistanceRange, err)
	}
}

func setExifMetadataSubjectLocation(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.SubjectLocation, err = getExifSubjectLocation(e)
	if err != nil {
		logSetExifMetadataError(exif.SubjectLocation, err)
	}
}

func setExifMetadataThumbJPEGInterchangeFormat(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ThumbJPEGInterchangeFormat, err = getExifThumbJPEGInterchangeFormat(e)
	if err != nil {
		logSetExifMetadataError(exif.ThumbJPEGInterchangeFormat, err)
	}
}

func setExifMetadataThumbJPEGInterchangeFormatLength(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.ThumbJPEGInterchangeFormatLength, err = getExifThumbJPEGInterchangeFormatLength(e)
	if err != nil {
		logSetExifMetadataError(exif.ThumbJPEGInterchangeFormatLength, err)
	}
}

func setExifMetadataUserComment(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.UserComment, err = getExifUserComment(e)
	if err != nil {
		logSetExifMetadataError(exif.UserComment, err)
	}
}

func setExifMetadataWhiteBalance(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.WhiteBalance, err = getExifWhiteBalance(e)
	if err != nil {
		logSetExifMetadataError(exif.WhiteBalance, err)
	}
}

func setExifMetadataXPAuthor(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.XPAuthor, err = getExifXPAuthor(e)
	if err != nil {
		logSetExifMetadataError(exif.XPAuthor, err)
	}
}

func setExifMetadataXPComment(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.XPComment, err = getExifXPComment(e)
	if err != nil {
		logSetExifMetadataError(exif.XPComment, err)
	}
}

func setExifMetadataXPKeywords(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.XPKeywords, err = getExifXPKeywords(e)
	if err != nil {
		logSetExifMetadataError(exif.XPKeywords, err)
	}
}

func setExifMetadataXPSubject(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.XPSubject, err = getExifXPSubject(e)
	if err != nil {
		logSetExifMetadataError(exif.XPSubject, err)
	}
}

func setExifMetadataXPTitle(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.XPTitle, err = getExifXPTitle(e)
	if err != nil {
		logSetExifMetadataError(exif.XPTitle, err)
	}
}

func setExifMetadataXResolution(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.XResolution, err = getExifXResolution(e)
	if err != nil {
		logSetExifMetadataError(exif.XResolution, err)
	}
}

func setExifMetadataYCbCrPositioning(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.YCbCrPositioning, err = getExifYCbCrPositioning(e)
	if err != nil {
		logSetExifMetadataError(exif.YCbCrPositioning, err)
	}
}

func setExifMetadataYCbCrSubSampling(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.YCbCrSubSampling, err = getExifYCbCrSubSampling(e)
	if err != nil {
		logSetExifMetadataError(exif.YCbCrSubSampling, err)
	}
}

func setExifMetadataYResolution(exifMetadata *ExifMetadata, e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) {
	var err error
	exifMetadata.YResolution, err = getExifYResolution(e)
	if err != nil {
		logSetExifMetadataError(exif.YResolution, err)
	}
}
