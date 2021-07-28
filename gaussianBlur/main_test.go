package gaussianBlur

import (
	"image"
	_ "image/jpeg"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestGetDenominator(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var sumAssert int
	testKernel := make([][]int, 5)
	for y := range testKernel {
		testKernel[y] = make([]int, 5)
		for x := range testKernel[y] {
			kernelVal := rand.Intn(10)
			sumAssert += kernelVal
			testKernel[y][x] = kernelVal
		}
	}
	if sumAssert != getGaussianDenominator(testKernel) {
		t.Fatal("Sum of kernel values incorrect")
	}
}

func TestApplyDenominator(t *testing.T) {
	testGaussianDenominator := 2
	testPixelMap := [][]int{
		{4, 4, 4},
		{4, 4, 4},
		{4, 4, 4},
	}
	expectedPixelMap := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}
	actualPixelMap := applyGaussianDenominator(
		testGaussianDenominator,
		testPixelMap,
	)

	if len(actualPixelMap) != len(expectedPixelMap) {
		t.Fatal("Pixel map y dimension not preserved")
	} else if len(actualPixelMap[0]) != len(expectedPixelMap[0]) {
		t.Fatal("Pixel map x dimension not preserved")
	}

	for y := range actualPixelMap {
		for x := range actualPixelMap[y] {
			if actualPixelMap[y][x] != expectedPixelMap[y][x] {
				t.Fatal("Expected value not met")
			}
		}
	}
}

func TestBlurShape(t *testing.T) {
	testKernel := &GaussianOperator{
		Kernel: [][]int{
			{1, 2, 1},
			{2, 4, 2},
			{1, 2, 1},
		},
	}
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Could not open image to begin testing")
	}
	defer file.Close()
	loadedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("sjhbdsfbjksfd")
	}
	xBound, yBound := loadedImg.Bounds().Max.X, loadedImg.Bounds().Max.Y
	outputPixelMap := testKernel.CalculateBlur(loadedImg)
	if len(outputPixelMap) != yBound {
		t.Fatal("Y dimension of image not preserved")
	} else if len(outputPixelMap[0]) != xBound {
		t.Fatal("X dimension of image not preserved")
	}
}
