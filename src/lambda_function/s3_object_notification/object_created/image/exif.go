package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
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
	Contrast                         *int        `json:"Contrast"`
	Copyright                        interface{} `json:"Copyright"`
	CustomRendered                   interface{} `json:"CustomRendered"`
	DateTime                         *string     `json:"DateTime"`
	DateTimeDigitized                *string     `json:"DateTimeDigitized"`
	DateTimeOriginal                 *string     `json:"DateTimeOriginal"`
	DeviceSettingDescription         interface{} `json:"DeviceSettingDescription"`
	DigitalZoomRatio                 *string     `json:"DigitalZoomRatio"`
	ExifIFDPointer                   *int        `json:"ExifIFDPointer"`
	ExifVersion                      *string     `json:"ExifVersion"`
	ExposureBiasValue                *string     `json:"ExposureBiasValue"`
	ExposureIndex                    interface{} `json:"ExposureIndex"`
	ExposureMode                     *int        `json:"ExposureMode"`
	ExposureProgram                  *int        `json:"ExposureProgram"`
	ExposureTime                     *string     `json:"ExposureTime"`
	FNumber                          *string     `json:"FNumber"`
	FileSource                       *string     `json:"FileSource"`
	Flash                            *int        `json:"Flash"`
	FlashEnergy                      interface{} `json:"FlashEnergy"`
	FlashpixVersion                  *string     `json:"FlashpixVersion"`
	FocalLength                      *string     `json:"FocalLength"`
	FocalLengthIn35mmFilm            *int        `json:"FocalLengthIn35mmFilm"`
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
	GPSLatitudeDegrees               *int        `json:"GPSLatitudeDegrees"`
	GPSLatitudeMinutes               *int        `json:"GPSLatitudeMinutes"`
	GPSLatitudeSeconds               *int        `json:"GPSLatitudeSeconds"`
	GPSLatitudeRef                   *string     `json:"GPSLatitudeRef"`
	GPSLongitudeDegrees              *int        `json:"GPSLongitudeDegrees"`
	GPSLongitudeMinutes              *int        `json:"GPSLongitudeMinutes"`
	GPSLongitudeSeconds              *int        `json:"GPSLongitudeSeconds"`
	GPSLongitudeRef                  *string     `json:"GPSLongitudeRef"`
	GPSMapDatum                      *string     `json:"GPSMapDatum"`
	GPSMeasureMode                   *int        `json:"GPSMeasureMode"`
	GPSProcessingMethod              interface{} `json:"GPSProcessingMethod"`
	GPSSatelites                     interface{} `json:"GPSSatelites"`
	GPSSpeed                         interface{} `json:"GPSSpeed"`
	GPSSpeedRef                      interface{} `json:"GPSSpeedRef"`
	GPSStatus                        *string     `json:"GPSStatus"`
	GPSTimeStampHours                *int        `json:"GPSTimeStampHours"`
	GPSTimeStampMinutes              *int        `json:"GPSTimeStampMinutes"`
	GPSTimeStampSeconds              *int        `json:"GPSTimeStampSeconds"`
	GPSTrack                         interface{} `json:"GPSTrack"`
	GPSTrackRef                      interface{} `json:"GPSTrackRef"`
	GPSVersionID                     *string     `json:"GPSVersionID"`
	GainControl                      interface{} `json:"GainControl"`
	ISOSpeedRatings                  *int        `json:"ISOSpeedRatings"`
	ImageDescription                 interface{} `json:"ImageDescription"`
	ImageLength                      interface{} `json:"ImageLength"`
	ImageUniqueID                    interface{} `json:"ImageUniqueID"`
	ImageWidth                       interface{} `json:"ImageWidth"`
	InteroperabilityIFDPointer       *int        `json:"InteroperabilityIFDPointer"`
	InteroperabilityIndex            *string     `json:"InteroperabilityIndex"`
	LensMake                         interface{} `json:"LensMake"`
	LensModel                        *string     `json:"LensModel"`
	LightSource                      *int        `json:"LightSource"`
	Make                             *string     `json:"Make"`
	MakerNote                        *string     `json:"MakerNote"`
	MaxApertureValue                 *string     `json:"MaxApertureValue"`
	MeteringMode                     *int        `json:"MeteringMode"`
	Model                            *string     `json:"Model"`
	OECF                             interface{} `json:"OECF"`
	Orientation                      *int        `json:"Orientation"`
	PhotometricInterpretation        interface{} `json:"PhotometricInterpretation"`
	PixelXDimension                  *int        `json:"PixelXDimension"`
	PixelYDimension                  *int        `json:"PixelYDimension"`
	PlanarConfiguration              interface{} `json:"PlanarConfiguration"`
	RelatedSoundFile                 interface{} `json:"RelatedSoundFile"`
	ResolutionUnit                   *int        `json:"ResolutionUnit"`
	SamplesPerPixel                  interface{} `json:"SamplesPerPixel"`
	Saturation                       *int        `json:"Saturation"`
	SceneCaptureType                 *int        `json:"SceneCaptureType"`
	SceneType                        *string     `json:"SceneType"`
	SensingMethod                    interface{} `json:"SensingMethod"`
	Sharpness                        *int        `json:"Sharpness"`
	ShutterSpeedValue                interface{} `json:"ShutterSpeedValue"`
	Software                         *string     `json:"Software"`
	SpatialFrequencyResponse         interface{} `json:"SpatialFrequencyResponse"`
	SpectralSensitivity              interface{} `json:"SpectralSensitivity"`
	SubSecTime                       *int        `json:"SubSecTime"`
	SubSecTimeDigitized              *int        `json:"SubSecTimeDigitized"`
	SubSecTimeOriginal               *int        `json:"SubSecTimeOriginal"`
	SubjectArea                      interface{} `json:"SubjectArea"`
	SubjectDistance                  interface{} `json:"SubjectDistance"`
	SubjectDistanceRange             interface{} `json:"SubjectDistanceRange"`
	SubjectLocation                  interface{} `json:"SubjectLocation"`
	ThumbJPEGInterchangeFormat       *int        `json:"ThumbJPEGInterchangeFormat"`
	ThumbJPEGInterchangeFormatLength *int        `json:"ThumbJPEGInterchangeFormatLength"`
	UserComment                      *string     `json:"UserComment"`
	WhiteBalance                     *int        `json:"WhiteBalance"`
	XPAuthor                         interface{} `json:"XPAuthor"`
	XPComment                        interface{} `json:"XPComment"`
	XPKeywords                       interface{} `json:"XPKeywords"`
	XPSubject                        interface{} `json:"XPSubject"`
	XPTitle                          interface{} `json:"XPTitle"`
	XResolution                      *string     `json:"XResolution"`
	YCbCrPositioning                 *int        `json:"YCbCrPositioning"`
	YCbCrSubSampling                 interface{} `json:"YCbCrSubSampling"`
	YResolution                      *string     `json:"YResolution"`
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
	gpsDifferential, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &gpsDifferential, nil
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
	gpsInfoIFDPointer, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &gpsInfoIFDPointer, nil
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
	return &gpsTrackRef, nil
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
	iSOSpeedRatings, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	interoperabilityIFDPointer, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	lightSource, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	meteringMode, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	orientation, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	pixelXDimension, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	pixelYDimension, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	resolutionUnit, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	samplesPerPixel, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	saturation, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	sceneCaptureType, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	sceneType = strings.Trim(tag.String(), "\"")
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
	sharpness, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	software = strings.Trim(tag.String(), "\"")
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

	subSecTime, err = strconv.Atoi(strings.ReplaceAll(strings.Trim(tag.String(), "\""), "\"", ""))
	if err != nil {
		return nil, err
	}
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
	subSecTimeDigitized, err = strconv.Atoi(strings.ReplaceAll(strings.Trim(tag.String(), "\""), "\"", ""))
	if err != nil {
		return nil, err
	}
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
	subSecTimeOriginal, err = strconv.Atoi(strings.ReplaceAll(strings.Trim(tag.String(), "\""), "\"", ""))
	if err != nil {
		return nil, err
	}
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
	thumbJPEGInterchangeFormat, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	thumbJPEGInterchangeFormatLength, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	userComment = strings.Trim(tag.String(), "\"")
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
	whiteBalance, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	xResolution = strings.Trim(tag.String(), "\"")
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
	yCbCrPositioning, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
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
	yResolution = strings.Trim(tag.String(), "\"")
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

