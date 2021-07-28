package grayscale

import (
	"testing"
	"image"
	_ "image/jpeg"
	"os"
)

func TestGrayScaleSize(t *testing.T) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Failed to open file, cannot conduct test")
	}
	defer file.Close()

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Failed to load image, cannot conduct test")
	}
	xBound, yBound := loadedImg.Bounds().Max.X, loadedImg.Bounds().Max.Y

	grayImg := ImageToGray(loadedImg)
	if grayImg.Bounds().Max.X != xBound {
		t.Fatal("Image X dimension not preserved")
	} else if grayImg.Bounds().Max.Y != yBound {
		t.Fatal("Image Y dimension not preserved")
	}
}

func TestGrayScaleValSize(t *testing.T) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Failed to open file, cannot conduct test")
	}
	defer file.Close()

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Failed to load image, cannot conduct test")
	}

	grayImg := ImageToGray(loadedImg)
	xBound, yBound := grayImg.Bounds().Max.X, grayImg.Bounds().Max.Y
	grayValues := GetGrayValues(grayImg)
	if len(grayValues.GrayValues) != yBound {
		t.Fatal("Image size not preserved when converting to value array")
	} else if len(grayValues.GrayValues[0]) != xBound {
		t.Fatal("Image size not preserved when converting to value array")
	}


}
