package convolution

import (
	"testing"
	"math/rand"
	"time"
)

func TestEmptyCopy(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randXLen, randYLen := rand.Intn(100), rand.Intn(100)
	newEmpty := emptyCopy(randYLen, randXLen)
	if len(newEmpty) != randYLen {
		t.Fatal("Original Y size not preserved in copy")
	} else if len(newEmpty[0]) != randXLen {
		t.Fatal("Original X size not preserved in copy")
	}
}

func TestNeighbourHood(t *testing.T) {
	testKernel := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testMapXDim, testMapYDim := rand.Intn(500), rand.Intn(500)
	testPixelMap := make([][]uint8, testMapYDim)
	for y := range testPixelMap {
		testPixelMap[y] = make([]uint8, testMapXDim)
		for x := range testPixelMap[y] {
			testPixelMap[y][x] = uint8(rand.Intn(256))
		}
	}
	randXPos, randYPos := rand.Intn(testMapXDim - 2) + 1, rand.Intn(testMapYDim - 2) + 1
	neighbourHoodValues := getNeighbourHood(len(testKernel), testPixelMap, randXPos, randYPos)
	if len(neighbourHoodValues) != len(testKernel) {
		t.Fatal("Resultant neighbourhood does not match kernel Y size")
	} else if len(neighbourHoodValues[0]) != len(testKernel[0]) {
		t.Fatal("Resultant neighbourhood does not macth kernel X size")
	} else if neighbourHoodValues[0][0] != int(testPixelMap[randYPos-1][randXPos-1]) {
		t.Fatal("Resultant neighbourhood not as expected")
	}
}

func TestSingleConvolution(t *testing.T) {
	testKernel := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	testNeighbourPixelMap := make([][]int, len(testKernel))
	var expectedSum int
	for y := range testNeighbourPixelMap {
		testNeighbourPixelMap[y] = make([]int, len(testKernel))
		for x := range testNeighbourPixelMap[y] {
			randomPixelVal := rand.Intn(256)
			testNeighbourPixelMap[y][x] = randomPixelVal
			expectedSum += randomPixelVal
		}
	}
	convolutionValue := getConvolvedPixel(testNeighbourPixelMap, testKernel)
	if convolutionValue != expectedSum {
		t.Fatal("Convolution calcualtion incorrect")
	}
}

func TestConvolutionShape(t *testing.T) {
	testKernel := &Kernel{
		Operator: [][]int{{0, -1, 0}, {-1, 5, -1}, {0, -1, 0}},
	}
	testMapXDim, testMapYDim := rand.Intn(500), rand.Intn(500)
	testPixelMap := make([][]uint8, testMapYDim)
	for y := range testPixelMap {
		testPixelMap[y] = make([]uint8, testMapXDim)
		for x := range testPixelMap[y] {
			testPixelMap[y][x] = uint8(rand.Intn(256))
		}
	}
	convolvedPixelMap := testKernel.Convolve(testPixelMap)
	if (len(convolvedPixelMap)) != len(testPixelMap) {
		t.Fatal("Y Dimension not preserved during convolution")
	} else if len(convolvedPixelMap[0]) != len(testPixelMap[0]) {
		t.Fatal("X Dimension not preserved during convoltuion")
	}

}