// setExifMetadataApertureValue sets the ExifMetadata.ApertureValue field.
func setExifMetadataApertureValue(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ApertureValue, err = getExifApertureValue(e)
	return err
}

// setExifMetadataArtist sets the ExifMetadata.Artist field.
func setExifMetadataArtist(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Artist, err = getExifArtist(e)
	return err
}

// setExifMetadataBitsPerSample sets the ExifMetadata.BitsPerSample field.
func setExifMetadataBitsPerSample(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.BitsPerSample, err = getExifBitsPerSample(e)
	return err
}

// setExifMetadataBrightnessValue sets the ExifMetadata.BrightnessValue field.
func setExifMetadataBrightnessValue(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.BrightnessValue, err = getExifBrightnessValue(e)
	return err
}

// setExifMetadataCFAPattern sets the ExifMetadata.CFAPattern field.
func setExifMetadataCFAPattern(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.CFAPattern, err = getExifCFAPattern(e)
	return err
}

// setExifMetadataColorSpace sets the ExifMetadata.ColorSpace field.
func setExifMetadataColorSpace(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ColorSpace, err = getExifColorSpace(e)
	return err
}

// setExifMetadataComponentsConfiguration sets the ExifMetadata.ComponentsConfiguration field.
func setExifMetadataComponentsConfiguration(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ComponentsConfiguration, err = getExifComponentsConfiguration(e)
	return err
}

