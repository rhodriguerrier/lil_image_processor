package sharpen

import (
	"testing"
	"os"
	"image"
	_ "image/jpeg"
)

func TestSharpenShape(t *testing.T) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Error opening test image file, cannot continue with test")
	}
	defer file.Close()

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Error decoding file to image format")
	}
	xBound, yBound := loadedImg.Bounds().Max.X, loadedImg.Bounds().Max.Y

	testSharpKernel := &SharpenOperator{
		Kernel: [][]int{
			{0, -1, 0},
			{-1, 5, -1},
			{0, -1, 0},
		},
	}
	returnColourMap := testSharpKernel.CalculateSharpen(loadedImg)
	if len(returnColourMap) != yBound {
		t.Fatal("Image Y dimension not preserved")
	} else if len(returnColourMap[0]) != xBound {
		t.Fatal("Image X dimension not preserved")
	}
}

