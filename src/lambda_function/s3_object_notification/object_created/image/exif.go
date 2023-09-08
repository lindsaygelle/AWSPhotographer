package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

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

func getExif(filename string) {
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
	setExifMetadataApertureValue(&exifMetadata, e)
	setExifMetadataArtist(&exifMetadata, e)
	setExifMetadataBitsPerSample(&exifMetadata, e)
	setExifMetadataBrightnessValue(&exifMetadata, e)
	setExifMetadataCFAPattern(&exifMetadata, e)
	setExifMetadataColorSpace(&exifMetadata, e)
	setExifMetadataComponentsConfiguration(&exifMetadata, e)
	setExifMetadataCompressedBitsPerPixel(&exifMetadata, e)
	setExifMetadataCompression(&exifMetadata, e)
	setExifMetadataContrast(&exifMetadata, e)
	setExifMetadataCopyright(&exifMetadata, e)
	setExifMetadataCustomRendered(&exifMetadata, e)
	setExifMetadataDateTime(&exifMetadata, e)
	setExifMetadataDateTimeDigitized(&exifMetadata, e)
	setExifMetadataDateTimeOriginal(&exifMetadata, e)
	setExifMetadataDeviceSettingDescription(&exifMetadata, e)
	setExifMetadataDigitalZoomRatio(&exifMetadata, e)
	setExifMetadataExifIFDPointer(&exifMetadata, e)
	setExifMetadataExifVersion(&exifMetadata, e)
	setExifMetadataExposureBiasValue(&exifMetadata, e)
	setExifMetadataExposureIndex(&exifMetadata, e)
	setExifMetadataExposureMode(&exifMetadata, e)
	setExifMetadataExposureProgram(&exifMetadata, e)
	setExifMetadataExposureTime(&exifMetadata, e)
	setExifMetadataFNumber(&exifMetadata, e)
	setExifMetadataFileSource(&exifMetadata, e)
	setExifMetadataFlash(&exifMetadata, e)
	setExifMetadataFlashEnergy(&exifMetadata, e)
	setExifMetadataFlashpixVersion(&exifMetadata, e)
	setExifMetadataFocalLength(&exifMetadata, e)
	setExifMetadataFocalLengthIn35mmFilm(&exifMetadata, e)
	setExifMetadataFocalPlaneResolutionUnit(&exifMetadata, e)
	setExifMetadataFocalPlaneXResolution(&exifMetadata, e)
	setExifMetadataFocalPlaneYResolution(&exifMetadata, e)
	setExifMetadataGPSAltitude(&exifMetadata, e)
	setExifMetadataGPSAltitudeRef(&exifMetadata, e)
	setExifMetadataGPSAreaInformation(&exifMetadata, e)
	setExifMetadataGPSDOP(&exifMetadata, e)
	setExifMetadataGPSDateStamp(&exifMetadata, e)
	setExifMetadataGPSDestBearing(&exifMetadata, e)
	setExifMetadataGPSDestBearingRef(&exifMetadata, e)
	setExifMetadataGPSDestDistance(&exifMetadata, e)
	setExifMetadataGPSDestDistanceRef(&exifMetadata, e)
	setExifMetadataGPSDestLatitude(&exifMetadata, e)
	setExifMetadataGPSDestLatitudeRef(&exifMetadata, e)
	setExifMetadataGPSDestLongitude(&exifMetadata, e)
	setExifMetadataGPSDestLongitudeRef(&exifMetadata, e)
	setExifMetadataGPSDifferential(&exifMetadata, e)
	setExifMetadataGPSImgDirection(&exifMetadata, e)
	setExifMetadataGPSImgDirectionRef(&exifMetadata, e)
	setExifMetadataGPSInfoIFDPointer(&exifMetadata, e)
	setExifMetadataGPSLatitude(&exifMetadata, e)
	setExifMetadataGPSLatitudeRef(&exifMetadata, e)
	setExifMetadataGPSLongitude(&exifMetadata, e)
	setExifMetadataGPSLongitudeRef(&exifMetadata, e)
	setExifMetadataGPSMapDatum(&exifMetadata, e)
	setExifMetadataGPSMeasureMode(&exifMetadata, e)
	setExifMetadataGPSProcessingMethod(&exifMetadata, e)
	setExifMetadataGPSSatelites(&exifMetadata, e)
	setExifMetadataGPSSpeed(&exifMetadata, e)
	setExifMetadataGPSSpeedRef(&exifMetadata, e)
	setExifMetadataGPSStatus(&exifMetadata, e)
	setExifMetadataGPSTimeStamp(&exifMetadata, e)
	setExifMetadataGPSTrack(&exifMetadata, e)
	setExifMetadataGPSTrackRef(&exifMetadata, e)
	setExifMetadataGPSVersionID(&exifMetadata, e)
	setExifMetadataGainControl(&exifMetadata, e)
	setExifMetadataISOSpeedRatings(&exifMetadata, e)
	setExifMetadataImageDescription(&exifMetadata, e)
	setExifMetadataImageLength(&exifMetadata, e)
	setExifMetadataImageUniqueID(&exifMetadata, e)
	setExifMetadataImageWidth(&exifMetadata, e)
	setExifMetadataInteroperabilityIFDPointer(&exifMetadata, e)
	setExifMetadataInteroperabilityIndex(&exifMetadata, e)
	setExifMetadataLensMake(&exifMetadata, e)
	setExifMetadataLensModel(&exifMetadata, e)
	setExifMetadataLightSource(&exifMetadata, e)
	setExifMetadataMake(&exifMetadata, e)
	setExifMetadataMakerNote(&exifMetadata, e)
	setExifMetadataMaxApertureValue(&exifMetadata, e)
	setExifMetadataMeteringMode(&exifMetadata, e)
	setExifMetadataModel(&exifMetadata, e)
	setExifMetadataOECF(&exifMetadata, e)
	setExifMetadataOrientation(&exifMetadata, e)
	setExifMetadataPhotometricInterpretation(&exifMetadata, e)
	setExifMetadataPixelXDimension(&exifMetadata, e)
	setExifMetadataPixelYDimension(&exifMetadata, e)
	setExifMetadataPlanarConfiguration(&exifMetadata, e)
	setExifMetadataRelatedSoundFile(&exifMetadata, e)
	setExifMetadataResolutionUnit(&exifMetadata, e)
	setExifMetadataSamplesPerPixel(&exifMetadata, e)
	setExifMetadataSaturation(&exifMetadata, e)
	setExifMetadataSceneCaptureType(&exifMetadata, e)
	setExifMetadataSceneType(&exifMetadata, e)
	setExifMetadataSensingMethod(&exifMetadata, e)
	setExifMetadataSharpness(&exifMetadata, e)
	setExifMetadataShutterSpeedValue(&exifMetadata, e)
	setExifMetadataSoftware(&exifMetadata, e)
	setExifMetadataSpatialFrequencyResponse(&exifMetadata, e)
	setExifMetadataSpectralSensitivity(&exifMetadata, e)
	setExifMetadataSubSecTime(&exifMetadata, e)
	setExifMetadataSubSecTimeDigitized(&exifMetadata, e)
	setExifMetadataSubSecTimeOriginal(&exifMetadata, e)
	setExifMetadataSubjectArea(&exifMetadata, e)
	setExifMetadataSubjectDistance(&exifMetadata, e)
	setExifMetadataSubjectDistanceRange(&exifMetadata, e)
	setExifMetadataSubjectLocation(&exifMetadata, e)
	setExifMetadataThumbJPEGInterchangeFormat(&exifMetadata, e)
	setExifMetadataThumbJPEGInterchangeFormatLength(&exifMetadata, e)
	setExifMetadataUserComment(&exifMetadata, e)
	setExifMetadataWhiteBalance(&exifMetadata, e)
	setExifMetadataXPAuthor(&exifMetadata, e)
	setExifMetadataXPComment(&exifMetadata, e)
	setExifMetadataXPKeywords(&exifMetadata, e)
	setExifMetadataXPSubject(&exifMetadata, e)
	setExifMetadataXPTitle(&exifMetadata, e)
	setExifMetadataXResolution(&exifMetadata, e)
	setExifMetadataYCbCrPositioning(&exifMetadata, e)
	setExifMetadataYCbCrSubSampling(&exifMetadata, e)
	setExifMetadataYResolution(&exifMetadata, e)
	fmt.Sprintf("%v", exifMetadata)
}