// setExifMetadataCompressedBitsPerPixel sets the ExifMetadata.CompressedBitsPerPixel field.
func setExifMetadataCompressedBitsPerPixel(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.CompressedBitsPerPixel, err = getExifCompressedBitsPerPixel(e)
	return err
}

// setExifMetadataCompression sets the ExifMetadata.Compression field.
func setExifMetadataCompression(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Compression, err = getExifCompression(e)
	return err
}

// setExifMetadataContrast sets the ExifMetadata.Contrast field.
func setExifMetadataContrast(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Contrast, err = getExifContrast(e)
	return err
}

// setExifMetadataCopyright sets the ExifMetadata.Copyright field.
func setExifMetadataCopyright(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Copyright, err = getExifCopyright(e)
	return err
}

// setExifMetadataCustomRendered sets the ExifMetadata.CustomRendered field.
func setExifMetadataCustomRendered(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.CustomRendered, err = getExifCustomRendered(e)
	return err
}

// setExifMetadataDateTime sets the ExifMetadata.DateTime field.
func setExifMetadataDateTime(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.DateTime, err = getExifDateTime(e)
	return err
}

// setExifMetadataDateTimeDigitized sets the ExifMetadata.DateTimeDigitized field.
func setExifMetadataDateTimeDigitized(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.DateTimeDigitized, err = getExifDateTimeDigitized(e)
	return err
}

// setExifMetadataDateTimeOriginal sets the ExifMetadata.DateTimeOriginal field.
func setExifMetadataDateTimeOriginal(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.DateTimeOriginal, err = getExifDateTimeOriginal(e)
	return err
}

// setExifMetadataDeviceSettingDescription sets the ExifMetadata.DeviceSettingDescription field.
func setExifMetadataDeviceSettingDescription(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.DeviceSettingDescription, err = getExifDeviceSettingDescription(e)
	return err
}

// setExifMetadataDigitalZoomRatio sets the ExifMetadata.DigitalZoomRatio field.
func setExifMetadataDigitalZoomRatio(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.DigitalZoomRatio, err = getExifDigitalZoomRatio(e)
	return err
}

// setExifMetadataExifIFDPointer sets the ExifMetadata.ExifIFDPointer field.
func setExifMetadataExifIFDPointer(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ExifIFDPointer, err = getExifExifIFDPointer(e)
	return err
}

// setExifMetadataExifVersion sets the ExifMetadata.ExifVersion field.
func setExifMetadataExifVersion(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ExifVersion, err = getExifExifVersion(e)
	return err
}

// setExifMetadataExposureBiasValue sets the ExifMetadata.ExposureBiasValue field.
func setExifMetadataExposureBiasValue(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ExposureBiasValue, err = getExifExposureBiasValue(e)
	return err
}

// setExifMetadataExposureIndex sets the ExifMetadata.ExposureIndex field.
func setExifMetadataExposureIndex(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ExposureIndex, err = getExifExposureIndex(e)
	return err
}

// setExifMetadataExposureMode sets the ExifMetadata.ExposureMode field.
func setExifMetadataExposureMode(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ExposureMode, err = getExifExposureMode(e)
	return err
}

// setExifMetadataExposureProgram sets the ExifMetadata.ExposureProgram field.
func setExifMetadataExposureProgram(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ExposureProgram, err = getExifExposureProgram(e)
	return err
}

// setExifMetadataExposureTime sets the ExifMetadata.ExposureTime field.
func setExifMetadataExposureTime(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ExposureTime, err = getExifExposureTime(e)
	return err
}

// setExifMetadataFNumber sets the ExifMetadata.FNumber field.
func setExifMetadataFNumber(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FNumber, err = getExifFNumber(e)
	return err
}

