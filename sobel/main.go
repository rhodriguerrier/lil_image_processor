package sobel

import (
	"math"

	"github.com/rhodriguerrier/lil_image_processor/convolution"
)

type SobelOperator struct {
	XKernel [][]int
	YKernel [][]int
}

func GetKernels() *SobelOperator {
	xOperator := [][]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	yOperator := [][]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	return &SobelOperator{XKernel: xOperator, YKernel: yOperator}
}

func findOrientation(xVal, yVal int) float64 {
	if yVal < 0 {
		return ((2 * math.Pi) + math.Atan2(float64(yVal), float64(xVal))) * (180 / math.Pi)
	}
	return math.Atan2(float64(yVal), float64(xVal)) * (180 / math.Pi)
}

func (s *SobelOperator) CalculateSobel(grayValues [][]uint8) ([][]uint8, [][]float64) {
	dxKernel := &convolution.Kernel{Operator: s.XKernel}
	dyKernel := &convolution.Kernel{Operator: s.YKernel}

	dxVals := dxKernel.Convolve(grayValues)
	dyVals := dyKernel.Convolve(grayValues)

	outputImgVals := make([][]uint8, len(dxVals))
	outputImgTheta := make([][]float64, len(dxVals))
	for y := range outputImgVals {
		outputImgVals[y] = make([]uint8, len(dxVals[y]))
		outputImgTheta[y] = make([]float64, len(dxVals[y]))
		for x := range outputImgVals[y] {
			finalPixVal := math.Sqrt(math.Pow(float64(dxVals[y][x]), 2) + math.Pow(float64(dyVals[y][x]), 2))
			outputImgTheta[y][x] = findOrientation(dxVals[y][x], dyVals[y][x])
			if finalPixVal > 255 {
				finalPixVal = 255
			}
			outputImgVals[y][x] = uint8(finalPixVal)
		}
	}
	return outputImgVals, outputImgTheta
}