func getExifApertureValue(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var apertureValue any
	tag, err := e.Get(exif.ApertureValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &apertureValue, nil
}

func getExifArtist(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var artist any
	tag, err := e.Get(exif.Artist)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &artist, nil
}

func getExifBitsPerSample(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var bitsPerSample any
	tag, err := e.Get(exif.BitsPerSample)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &bitsPerSample, nil
}

func getExifBrightnessValue(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (*string, error) {
	var brightnessValue string
	tag, err := e.Get(exif.BrightnessValue)
	if err != nil {
		return nil, err
	}
	rational, err := tag.Rat(0)
	if err != nil {
		panic(err)
	}
	brightnessValue = rational.String()
	return &brightnessValue, nil
}

func getExifCFAPattern(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var cFAPattern any
	tag, err := e.Get(exif.CFAPattern)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &cFAPattern, nil
}

func getExifColorSpace(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (*int, error) {
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

func getExifComponentsConfiguration(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (*string, error) {
	var componentsConfiguration string
	tag, err := e.Get(exif.ComponentsConfiguration)
	if err != nil {
		return nil, err
	}
	componentsConfiguration = tag.String()
	return &componentsConfiguration, nil
}

func getExifCompressedBitsPerPixel(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (*string, error) {
	var compressedBitsPerPixel string
	tag, err := e.Get(exif.CompressedBitsPerPixel)
	if err != nil {
		return nil, err
	}
	rational, err := tag.Rat(0)
	if err != nil {
		return nil, err
	}
	compressedBitsPerPixel = rational.String()
	return &compressedBitsPerPixel, nil
}

func getExifCompression(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var compression any
	tag, err := e.Get(exif.Compression)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &compression, nil
}

func getExifContrast(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
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

func getExifCopyright(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var copyright any
	tag, err := e.Get(exif.Copyright)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &copyright, nil
}

func getExifCustomRendered(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (*int, error) {
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

func getExifDateTime(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var dateTime any
	tag, err := e.Get(exif.DateTime)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &dateTime, nil
}

func getExifDateTimeDigitized(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var dateTimeDigitized any
	tag, err := e.Get(exif.DateTimeDigitized)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &dateTimeDigitized, nil
}

func getExifDateTimeOriginal(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var dateTimeOriginal any
	tag, err := e.Get(exif.DateTimeOriginal)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &dateTimeOriginal, nil
}

func getExifDeviceSettingDescription(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var deviceSettingDescription any
	tag, err := e.Get(exif.DeviceSettingDescription)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &deviceSettingDescription, nil
}

func getExifDigitalZoomRatio(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var digitalZoomRatio any
	tag, err := e.Get(exif.DigitalZoomRatio)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &digitalZoomRatio, nil
}

func getExifExifIFDPointer(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var exifIFDPointer any
	tag, err := e.Get(exif.ExifIFDPointer)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exifIFDPointer, nil
}

func getExifExifVersion(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var exifVersion any
	tag, err := e.Get(exif.ExifVersion)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exifVersion, nil
}

func getExifExposureBiasValue(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var exposureBiasValue any
	tag, err := e.Get(exif.ExposureBiasValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exposureBiasValue, nil
}

func getExifExposureIndex(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var exposureIndex any
	tag, err := e.Get(exif.ExposureIndex)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exposureIndex, nil
}

func getExifExposureMode(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var exposureMode any
	tag, err := e.Get(exif.ExposureMode)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exposureMode, nil
}

func getExifExposureProgram(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var exposureProgram any
	tag, err := e.Get(exif.ExposureProgram)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exposureProgram, nil
}

func getExifExposureTime(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var exposureTime any
	tag, err := e.Get(exif.ExposureTime)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &exposureTime, nil
}

func getExifFNumber(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var fNumber any
	tag, err := e.Get(exif.FNumber)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &fNumber, nil
}

func getExifFileSource(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var fileSource any
	tag, err := e.Get(exif.FileSource)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &fileSource, nil
}

func getExifFlash(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var flash any
	tag, err := e.Get(exif.Flash)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &flash, nil
}

func getExifFlashEnergy(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var flashEnergy any
	tag, err := e.Get(exif.FlashEnergy)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &flashEnergy, nil
}

func getExifFlashpixVersion(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var flashpixVersion any
	tag, err := e.Get(exif.FlashpixVersion)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &flashpixVersion, nil
}

func getExifFocalLength(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var focalLength any
	tag, err := e.Get(exif.FocalLength)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalLength, nil
}

func getExifFocalLengthIn35mmFilm(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var focalLengthIn35mmFilm any
	tag, err := e.Get(exif.FocalLengthIn35mmFilm)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalLengthIn35mmFilm, nil
}

func getExifFocalPlaneResolutionUnit(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var focalPlaneResolutionUnit any
	tag, err := e.Get(exif.FocalPlaneResolutionUnit)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalPlaneResolutionUnit, nil
}

func getExifFocalPlaneXResolution(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var focalPlaneXResolution any
	tag, err := e.Get(exif.FocalPlaneXResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalPlaneXResolution, nil
}

func getExifFocalPlaneYResolution(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var focalPlaneYResolution any
	tag, err := e.Get(exif.FocalPlaneYResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &focalPlaneYResolution, nil
}

func getExifGPSAltitude(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSAltitude any
	tag, err := e.Get(exif.GPSAltitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSAltitude, nil
}

func getExifGPSAltitudeRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSAltitudeRef any
	tag, err := e.Get(exif.GPSAltitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSAltitudeRef, nil
}

func getExifGPSAreaInformation(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSAreaInformation any
	tag, err := e.Get(exif.GPSAreaInformation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSAreaInformation, nil
}

func getExifGPSDOP(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDOP any
	tag, err := e.Get(exif.GPSDOP)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDOP, nil
}

func getExifGPSDateStamp(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDateStamp any
	tag, err := e.Get(exif.GPSDateStamp)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDateStamp, nil
}

func getExifGPSDestBearing(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestBearing any
	tag, err := e.Get(exif.GPSDestBearing)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestBearing, nil
}

func getExifGPSDestBearingRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestBearingRef any
	tag, err := e.Get(exif.GPSDestBearingRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestBearingRef, nil
}

func getExifGPSDestDistance(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestDistance any
	tag, err := e.Get(exif.GPSDestDistance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestDistance, nil
}

func getExifGPSDestDistanceRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestDistanceRef any
	tag, err := e.Get(exif.GPSDestDistanceRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestDistanceRef, nil
}

func getExifGPSDestLatitude(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestLatitude any
	tag, err := e.Get(exif.GPSDestLatitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestLatitude, nil
}

func getExifGPSDestLatitudeRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestLatitudeRef any
	tag, err := e.Get(exif.GPSDestLatitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestLatitudeRef, nil
}

func getExifGPSDestLongitude(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestLongitude any
	tag, err := e.Get(exif.GPSDestLongitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestLongitude, nil
}

func getExifGPSDestLongitudeRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDestLongitudeRef any
	tag, err := e.Get(exif.GPSDestLongitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDestLongitudeRef, nil
}

func getExifGPSDifferential(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSDifferential any
	tag, err := e.Get(exif.GPSDifferential)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSDifferential, nil
}

func getExifGPSImgDirection(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSImgDirection any
	tag, err := e.Get(exif.GPSImgDirection)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSImgDirection, nil
}

func getExifGPSImgDirectionRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSImgDirectionRef any
	tag, err := e.Get(exif.GPSImgDirectionRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSImgDirectionRef, nil
}

func getExifGPSInfoIFDPointer(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSInfoIFDPointer any
	tag, err := e.Get(exif.GPSInfoIFDPointer)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSInfoIFDPointer, nil
}

func getExifGPSLatitude(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSLatitude any
	tag, err := e.Get(exif.GPSLatitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSLatitude, nil
}

func getExifGPSLatitudeRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSLatitudeRef any
	tag, err := e.Get(exif.GPSLatitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSLatitudeRef, nil
}

func getExifGPSLongitude(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSLongitude any
	tag, err := e.Get(exif.GPSLongitude)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSLongitude, nil
}

func getExifGPSLongitudeRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSLongitudeRef any
	tag, err := e.Get(exif.GPSLongitudeRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSLongitudeRef, nil
}

func getExifGPSMapDatum(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSMapDatum any
	tag, err := e.Get(exif.GPSMapDatum)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSMapDatum, nil
}

func getExifGPSMeasureMode(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSMeasureMode any
	tag, err := e.Get(exif.GPSMeasureMode)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSMeasureMode, nil
}

func getExifGPSProcessingMethod(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSProcessingMethod any
	tag, err := e.Get(exif.GPSProcessingMethod)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSProcessingMethod, nil
}

func getExifGPSSatelites(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSSatelites any
	tag, err := e.Get(exif.GPSSatelites)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSSatelites, nil
}

func getExifGPSSpeed(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSSpeed any
	tag, err := e.Get(exif.GPSSpeed)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSSpeed, nil
}

func getExifGPSSpeedRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSSpeedRef any
	tag, err := e.Get(exif.GPSSpeedRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSSpeedRef, nil
}

func getExifGPSStatus(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSStatus any
	tag, err := e.Get(exif.GPSStatus)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSStatus, nil
}

func getExifGPSTimeStamp(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSTimeStamp any
	tag, err := e.Get(exif.GPSTimeStamp)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSTimeStamp, nil
}

func getExifGPSTrack(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSTrack any
	tag, err := e.Get(exif.GPSTrack)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSTrack, nil
}

func getExifGPSTrackRef(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSTrackRef any
	tag, err := e.Get(exif.GPSTrackRef)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSTrackRef, nil
}

func getExifGPSVersionID(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gPSVersionID any
	tag, err := e.Get(exif.GPSVersionID)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gPSVersionID, nil
}

func getExifGainControl(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var gainControl any
	tag, err := e.Get(exif.GainControl)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &gainControl, nil
}

func getExifISOSpeedRatings(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var iSOSpeedRatings any
	tag, err := e.Get(exif.ISOSpeedRatings)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &iSOSpeedRatings, nil
}

func getExifImageDescription(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var imageDescription any
	tag, err := e.Get(exif.ImageDescription)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageDescription, nil
}

func getExifImageLength(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var imageLength any
	tag, err := e.Get(exif.ImageLength)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageLength, nil
}

func getExifImageUniqueID(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var imageUniqueID any
	tag, err := e.Get(exif.ImageUniqueID)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageUniqueID, nil
}

func getExifImageWidth(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var imageWidth any
	tag, err := e.Get(exif.ImageWidth)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &imageWidth, nil
}

func getExifInteroperabilityIFDPointer(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var interoperabilityIFDPointer any
	tag, err := e.Get(exif.InteroperabilityIFDPointer)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &interoperabilityIFDPointer, nil
}

func getExifInteroperabilityIndex(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var interoperabilityIndex any
	tag, err := e.Get(exif.InteroperabilityIndex)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &interoperabilityIndex, nil
}

func getExifLensMake(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var lensMake any
	tag, err := e.Get(exif.LensMake)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &lensMake, nil
}

func getExifLensModel(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var lensModel any
	tag, err := e.Get(exif.LensModel)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &lensModel, nil
}

func getExifLightSource(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var lightSource any
	tag, err := e.Get(exif.LightSource)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &lightSource, nil
}

func getExifMake(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var make any
	tag, err := e.Get(exif.Make)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &make, nil
}

func getExifMakerNote(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var makerNote any
	tag, err := e.Get(exif.MakerNote)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &makerNote, nil
}

func getExifMaxApertureValue(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var maxApertureValue any
	tag, err := e.Get(exif.MaxApertureValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &maxApertureValue, nil
}

func getExifMeteringMode(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var meteringMode any
	tag, err := e.Get(exif.MeteringMode)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &meteringMode, nil
}

func getExifModel(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var model any
	tag, err := e.Get(exif.Model)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &model, nil
}

func getExifOECF(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var oECF any
	tag, err := e.Get(exif.OECF)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &oECF, nil
}

func getExifOrientation(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var orientation any
	tag, err := e.Get(exif.Orientation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &orientation, nil
}

func getExifPhotometricInterpretation(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var photometricInterpretation any
	tag, err := e.Get(exif.PhotometricInterpretation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &photometricInterpretation, nil
}

func getExifPixelXDimension(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var pixelXDimension any
	tag, err := e.Get(exif.PixelXDimension)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &pixelXDimension, nil
}

func getExifPixelYDimension(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var pixelYDimension any
	tag, err := e.Get(exif.PixelYDimension)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &pixelYDimension, nil
}

func getExifPlanarConfiguration(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var planarConfiguration any
	tag, err := e.Get(exif.PlanarConfiguration)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &planarConfiguration, nil
}

func getExifRelatedSoundFile(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var relatedSoundFile any
	tag, err := e.Get(exif.RelatedSoundFile)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &relatedSoundFile, nil
}

func getExifResolutionUnit(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var resolutionUnit any
	tag, err := e.Get(exif.ResolutionUnit)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &resolutionUnit, nil
}

func getExifSamplesPerPixel(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var samplesPerPixel any
	tag, err := e.Get(exif.SamplesPerPixel)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &samplesPerPixel, nil
}

func getExifSaturation(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var saturation any
	tag, err := e.Get(exif.Saturation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &saturation, nil
}

func getExifSceneCaptureType(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var sceneCaptureType any
	tag, err := e.Get(exif.SceneCaptureType)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sceneCaptureType, nil
}

func getExifSceneType(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var sceneType any
	tag, err := e.Get(exif.SceneType)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sceneType, nil
}

func getExifSensingMethod(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var sensingMethod any
	tag, err := e.Get(exif.SensingMethod)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sensingMethod, nil
}

func getExifSharpness(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var sharpness any
	tag, err := e.Get(exif.Sharpness)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &sharpness, nil
}

func getExifShutterSpeedValue(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var shutterSpeedValue any
	tag, err := e.Get(exif.ShutterSpeedValue)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &shutterSpeedValue, nil
}

func getExifSoftware(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var software any
	tag, err := e.Get(exif.Software)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &software, nil
}

func getExifSpatialFrequencyResponse(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var spatialFrequencyResponse any
	tag, err := e.Get(exif.SpatialFrequencyResponse)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &spatialFrequencyResponse, nil
}

func getExifSpectralSensitivity(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var spectralSensitivity any
	tag, err := e.Get(exif.SpectralSensitivity)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &spectralSensitivity, nil
}

func getExifSubSecTime(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var subSecTime any
	tag, err := e.Get(exif.SubSecTime)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subSecTime, nil
}

func getExifSubSecTimeDigitized(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var subSecTimeDigitized any
	tag, err := e.Get(exif.SubSecTimeDigitized)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subSecTimeDigitized, nil
}

func getExifSubSecTimeOriginal(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var subSecTimeOriginal any
	tag, err := e.Get(exif.SubSecTimeOriginal)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subSecTimeOriginal, nil
}

func getExifSubjectArea(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var subjectArea any
	tag, err := e.Get(exif.SubjectArea)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectArea, nil
}

func getExifSubjectDistance(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var subjectDistance any
	tag, err := e.Get(exif.SubjectDistance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectDistance, nil
}

func getExifSubjectDistanceRange(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var subjectDistanceRange any
	tag, err := e.Get(exif.SubjectDistanceRange)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectDistanceRange, nil
}

func getExifSubjectLocation(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var subjectLocation any
	tag, err := e.Get(exif.SubjectLocation)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &subjectLocation, nil
}

func getExifThumbJPEGInterchangeFormat(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var thumbJPEGInterchangeFormat any
	tag, err := e.Get(exif.ThumbJPEGInterchangeFormat)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &thumbJPEGInterchangeFormat, nil
}

func getExifThumbJPEGInterchangeFormatLength(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var thumbJPEGInterchangeFormatLength any
	tag, err := e.Get(exif.ThumbJPEGInterchangeFormatLength)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &thumbJPEGInterchangeFormatLength, nil
}

func getExifUserComment(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var userComment any
	tag, err := e.Get(exif.UserComment)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &userComment, nil
}

func getExifWhiteBalance(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var whiteBalance any
	tag, err := e.Get(exif.WhiteBalance)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &whiteBalance, nil
}

func getExifXPAuthor(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var xPAuthor any
	tag, err := e.Get(exif.XPAuthor)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPAuthor, nil
}

func getExifXPComment(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var xPComment any
	tag, err := e.Get(exif.XPComment)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPComment, nil
}

func getExifXPKeywords(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var xPKeywords any
	tag, err := e.Get(exif.XPKeywords)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPKeywords, nil
}

func getExifXPSubject(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var xPSubject any
	tag, err := e.Get(exif.XPSubject)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPSubject, nil
}

func getExifXPTitle(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var xPTitle any
	tag, err := e.Get(exif.XPTitle)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xPTitle, nil
}

func getExifXResolution(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var xResolution any
	tag, err := e.Get(exif.XResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &xResolution, nil
}

func getExifYCbCrPositioning(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var yCbCrPositioning any
	tag, err := e.Get(exif.YCbCrPositioning)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &yCbCrPositioning, nil
}

func getExifYCbCrSubSampling(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var yCbCrSubSampling any
	tag, err := e.Get(exif.YCbCrSubSampling)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &yCbCrSubSampling, nil
}

func getExifYResolution(e interface {
	Get(exif.FieldName) (*tiff.Tag, error)
}) (any, error) {
	var yResolution any
	tag, err := e.Get(exif.YResolution)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%v", tag)
	return &yResolution, nil
}

func logSetExifMetadataError(fieldName exif.FieldName, err error) {
	log.Printf("ExifMetadata.%s: Error=%s", fieldName, err)
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