// setExifMetadataFileSource sets the ExifMetadata.FileSource field.
func setExifMetadataFileSource(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FileSource, err = getExifFileSource(e)
	return err
}

// setExifMetadataFlash sets the ExifMetadata.Flash field.
func setExifMetadataFlash(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Flash, err = getExifFlash(e)
	return err
}

// setExifMetadataFlashEnergy sets the ExifMetadata.FlashEnergy field.
func setExifMetadataFlashEnergy(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FlashEnergy, err = getExifFlashEnergy(e)
	return err
}

// setExifMetadataFlashpixVersion sets the ExifMetadata.FlashpixVersion field.
func setExifMetadataFlashpixVersion(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FlashpixVersion, err = getExifFlashpixVersion(e)
	return err
}

// setExifMetadataFocalLength sets the ExifMetadata.FocalLength field.
func setExifMetadataFocalLength(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FocalLength, err = getExifFocalLength(e)
	return err
}

// setExifMetadataFocalLengthIn35mmFilm sets the ExifMetadata.FocalLengthIn35mmFilm field.
func setExifMetadataFocalLengthIn35mmFilm(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FocalLengthIn35mmFilm, err = getExifFocalLengthIn35mmFilm(e)
	return err
}

// setExifMetadataFocalPlaneResolutionUnit sets the ExifMetadata.FocalPlaneResolutionUnit field.
func setExifMetadataFocalPlaneResolutionUnit(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FocalPlaneResolutionUnit, err = getExifFocalPlaneResolutionUnit(e)
	return err
}

// setExifMetadataFocalPlaneXResolution sets the ExifMetadata.FocalPlaneXResolution field.
func setExifMetadataFocalPlaneXResolution(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FocalPlaneXResolution, err = getExifFocalPlaneXResolution(e)
	return err
}

// setExifMetadataFocalPlaneYResolution sets the ExifMetadata.FocalPlaneYResolution field.
func setExifMetadataFocalPlaneYResolution(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.FocalPlaneYResolution, err = getExifFocalPlaneYResolution(e)
	return err
}

// setExifMetadataGPSAltitude sets the ExifMetadata.GPSAltitude field.
func setExifMetadataGPSAltitude(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSAltitude, err = getExifGPSAltitude(e)
	return err
}

// setExifMetadataGPSAltitudeRef sets the ExifMetadata.GPSAltitudeRef field.
func setExifMetadataGPSAltitudeRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSAltitudeRef, err = getExifGPSAltitudeRef(e)
	return err
}

// setExifMetadataGPSAreaInformation sets the ExifMetadata.GPSAreaInformation field.
func setExifMetadataGPSAreaInformation(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSAreaInformation, err = getExifGPSAreaInformation(e)
	return err
}

// setExifMetadataGPSDOP sets the ExifMetadata.GPSDOP field.
func setExifMetadataGPSDOP(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDOP, err = getExifGPSDOP(e)
	return err
}

// setExifMetadataGPSDateStamp sets the ExifMetadata.GPSDateStamp field.
func setExifMetadataGPSDateStamp(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDateStamp, err = getExifGPSDateStamp(e)
	return err
}

// setExifMetadataGPSDestBearing sets the ExifMetadata.GPSDestBearing field.
func setExifMetadataGPSDestBearing(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestBearing, err = getExifGPSDestBearing(e)
	return err
}

// setExifMetadataGPSDestBearingRef sets the ExifMetadata.GPSDestBearingRef field.
func setExifMetadataGPSDestBearingRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestBearingRef, err = getExifGPSDestBearingRef(e)
	return err
}

// setExifMetadataGPSDestDistance sets the ExifMetadata.GPSDestDistance field.
func setExifMetadataGPSDestDistance(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestDistance, err = getExifGPSDestDistance(e)
	return err
}

// setExifMetadataGPSDestDistanceRef sets the ExifMetadata.GPSDestDistanceRef field.
func setExifMetadataGPSDestDistanceRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestDistanceRef, err = getExifGPSDestDistanceRef(e)
	return err
}

// setExifMetadataGPSDestLatitude sets the ExifMetadata.GPSDestLatitude field.
func setExifMetadataGPSDestLatitude(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestLatitude, err = getExifGPSDestLatitude(e)
	return err
}

// setExifMetadataGPSDestLatitudeRef sets the ExifMetadata.GPSDestLatitudeRef field.
func setExifMetadataGPSDestLatitudeRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestLatitudeRef, err = getExifGPSDestLatitudeRef(e)
	return err
}

