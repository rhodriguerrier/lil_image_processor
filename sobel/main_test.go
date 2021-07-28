package sobel

import (
	"image"
	_ "image/jpeg"
	"os"
	"testing"

	"github.com/rhodriguerrier/lil_image_processor/grayscale"
)

func TestSobelSizes(t *testing.T) {
	testSobelOperator := GetKernels()
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Failed to open test image, cannot conduct test")
	}

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Failed to load image, cannot conduct test")
	}
	grayImg := grayscale.ImageToGray(loadedImg)
	grayValues := grayscale.GetGrayValues(grayImg)
	sobelPixelMap, sobelThetaMap := testSobelOperator.CalculateSobel(grayValues.GrayValues)
	if len(sobelPixelMap) != loadedImg.Bounds().Max.Y {
		t.Fatal("Image Y size not preserved in pixel values map")
	} else if len(sobelPixelMap[0]) != loadedImg.Bounds().Max.X {
		t.Fatal("Image X size not preserved in pixel values map")
	} else if len(sobelThetaMap) != loadedImg.Bounds().Max.Y {
		t.Fatal("Image Y size not preserved in pixel theta map")
	} else if len(sobelThetaMap[0]) != loadedImg.Bounds().Max.X {
		t.Fatal("Image X sixe not preserved in pixel theta map")
	}
}

func TestOrientationCalculation(t *testing.T) {
	if findOrientation(1, 1) != 45.0 {
		t.Fatal("Orientation calculation incorrect")
	} else if findOrientation(-1, 1) != 135.0 {
		t.Fatal("Orientation calculation incorrect")
	} else if findOrientation(-1, -1) != 225.0 {
		t.Fatal("Orientation calcultion incorrect")
	} else if findOrientation(1, -1) != 315.0 {
		t.Fatal("Orientation calculation incorrect")
	}
}
