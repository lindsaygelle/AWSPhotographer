package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

type Exif interface {
	Get(exif.FieldName) (ExifTag, error)
}

type ExifTag interface {
	Int(int) (int, error)
	String() string
	StringVal() (string, error)
}

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

type exifContainer struct {
	*exif.Exif
}

func (e *exifContainer) Get(fieldName exif.FieldName) (ExifTag, error) {
	return e.Exif.Get(fieldName)
}

func getExif(filename string) *ExifMetadata {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	e, err := exif.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	exifMetadata := ExifMetadata{}
	setExif(&exifMetadata, &exifContainer{e})
	return &exifMetadata
}

func getExifApertureValue(e Exif) (any, error) {
	var apertureValue any
	tag, err := e.Get(exif.ApertureValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &apertureValue, nil
}

func getExifArtist(e Exif) (any, error) {
	var artist any
	tag, err := e.Get(exif.Artist)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &artist, nil
}

func getExifBitsPerSample(e Exif) (any, error) {
	var bitsPerSample any
	tag, err := e.Get(exif.BitsPerSample)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &bitsPerSample, nil
}

func getExifBrightnessValue(e Exif) (*string, error) {
	var brightnessValue string
	tag, err := e.Get(exif.BrightnessValue)
	if err != nil {
		return nil, err
	}
	brightnessValue = strings.Trim(tag.String(), "\"")
	return &brightnessValue, nil
}

func getExifCFAPattern(e Exif) (any, error) {
	var cFAPattern any
	tag, err := e.Get(exif.CFAPattern)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &cFAPattern, nil
}

func getExifColorSpace(e Exif) (*int, error) {
	var colorSpace int
	tag, err := e.Get(exif.ColorSpace)
	if err != nil {
		return nil, err
	}
	colorSpace, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &colorSpace, nil
}

func getExifComponentsConfiguration(e Exif) (*string, error) {
	var componentsConfiguration string
	tag, err := e.Get(exif.ComponentsConfiguration)
	if err != nil {
		return nil, err
	}
	componentsConfiguration = strings.Trim(tag.String(), "\"")
	return &componentsConfiguration, nil
}

func getExifCompressedBitsPerPixel(e Exif) (*string, error) {
	var compressedBitsPerPixel string
	tag, err := e.Get(exif.CompressedBitsPerPixel)
	if err != nil {
		return nil, err
	}
	compressedBitsPerPixel = strings.Trim(tag.String(), "\"")
	return &compressedBitsPerPixel, nil
}

func getExifCompression(e Exif) (any, error) {
	var compression any
	tag, err := e.Get(exif.Compression)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &compression, nil
}

func getExifContrast(e Exif) (*int, error) {
	var contrast int
	tag, err := e.Get(exif.Contrast)
	if err != nil {
		return nil, err
	}
	contrast, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &contrast, nil
}

func getExifCopyright(e Exif) (any, error) {
	var copyright any
	tag, err := e.Get(exif.Copyright)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &copyright, nil
}

func getExifCustomRendered(e Exif) (*int, error) {
	var customRendered int
	tag, err := e.Get(exif.CustomRendered)
	if err != nil {
		return nil, err
	}
	customRendered, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &customRendered, nil
}

func getExifDateTime(e Exif) (*string, error) {
	var dateTime string
	tag, err := e.Get(exif.DateTime)
	if err != nil {
		return nil, err
	}
	dateTime, err = tag.StringVal()
	if err != nil {
		return nil, err
	}
	return &dateTime, nil
}

func getExifDateTimeDigitized(e Exif) (*string, error) {
	var dateTimeDigitized string
	tag, err := e.Get(exif.DateTimeDigitized)
	if err != nil {
		return nil, err
	}
	dateTimeDigitized, err = tag.StringVal()
	if err != nil {
		return nil, err
	}
	return &dateTimeDigitized, nil
}

func getExifDateTimeOriginal(e Exif) (*string, error) {
	var dateTimeOriginal string
	tag, err := e.Get(exif.DateTimeOriginal)
	if err != nil {
		return nil, err
	}
	dateTimeOriginal, err = tag.StringVal()
	if err != nil {
		return nil, err
	}
	return &dateTimeOriginal, nil
}

func getExifDeviceSettingDescription(e Exif) (any, error) {
	var deviceSettingDescription any
	tag, err := e.Get(exif.DeviceSettingDescription)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &deviceSettingDescription, nil
}

func getExifDigitalZoomRatio(e Exif) (*string, error) {
	var digitalZoomRatio string
	tag, err := e.Get(exif.DigitalZoomRatio)
	if err != nil {
		return nil, err
	}
	digitalZoomRatio = strings.Trim(tag.String(), "\"")
	return &digitalZoomRatio, nil
}

func getExifExifIFDPointer(e Exif) (*int, error) {
	var exifIFDPointer int
	tag, err := e.Get(exif.ExifIFDPointer)
	if err != nil {
		return nil, err
	}
	exifIFDPointer, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &exifIFDPointer, nil
}

func getExifExifVersion(e Exif) (*string, error) {
	var exifVersion string
	tag, err := e.Get(exif.ExifVersion)
	if err != nil {
		return nil, err
	}
	exifVersion = strings.Trim(tag.String(), "\"")
	return &exifVersion, nil
}

func getExifExposureBiasValue(e Exif) (*string, error) {
	var exposureBiasValue string
	tag, err := e.Get(exif.ExposureBiasValue)
	if err != nil {
		return nil, err
	}
	exposureBiasValue = strings.Trim(tag.String(), "\"")
	return &exposureBiasValue, nil
}

func getExifExposureIndex(e Exif) (any, error) {
	var exposureIndex any
	tag, err := e.Get(exif.ExposureIndex)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exposureIndex, nil
}

func getExifExposureMode(e Exif) (*int, error) {
	var exposureMode int
	tag, err := e.Get(exif.ExposureMode)
	if err != nil {
		return nil, err
	}
	exposureMode, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &exposureMode, nil
}

func getExifExposureProgram(e Exif) (*int, error) {
	var exposureProgram int
	tag, err := e.Get(exif.ExposureProgram)
	if err != nil {
		return nil, err
	}
	exposureProgram, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &exposureProgram, nil
}

func getExifExposureTime(e Exif) (*string, error) {
	var exposureTime string
	tag, err := e.Get(exif.ExposureTime)
	if err != nil {
		return nil, err
	}
	exposureTime = strings.Trim(tag.String(), "\"")
	return &exposureTime, nil
}

func getExifFNumber(e Exif) (*string, error) {
	var fNumber string
	tag, err := e.Get(exif.FNumber)
	if err != nil {
		return nil, err
	}
	fNumber = strings.Trim(tag.String(), "\"")
	return &fNumber, nil
}

func getExifFileSource(e Exif) (*string, error) {
	var fileSource string
	tag, err := e.Get(exif.FileSource)
	if err != nil {
		return nil, err
	}
	fileSource = strings.Trim(tag.String(), "\"")
	return &fileSource, nil
}

func getExifFlash(e Exif) (*int, error) {
	var flash int
	tag, err := e.Get(exif.Flash)
	if err != nil {
		return nil, err
	}
	flash, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &flash, nil
}

func getExifFlashEnergy(e Exif) (any, error) {
	var flashEnergy any
	tag, err := e.Get(exif.FlashEnergy)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &flashEnergy, nil
}

func getExifFlashpixVersion(e Exif) (*string, error) {
	var flashpixVersion string
	tag, err := e.Get(exif.FlashpixVersion)
	if err != nil {
		return nil, err
	}
	flashpixVersion = strings.Trim(tag.String(), "\"")
	return &flashpixVersion, nil
}

func getExifFocalLength(e Exif) (*string, error) {
	var focalLength string
	tag, err := e.Get(exif.FocalLength)
	if err != nil {
		return nil, err
	}
	focalLength = strings.Trim(tag.String(), "\"")
	return &focalLength, nil
}

func getExifFocalLengthIn35mmFilm(e Exif) (*int, error) {
	var focalLengthIn35mmFilm int
	tag, err := e.Get(exif.FocalLengthIn35mmFilm)
	if err != nil {
		return nil, err
	}
	focalLengthIn35mmFilm, err = tag.Int(0)
	if err != nil {
		return nil, err
	}
	return &focalLengthIn35mmFilm, nil
}

func getExifFocalPlaneResolutionUnit(e Exif) (any, error) {
	var focalPlaneResolutionUnit any
	tag, err := e.Get(exif.FocalPlaneResolutionUnit)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalPlaneResolutionUnit, nil
}

func getExifFocalPlaneXResolution(e Exif) (any, error) {
	var focalPlaneXResolution any
	tag, err := e.Get(exif.FocalPlaneXResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalPlaneXResolution, nil
}

func getExifFocalPlaneYResolution(e Exif) (any, error) {
	var focalPlaneYResolution any
	tag, err := e.Get(exif.FocalPlaneYResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalPlaneYResolution, nil
}

func getExifGPSAltitude(e Exif) (any, error) {
	var gpsAltitude any
	tag, err := e.Get(exif.GPSAltitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsAltitude, nil
}

func getExifGPSAltitudeRef(e Exif) (any, error) {
	var gpsAltitudeRef any
	tag, err := e.Get(exif.GPSAltitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsAltitudeRef, nil
}

func getExifGPSAreaInformation(e Exif) (any, error) {
	var gpsAreaInformation any
	tag, err := e.Get(exif.GPSAreaInformation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsAreaInformation, nil
}

func getExifGPSDOP(e Exif) (any, error) {
	var gpsDOP any
	tag, err := e.Get(exif.GPSDOP)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDOP, nil
}

func getExifGPSDateStamp(e Exif) (*string, error) {
	var gpsDateStamp string
	tag, err := e.Get(exif.GPSDateStamp)
	if err != nil {
		return nil, err
	}
	gpsDateStamp = strings.Trim(tag.String(), "\"")
	return &gpsDateStamp, nil
}

func getExifGPSDestBearing(e Exif) (any, error) {
	var gpsDestBearing any
	tag, err := e.Get(exif.GPSDestBearing)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestBearing, nil
}

func getExifGPSDestBearingRef(e Exif) (any, error) {
	var gpsDestBearingRef any
	tag, err := e.Get(exif.GPSDestBearingRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestBearingRef, nil
}

func getExifGPSDestDistance(e Exif) (any, error) {
	var gpsDestDistance any
	tag, err := e.Get(exif.GPSDestDistance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestDistance, nil
}

func getExifGPSDestDistanceRef(e Exif) (any, error) {
	var gpsDestDistanceRef any
	tag, err := e.Get(exif.GPSDestDistanceRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestDistanceRef, nil
}

func getExifGPSDestLatitude(e Exif) (any, error) {
	var gpsDestLatitude any
	tag, err := e.Get(exif.GPSDestLatitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestLatitude, nil
}

func getExifGPSDestLatitudeRef(e Exif) (any, error) {
	var gpsDestLatitudeRef any
	tag, err := e.Get(exif.GPSDestLatitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestLatitudeRef, nil
}

func getExifGPSDestLongitude(e Exif) (any, error) {
	var gpsDestLongitude any
	tag, err := e.Get(exif.GPSDestLongitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestLongitude, nil
}

func getExifGPSDestLongitudeRef(e Exif) (any, error) {
	var gpsDestLongitudeRef any
	tag, err := e.Get(exif.GPSDestLongitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsDestLongitudeRef, nil
}

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

func getExifGPSImgDirection(e Exif) (any, error) {
	var gpsImgDirection any
	tag, err := e.Get(exif.GPSImgDirection)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsImgDirection, nil
}

func getExifGPSImgDirectionRef(e Exif) (any, error) {
	var gpsImgDirectionRef any
	tag, err := e.Get(exif.GPSImgDirectionRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsImgDirectionRef, nil
}

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

func getExifGPSLatitude(e Exif) ([]string, error) {
	var gpsLatitude []string
	tag, err := e.Get(exif.GPSLatitude)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(tag.String()), &gpsLatitude)
	if err != nil {
		return nil, err
	}
	return gpsLatitude, nil
}

func getExifGPSLatitudeRef(e Exif) (*string, error) {
	var gpsLatitudeRef string
	tag, err := e.Get(exif.GPSLatitudeRef)
	if err != nil {
		return nil, err
	}
	gpsLatitudeRef = strings.Trim(tag.String(), "\"")
	return &gpsLatitudeRef, nil
}

func getExifGPSLongitude(e Exif) ([]string, error) {
	var gpsLongitude []string
	tag, err := e.Get(exif.GPSLongitude)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(tag.String()), &gpsLongitude)
	if err != nil {
		return nil, err
	}
	return gpsLongitude, nil
}

func getExifGPSLongitudeRef(e Exif) (*string, error) {
	var gpsLongitudeRef string
	tag, err := e.Get(exif.GPSLongitudeRef)
	if err != nil {
		return nil, err
	}
	gpsLongitudeRef = strings.Trim(tag.String(), "\"")
	return &gpsLongitudeRef, nil
}

func getExifGPSMapDatum(e Exif) (*string, error) {
	var gpsMapDatum string
	tag, err := e.Get(exif.GPSMapDatum)
	if err != nil {
		return nil, err
	}
	gpsMapDatum = strings.Trim(tag.String(), "\"")
	return &gpsMapDatum, nil
}

func getExifGPSMeasureMode(e Exif) (*int, error) {
	var gpsMeasureMode int
	tag, err := e.Get(exif.GPSMeasureMode)
	if err != nil {
		return nil, err
	}
	gpsMeasureMode, err = strconv.Atoi(strings.Trim(tag.String(), "\""))
	if err != nil {
		return nil, err
	}
	return &gpsMeasureMode, nil
}

func getExifGPSProcessingMethod(e Exif) (any, error) {
	var gpsProcessingMethod any
	tag, err := e.Get(exif.GPSProcessingMethod)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsProcessingMethod, nil
}

func getExifGPSSatelites(e Exif) (any, error) {
	var gpsSatelites any
	tag, err := e.Get(exif.GPSSatelites)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsSatelites, nil
}

func getExifGPSSpeed(e Exif) (any, error) {
	var gpsSpeed any
	tag, err := e.Get(exif.GPSSpeed)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsSpeed, nil
}

func getExifGPSSpeedRef(e Exif) (any, error) {
	var gpsSpeedRef any
	tag, err := e.Get(exif.GPSSpeedRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsSpeedRef, nil
}

func getExifGPSStatus(e Exif) (*string, error) {
	var gpsStatus string
	tag, err := e.Get(exif.GPSStatus)
	if err != nil {
		return nil, err
	}
	gpsStatus = strings.Trim(tag.String(), "\"")
	return &gpsStatus, nil
}

func getExifGPSTimeStamp(e Exif) ([]string, error) {
	var gpsTimeStamp []string
	tag, err := e.Get(exif.GPSTimeStamp)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(tag.String()), &gpsTimeStamp)
	if err != nil {
		return nil, err
	}
	return gpsTimeStamp, nil
}

func getExifGPSTrack(e Exif) (any, error) {
	var gpsTrack any
	tag, err := e.Get(exif.GPSTrack)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsTrack, nil
}

func getExifGPSTrackRef(e Exif) (any, error) {
	var gpsTrackRef any
	tag, err := e.Get(exif.GPSTrackRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gpsTrackRef, nil
}

func getExifGPSVersionID(e Exif) ([]string, error) {
	var gpsVersionID []string
	tag, err := e.Get(exif.GPSVersionID)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(tag.String()), &gpsVersionID)
	if err != nil {
		return nil, err
	}
	return gpsVersionID, nil
}

func getExifGainControl(e Exif) (any, error) {
	var gainControl any
	tag, err := e.Get(exif.GainControl)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gainControl, nil
}

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

func getExifImageDescription(e Exif) (any, error) {
	var imageDescription any
	tag, err := e.Get(exif.ImageDescription)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageDescription, nil
}

func getExifImageLength(e Exif) (any, error) {
	var imageLength any
	tag, err := e.Get(exif.ImageLength)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageLength, nil
}

func getExifImageUniqueID(e Exif) (any, error) {
	var imageUniqueID any
	tag, err := e.Get(exif.ImageUniqueID)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageUniqueID, nil
}

func getExifImageWidth(e Exif) (any, error) {
	var imageWidth any
	tag, err := e.Get(exif.ImageWidth)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageWidth, nil
}

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

func getExifInteroperabilityIndex(e Exif) (*string, error) {
	var interoperabilityIndex string
	tag, err := e.Get(exif.InteroperabilityIndex)
	if err != nil {
		return nil, err
	}
	interoperabilityIndex = strings.Trim(tag.String(), "\"")
	return &interoperabilityIndex, nil
}

func getExifLensMake(e Exif) (any, error) {
	var lensMake any
	tag, err := e.Get(exif.LensMake)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &lensMake, nil
}

func getExifLensModel(e Exif) (*string, error) {
	var lensModel string
	tag, err := e.Get(exif.LensModel)
	if err != nil {
		return nil, err
	}
	lensModel = strings.Trim(tag.String(), "\"")
	return &lensModel, nil
}

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

func getExifMake(e Exif) (*string, error) {
	var make string
	tag, err := e.Get(exif.Make)
	if err != nil {
		return nil, err
	}
	make = strings.Trim(tag.String(), "\"")
	return &make, nil
}

func getExifMakerNote(e Exif) (*string, error) {
	var makerNote string
	tag, err := e.Get(exif.MakerNote)
	if err != nil {
		return nil, err
	}
	makerNote = strings.Trim(tag.String(), "\"")
	return &makerNote, nil
}

func getExifMaxApertureValue(e Exif) (*string, error) {
	var maxApertureValue string
	tag, err := e.Get(exif.MaxApertureValue)
	if err != nil {
		return nil, err
	}
	maxApertureValue = strings.Trim(tag.String(), "\"")
	return &maxApertureValue, nil
}

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

func getExifModel(e Exif) (*string, error) {
	var model string
	tag, err := e.Get(exif.Model)
	if err != nil {
		return nil, err
	}
	model = strings.Trim(tag.String(), "\"")
	return &model, nil
}

func getExifOECF(e Exif) (any, error) {
	var oECF any
	tag, err := e.Get(exif.OECF)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &oECF, nil
}

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

func getExifPhotometricInterpretation(e Exif) (any, error) {
	var photometricInterpretation any
	tag, err := e.Get(exif.PhotometricInterpretation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &photometricInterpretation, nil
}

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

func getExifPlanarConfiguration(e Exif) (any, error) {
	var planarConfiguration any
	tag, err := e.Get(exif.PlanarConfiguration)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &planarConfiguration, nil
}

func getExifRelatedSoundFile(e Exif) (any, error) {
	var relatedSoundFile any
	tag, err := e.Get(exif.RelatedSoundFile)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &relatedSoundFile, nil
}

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

func getExifSceneType(e Exif) (*string, error) {
	var sceneType string
	tag, err := e.Get(exif.SceneType)
	if err != nil {
		return nil, err
	}
	sceneType = strings.Trim(tag.String(), "\"")
	return &sceneType, nil
}

func getExifSensingMethod(e Exif) (any, error) {
	var sensingMethod any
	tag, err := e.Get(exif.SensingMethod)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sensingMethod, nil
}

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

func getExifShutterSpeedValue(e Exif) (any, error) {
	var shutterSpeedValue any
	tag, err := e.Get(exif.ShutterSpeedValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &shutterSpeedValue, nil
}

func getExifSoftware(e Exif) (*string, error) {
	var software string
	tag, err := e.Get(exif.Software)
	if err != nil {
		return nil, err
	}
	software = strings.Trim(tag.String(), "\"")
	return &software, nil
}

func getExifSpatialFrequencyResponse(e Exif) (any, error) {
	var spatialFrequencyResponse any
	tag, err := e.Get(exif.SpatialFrequencyResponse)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &spatialFrequencyResponse, nil
}

func getExifSpectralSensitivity(e Exif) (any, error) {
	var spectralSensitivity any
	tag, err := e.Get(exif.SpectralSensitivity)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &spectralSensitivity, nil
}

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

func getExifSubjectArea(e Exif) (any, error) {
	var subjectArea any
	tag, err := e.Get(exif.SubjectArea)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectArea, nil
}

func getExifSubjectDistance(e Exif) (any, error) {
	var subjectDistance any
	tag, err := e.Get(exif.SubjectDistance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectDistance, nil
}

func getExifSubjectDistanceRange(e Exif) (any, error) {
	var subjectDistanceRange any
	tag, err := e.Get(exif.SubjectDistanceRange)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectDistanceRange, nil
}

func getExifSubjectLocation(e Exif) (any, error) {
	var subjectLocation any
	tag, err := e.Get(exif.SubjectLocation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectLocation, nil
}

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

func getExifUserComment(e Exif) (*string, error) {
	var userComment string
	tag, err := e.Get(exif.UserComment)
	if err != nil {
		return nil, err
	}
	userComment = strings.Trim(tag.String(), "\"")
	return &userComment, nil
}

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

func getExifXPAuthor(e Exif) (any, error) {
	var xPAuthor any
	tag, err := e.Get(exif.XPAuthor)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPAuthor, nil
}

func getExifXPComment(e Exif) (any, error) {
	var xPComment any
	tag, err := e.Get(exif.XPComment)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPComment, nil
}

func getExifXPKeywords(e Exif) (any, error) {
	var xPKeywords any
	tag, err := e.Get(exif.XPKeywords)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPKeywords, nil
}

func getExifXPSubject(e Exif) (any, error) {
	var xPSubject any
	tag, err := e.Get(exif.XPSubject)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPSubject, nil
}

func getExifXPTitle(e Exif) (any, error) {
	var xPTitle any
	tag, err := e.Get(exif.XPTitle)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPTitle, nil
}

func getExifXResolution(e Exif) (*string, error) {
	var xResolution string
	tag, err := e.Get(exif.XResolution)
	if err != nil {
		return nil, err
	}
	xResolution = strings.Trim(tag.String(), "\"")
	return &xResolution, nil
}

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

func getExifYCbCrSubSampling(e Exif) (any, error) {
	var yCbCrSubSampling any
	tag, err := e.Get(exif.YCbCrSubSampling)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &yCbCrSubSampling, nil
}

func getExifYResolution(e Exif) (*string, error) {
	var yResolution string
	tag, err := e.Get(exif.YResolution)
	if err != nil {
		return nil, err
	}
	yResolution = strings.Trim(tag.String(), "\"")
	return &yResolution, nil
}

func setExif(exifMetadata *ExifMetadata, e Exif) {
	setExifMetadataApertureValue(exifMetadata, e)
	setExifMetadataArtist(exifMetadata, e)
	setExifMetadataBitsPerSample(exifMetadata, e)
	setExifMetadataBrightnessValue(exifMetadata, e)
	setExifMetadataCFAPattern(exifMetadata, e)
	setExifMetadataColorSpace(exifMetadata, e)
	setExifMetadataComponentsConfiguration(exifMetadata, e)
	setExifMetadataCompressedBitsPerPixel(exifMetadata, e)
	setExifMetadataCompression(exifMetadata, e)
	setExifMetadataContrast(exifMetadata, e)
	setExifMetadataCopyright(exifMetadata, e)
	setExifMetadataCustomRendered(exifMetadata, e)
	setExifMetadataDateTime(exifMetadata, e)
	setExifMetadataDateTimeDigitized(exifMetadata, e)
	setExifMetadataDateTimeOriginal(exifMetadata, e)
	setExifMetadataDeviceSettingDescription(exifMetadata, e)
	setExifMetadataDigitalZoomRatio(exifMetadata, e)
	setExifMetadataExifIFDPointer(exifMetadata, e)
	setExifMetadataExifVersion(exifMetadata, e)
	setExifMetadataExposureBiasValue(exifMetadata, e)
	setExifMetadataExposureIndex(exifMetadata, e)
	setExifMetadataExposureMode(exifMetadata, e)
	setExifMetadataExposureProgram(exifMetadata, e)
	setExifMetadataExposureTime(exifMetadata, e)
	setExifMetadataFNumber(exifMetadata, e)
	setExifMetadataFileSource(exifMetadata, e)
	setExifMetadataFlash(exifMetadata, e)
	setExifMetadataFlashEnergy(exifMetadata, e)
	setExifMetadataFlashpixVersion(exifMetadata, e)
	setExifMetadataFocalLength(exifMetadata, e)
	setExifMetadataFocalLengthIn35mmFilm(exifMetadata, e)
	setExifMetadataFocalPlaneResolutionUnit(exifMetadata, e)
	setExifMetadataFocalPlaneXResolution(exifMetadata, e)
	setExifMetadataFocalPlaneYResolution(exifMetadata, e)
	setExifMetadataGPSAltitude(exifMetadata, e)
	setExifMetadataGPSAltitudeRef(exifMetadata, e)
	setExifMetadataGPSAreaInformation(exifMetadata, e)
	setExifMetadataGPSDOP(exifMetadata, e)
	setExifMetadataGPSDateStamp(exifMetadata, e)
	setExifMetadataGPSDestBearing(exifMetadata, e)
	setExifMetadataGPSDestBearingRef(exifMetadata, e)
	setExifMetadataGPSDestDistance(exifMetadata, e)
	setExifMetadataGPSDestDistanceRef(exifMetadata, e)
	setExifMetadataGPSDestLatitude(exifMetadata, e)
	setExifMetadataGPSDestLatitudeRef(exifMetadata, e)
	setExifMetadataGPSDestLongitude(exifMetadata, e)
	setExifMetadataGPSDestLongitudeRef(exifMetadata, e)
	setExifMetadataGPSDifferential(exifMetadata, e)
	setExifMetadataGPSImgDirection(exifMetadata, e)
	setExifMetadataGPSImgDirectionRef(exifMetadata, e)
	setExifMetadataGPSInfoIFDPointer(exifMetadata, e)
	setExifMetadataGPSLatitude(exifMetadata, e)
	setExifMetadataGPSLatitudeRef(exifMetadata, e)
	setExifMetadataGPSLongitude(exifMetadata, e)
	setExifMetadataGPSLongitudeRef(exifMetadata, e)
	setExifMetadataGPSMapDatum(exifMetadata, e)
	setExifMetadataGPSMeasureMode(exifMetadata, e)
	setExifMetadataGPSProcessingMethod(exifMetadata, e)
	setExifMetadataGPSSatelites(exifMetadata, e)
	setExifMetadataGPSSpeed(exifMetadata, e)
	setExifMetadataGPSSpeedRef(exifMetadata, e)
	setExifMetadataGPSStatus(exifMetadata, e)
	setExifMetadataGPSTimeStamp(exifMetadata, e)
	setExifMetadataGPSTrack(exifMetadata, e)
	setExifMetadataGPSTrackRef(exifMetadata, e)
	setExifMetadataGPSVersionID(exifMetadata, e)
	setExifMetadataGainControl(exifMetadata, e)
	setExifMetadataISOSpeedRatings(exifMetadata, e)
	setExifMetadataImageDescription(exifMetadata, e)
	setExifMetadataImageLength(exifMetadata, e)
	setExifMetadataImageUniqueID(exifMetadata, e)
	setExifMetadataImageWidth(exifMetadata, e)
	setExifMetadataInteroperabilityIFDPointer(exifMetadata, e)
	setExifMetadataInteroperabilityIndex(exifMetadata, e)
	setExifMetadataLensMake(exifMetadata, e)
	setExifMetadataLensModel(exifMetadata, e)
	setExifMetadataLightSource(exifMetadata, e)
	setExifMetadataMake(exifMetadata, e)
	setExifMetadataMakerNote(exifMetadata, e)
	setExifMetadataMaxApertureValue(exifMetadata, e)
	setExifMetadataMeteringMode(exifMetadata, e)
	setExifMetadataModel(exifMetadata, e)
	setExifMetadataOECF(exifMetadata, e)
	setExifMetadataOrientation(exifMetadata, e)
	setExifMetadataPhotometricInterpretation(exifMetadata, e)
	setExifMetadataPixelXDimension(exifMetadata, e)
	setExifMetadataPixelYDimension(exifMetadata, e)
	setExifMetadataPlanarConfiguration(exifMetadata, e)
	setExifMetadataRelatedSoundFile(exifMetadata, e)
	setExifMetadataResolutionUnit(exifMetadata, e)
	setExifMetadataSamplesPerPixel(exifMetadata, e)
	setExifMetadataSaturation(exifMetadata, e)
	setExifMetadataSceneCaptureType(exifMetadata, e)
	setExifMetadataSceneType(exifMetadata, e)
	setExifMetadataSensingMethod(exifMetadata, e)
	setExifMetadataSharpness(exifMetadata, e)
	setExifMetadataShutterSpeedValue(exifMetadata, e)
	setExifMetadataSoftware(exifMetadata, e)
	setExifMetadataSpatialFrequencyResponse(exifMetadata, e)
	setExifMetadataSpectralSensitivity(exifMetadata, e)
	setExifMetadataSubSecTime(exifMetadata, e)
	setExifMetadataSubSecTimeDigitized(exifMetadata, e)
	setExifMetadataSubSecTimeOriginal(exifMetadata, e)
	setExifMetadataSubjectArea(exifMetadata, e)
	setExifMetadataSubjectDistance(exifMetadata, e)
	setExifMetadataSubjectDistanceRange(exifMetadata, e)
	setExifMetadataSubjectLocation(exifMetadata, e)
	setExifMetadataThumbJPEGInterchangeFormat(exifMetadata, e)
	setExifMetadataThumbJPEGInterchangeFormatLength(exifMetadata, e)
	setExifMetadataUserComment(exifMetadata, e)
	setExifMetadataWhiteBalance(exifMetadata, e)
	setExifMetadataXPAuthor(exifMetadata, e)
	setExifMetadataXPComment(exifMetadata, e)
	setExifMetadataXPKeywords(exifMetadata, e)
	setExifMetadataXPSubject(exifMetadata, e)
	setExifMetadataXPTitle(exifMetadata, e)
	setExifMetadataXResolution(exifMetadata, e)
	setExifMetadataYCbCrPositioning(exifMetadata, e)
	setExifMetadataYCbCrSubSampling(exifMetadata, e)
	setExifMetadataYResolution(exifMetadata, e)
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