// setExifMetadataGPSDestLongitude sets the ExifMetadata.GPSDestLongitude field.
func setExifMetadataGPSDestLongitude(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestLongitude, err = getExifGPSDestLongitude(e)
	return err
}

// setExifMetadataGPSDestLongitudeRef sets the ExifMetadata.GPSDestLongitudeRef field.
func setExifMetadataGPSDestLongitudeRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDestLongitudeRef, err = getExifGPSDestLongitudeRef(e)
	return err
}

// setExifMetadataGPSDifferential sets the ExifMetadata.GPSDifferential field.
func setExifMetadataGPSDifferential(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSDifferential, err = getExifGPSDifferential(e)
	return err
}

// setExifMetadataGPSImgDirection sets the ExifMetadata.GPSImgDirection field.
func setExifMetadataGPSImgDirection(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSImgDirection, err = getExifGPSImgDirection(e)
	return err
}

// setExifMetadataGPSImgDirectionRef sets the ExifMetadata.GPSImgDirectionRef field.
func setExifMetadataGPSImgDirectionRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSImgDirectionRef, err = getExifGPSImgDirectionRef(e)
	return err
}

// setExifMetadataGPSInfoIFDPointer sets the ExifMetadata.GPSInfoIFDPointer field.
func setExifMetadataGPSInfoIFDPointer(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSInfoIFDPointer, err = getExifGPSInfoIFDPointer(e)
	return err
}

// setExifMetadataGPSLatitude sets the ExifMetadata.GPSLatitude field.
func setExifMetadataGPSLatitude(exifMetadata *ExifMetadata, e Exif) error {
	gpsLatitude, err := getExifGPSLatitude(e)
	if err != nil {
		return err
	}
	value := gpsLatitude[0]
	gpsLatitudeDegree, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	exifMetadata.GPSLatitudeDegrees = &gpsLatitudeDegree
	value = gpsLatitude[1]
	gpsLatitudeMinutes, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	exifMetadata.GPSLatitudeMinutes = &gpsLatitudeMinutes
	value = gpsLatitude[2]
	gpsLatitudeSeconds, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	gpsLatitudeSeconds = (gpsLatitudeSeconds / 100)
	exifMetadata.GPSLatitudeSeconds = &gpsLatitudeSeconds
	return nil
}

// setExifMetadataGPSLatitudeRef sets the ExifMetadata.GPSLatitudeRef field.
func setExifMetadataGPSLatitudeRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSLatitudeRef, err = getExifGPSLatitudeRef(e)
	return err
}

// setExifMetadataGPSLongitude sets the ExifMetadata.GPSLongitude field.
func setExifMetadataGPSLongitude(exifMetadata *ExifMetadata, e Exif) error {
	gpsLongitude, err := getExifGPSLongitude(e)
	if err != nil {
		return err
	}
	value := gpsLongitude[0]
	gpsLongitudeDegree, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	exifMetadata.GPSLongitudeDegrees = &gpsLongitudeDegree
	value = gpsLongitude[1]
	gpsLongitudeMinutes, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	exifMetadata.GPSLongitudeMinutes = &gpsLongitudeMinutes
	value = gpsLongitude[2]
	gpsLongitudeSeconds, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	gpsLongitudeSeconds = (gpsLongitudeSeconds / 100)
	exifMetadata.GPSLongitudeSeconds = &gpsLongitudeSeconds
	return nil
}

// setExifMetadataGPSLongitudeRef sets the ExifMetadata.GPSLongitudeRef field.
func setExifMetadataGPSLongitudeRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSLongitudeRef, err = getExifGPSLongitudeRef(e)
	return err
}

// setExifMetadataGPSMapDatum sets the ExifMetadata.GPSMapDatum field.
func setExifMetadataGPSMapDatum(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSMapDatum, err = getExifGPSMapDatum(e)
	return err
}

// setExifMetadataGPSMeasureMode sets the ExifMetadata.GPSMeasureMode field.
func setExifMetadataGPSMeasureMode(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSMeasureMode, err = getExifGPSMeasureMode(e)
	return err
}

// setExifMetadataGPSProcessingMethod sets the ExifMetadata.GPSProcessingMethod field.
func setExifMetadataGPSProcessingMethod(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSProcessingMethod, err = getExifGPSProcessingMethod(e)
	return err
}

// setExifMetadataGPSSatelites sets the ExifMetadata.GPSSatelites field.
func setExifMetadataGPSSatelites(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSSatelites, err = getExifGPSSatelites(e)
	return err
}

// setExifMetadataGPSSpeed sets the ExifMetadata.GPSSpeed field.
func setExifMetadataGPSSpeed(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSSpeed, err = getExifGPSSpeed(e)
	return err
}

// setExifMetadataGPSSpeedRef sets the ExifMetadata.GPSSpeedRef field.
func setExifMetadataGPSSpeedRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSSpeedRef, err = getExifGPSSpeedRef(e)
	return err
}

// setExifMetadataGPSStatus sets the ExifMetadata.GPSStatus field.
func setExifMetadataGPSStatus(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSStatus, err = getExifGPSStatus(e)
	return err
}

// setExifMetadataGPSTimeStamp sets the ExifMetadata.GPSTimeStamp field.
func setExifMetadataGPSTimeStamp(exifMetadata *ExifMetadata, e Exif) error {
	gpsTimeStamp, err := getExifGPSTimeStamp(e)
	if err != nil {
		return err
	}
	value := gpsTimeStamp[0]
	gpsTimeStampHours, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	exifMetadata.GPSTimeStampHours = &gpsTimeStampHours
	value = gpsTimeStamp[1]
	gpsTimeStampMinutes, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	exifMetadata.GPSTimeStampMinutes = &gpsTimeStampMinutes
	value = gpsTimeStamp[2]
	gpsTimeStampSeconds, err := strconv.Atoi(strings.Split(value, "/")[0])
	if err != nil {
		return err
	}
	gpsTimeStampSeconds = (gpsTimeStampSeconds / 100)
	exifMetadata.GPSTimeStampSeconds = &gpsTimeStampSeconds
	return nil
}

// setExifMetadataGPSTrack sets the ExifMetadata.GPSTrack field.
func setExifMetadataGPSTrack(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSTrack, err = getExifGPSTrack(e)
	return err
}

// setExifMetadataGPSTrackRef sets the ExifMetadata.GPSTrackRef field.
func setExifMetadataGPSTrackRef(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GPSTrackRef, err = getExifGPSTrackRef(e)
	return err
}

// setExifMetadataGPSVersionID sets the ExifMetadata.GPSVersionID field.
func setExifMetadataGPSVersionID(exifMetadata *ExifMetadata, e Exif) error {
	value, err := getExifGPSVersionID(e)
	if err != nil {
		return err
	}
	gpsVersionID := strings.Join(value, ".")
	exifMetadata.GPSVersionID = &gpsVersionID
	return nil
}

// setExifMetadataGainControl sets the ExifMetadata.GainControl field.
func setExifMetadataGainControl(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.GainControl, err = getExifGainControl(e)
	return err
}

// setExifMetadataISOSpeedRatings sets the ExifMetadata.ISOSpeedRatings field.
func setExifMetadataISOSpeedRatings(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ISOSpeedRatings, err = getExifISOSpeedRatings(e)
	return err
}

// setExifMetadataImageDescription sets the ExifMetadata.ImageDescription field.
func setExifMetadataImageDescription(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ImageDescription, err = getExifImageDescription(e)
	return err
}

// setExifMetadataImageLength sets the ExifMetadata.ImageLength field.
func setExifMetadataImageLength(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ImageLength, err = getExifImageLength(e)
	return err
}

// setExifMetadataImageUniqueID sets the ExifMetadata.ImageUniqueID field.
func setExifMetadataImageUniqueID(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ImageUniqueID, err = getExifImageUniqueID(e)
	return err
}

// setExifMetadataImageWidth sets the ExifMetadata.ImageWidth field.
func setExifMetadataImageWidth(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ImageWidth, err = getExifImageWidth(e)
	return err
}

// setExifMetadataInteroperabilityIFDPointer sets the ExifMetadata.InteroperabilityIFDPointer field.
func setExifMetadataInteroperabilityIFDPointer(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.InteroperabilityIFDPointer, err = getExifInteroperabilityIFDPointer(e)
	return err
}

// setExifMetadataInteroperabilityIndex sets the ExifMetadata.InteroperabilityIndex field.
func setExifMetadataInteroperabilityIndex(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.InteroperabilityIndex, err = getExifInteroperabilityIndex(e)
	return err
}

// setExifMetadataLensMake sets the ExifMetadata.LensMake field.
func setExifMetadataLensMake(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.LensMake, err = getExifLensMake(e)
	return err
}

// setExifMetadataLensModel sets the ExifMetadata.LensModel field.
func setExifMetadataLensModel(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.LensModel, err = getExifLensModel(e)
	return err
}

// setExifMetadataLightSource sets the ExifMetadata.LightSource field.
func setExifMetadataLightSource(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.LightSource, err = getExifLightSource(e)
	return err
}

// setExifMetadataMake sets the ExifMetadata.Make field.
func setExifMetadataMake(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Make, err = getExifMake(e)
	return err
}

// setExifMetadataMakerNote sets the ExifMetadata.MakerNote field.
func setExifMetadataMakerNote(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.MakerNote, err = getExifMakerNote(e)
	return err
}

// setExifMetadataMaxApertureValue sets the ExifMetadata.MaxApertureValue field.
func setExifMetadataMaxApertureValue(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.MaxApertureValue, err = getExifMaxApertureValue(e)
	return err
}

// setExifMetadataMeteringMode sets the ExifMetadata.MeteringMode field.
func setExifMetadataMeteringMode(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.MeteringMode, err = getExifMeteringMode(e)
	return err
}

// setExifMetadataModel sets the ExifMetadata.Model field.
func setExifMetadataModel(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Model, err = getExifModel(e)
	return err
}

// setExifMetadataOECF sets the ExifMetadata.OECF field.
func setExifMetadataOECF(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.OECF, err = getExifOECF(e)
	return err
}

// setExifMetadataOrientation sets the ExifMetadata.Orientation field.
func setExifMetadataOrientation(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Orientation, err = getExifOrientation(e)
	return err
}

// setExifMetadataPhotometricInterpretation sets the ExifMetadata.PhotometricInterpretation field.
func setExifMetadataPhotometricInterpretation(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.PhotometricInterpretation, err = getExifPhotometricInterpretation(e)
	return err
}

// setExifMetadataPixelXDimension sets the ExifMetadata.PixelXDimension field.
func setExifMetadataPixelXDimension(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.PixelXDimension, err = getExifPixelXDimension(e)
	return err
}

// setExifMetadataPixelYDimension sets the ExifMetadata.PixelYDimension field.
func setExifMetadataPixelYDimension(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.PixelYDimension, err = getExifPixelYDimension(e)
	return err
}

// setExifMetadataPlanarConfiguration sets the ExifMetadata.PlanarConfiguration field.
func setExifMetadataPlanarConfiguration(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.PlanarConfiguration, err = getExifPlanarConfiguration(e)
	return err
}

// setExifMetadataRelatedSoundFile sets the ExifMetadata.RelatedSoundFile field.
func setExifMetadataRelatedSoundFile(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.RelatedSoundFile, err = getExifRelatedSoundFile(e)
	return err
}

// setExifMetadataResolutionUnit sets the ExifMetadata.ResolutionUnit field.
func setExifMetadataResolutionUnit(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ResolutionUnit, err = getExifResolutionUnit(e)
	return err
}

// setExifMetadataSamplesPerPixel sets the ExifMetadata.SamplesPerPixel field.
func setExifMetadataSamplesPerPixel(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SamplesPerPixel, err = getExifSamplesPerPixel(e)
	return err
}

// setExifMetadataSaturation sets the ExifMetadata.Saturation field.
func setExifMetadataSaturation(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Saturation, err = getExifSaturation(e)
	return err
}

// setExifMetadataSceneCaptureType sets the ExifMetadata.SceneCaptureType field.
func setExifMetadataSceneCaptureType(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SceneCaptureType, err = getExifSceneCaptureType(e)
	return err
}

// setExifMetadataSceneType sets the ExifMetadata.SceneType field.
func setExifMetadataSceneType(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SceneType, err = getExifSceneType(e)
	return err
}

// setExifMetadataSensingMethod sets the ExifMetadata.SensingMethod field.
func setExifMetadataSensingMethod(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SensingMethod, err = getExifSensingMethod(e)
	return err
}

// setExifMetadataSharpness sets the ExifMetadata.Sharpness field.
func setExifMetadataSharpness(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Sharpness, err = getExifSharpness(e)
	return err
}

// setExifMetadataShutterSpeedValue sets the ExifMetadata.ShutterSpeedValue field.
func setExifMetadataShutterSpeedValue(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ShutterSpeedValue, err = getExifShutterSpeedValue(e)
	return err
}

// setExifMetadataSoftware sets the ExifMetadata.Software field.
func setExifMetadataSoftware(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.Software, err = getExifSoftware(e)
	return err
}

// setExifMetadataSpatialFrequencyResponse sets the ExifMetadata.SpatialFrequencyResponse field.
func setExifMetadataSpatialFrequencyResponse(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SpatialFrequencyResponse, err = getExifSpatialFrequencyResponse(e)
	return err
}

// setExifMetadataSpectralSensitivity sets the ExifMetadata.SpectralSensitivity field.
func setExifMetadataSpectralSensitivity(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SpectralSensitivity, err = getExifSpectralSensitivity(e)
	return err
}

// setExifMetadataSubSecTime sets the ExifMetadata.SubSecTime field.
func setExifMetadataSubSecTime(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SubSecTime, err = getExifSubSecTime(e)
	return err
}

// setExifMetadataSubSecTimeDigitized sets the ExifMetadata.SubSecTimeDigitized field.
func setExifMetadataSubSecTimeDigitized(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SubSecTimeDigitized, err = getExifSubSecTimeDigitized(e)
	return err
}

// setExifMetadataSubSecTimeOriginal sets the ExifMetadata.SubSecTimeOriginal field.
func setExifMetadataSubSecTimeOriginal(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SubSecTimeOriginal, err = getExifSubSecTimeOriginal(e)
	return err
}

// setExifMetadataSubjectArea sets the ExifMetadata.SubjectArea field.
func setExifMetadataSubjectArea(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SubjectArea, err = getExifSubjectArea(e)
	return err
}

// setExifMetadataSubjectDistance sets the ExifMetadata.SubjectDistance field.
func setExifMetadataSubjectDistance(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SubjectDistance, err = getExifSubjectDistance(e)
	return err
}

// setExifMetadataSubjectDistanceRange sets the ExifMetadata.SubjectDistanceRange field.
func setExifMetadataSubjectDistanceRange(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SubjectDistanceRange, err = getExifSubjectDistanceRange(e)
	return err
}

// setExifMetadataSubjectLocation sets the ExifMetadata.SubjectLocation field.
func setExifMetadataSubjectLocation(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.SubjectLocation, err = getExifSubjectLocation(e)
	return err
}

// setExifMetadataThumbJPEGInterchangeFormat sets the ExifMetadata.ThumbJPEGInterchangeFormat field.
func setExifMetadataThumbJPEGInterchangeFormat(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ThumbJPEGInterchangeFormat, err = getExifThumbJPEGInterchangeFormat(e)
	return err
}

// setExifMetadataThumbJPEGInterchangeFormatLength sets the ExifMetadata.ThumbJPEGInterchangeFormatLength field.
func setExifMetadataThumbJPEGInterchangeFormatLength(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.ThumbJPEGInterchangeFormatLength, err = getExifThumbJPEGInterchangeFormatLength(e)
	return err
}

// setExifMetadataUserComment sets the ExifMetadata.UserComment field.
func setExifMetadataUserComment(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.UserComment, err = getExifUserComment(e)
	return err
}

// setExifMetadataWhiteBalance sets the ExifMetadata.WhiteBalance field.
func setExifMetadataWhiteBalance(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.WhiteBalance, err = getExifWhiteBalance(e)
	return err
}

// setExifMetadataXPAuthor sets the ExifMetadata.XPAuthor field.
func setExifMetadataXPAuthor(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.XPAuthor, err = getExifXPAuthor(e)
	return err
}

// setExifMetadataXPComment sets the ExifMetadata.XPComment field.
func setExifMetadataXPComment(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.XPComment, err = getExifXPComment(e)
	return err
}

// setExifMetadataXPKeywords sets the ExifMetadata.XPKeywords field.
func setExifMetadataXPKeywords(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.XPKeywords, err = getExifXPKeywords(e)
	return err
}

// setExifMetadataXPSubject sets the ExifMetadata.XPSubject field.
func setExifMetadataXPSubject(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.XPSubject, err = getExifXPSubject(e)
	return err
}

// setExifMetadataXPTitle sets the ExifMetadata.XPTitle field.
func setExifMetadataXPTitle(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.XPTitle, err = getExifXPTitle(e)
	return err
}

// setExifMetadataXResolution sets the ExifMetadata.XResolution field.
func setExifMetadataXResolution(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.XResolution, err = getExifXResolution(e)
	return err
}

// setExifMetadataYCbCrPositioning sets the ExifMetadata.YCbCrPositioning field.
func setExifMetadataYCbCrPositioning(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.YCbCrPositioning, err = getExifYCbCrPositioning(e)
	return err
}

// setExifMetadataYCbCrSubSampling sets the ExifMetadata.YCbCrSubSampling field.
func setExifMetadataYCbCrSubSampling(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.YCbCrSubSampling, err = getExifYCbCrSubSampling(e)
	return err
}

// setExifMetadataYResolution sets the ExifMetadata.YResolution field.
func setExifMetadataYResolution(exifMetadata *ExifMetadata, e Exif) error {
	var err error
	exifMetadata.YResolution, err = getExifYResolution(e)
	return err
}
