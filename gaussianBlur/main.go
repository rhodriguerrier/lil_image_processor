package gaussianBlur

import (
	"image"
	"image/color"
	"reflect"

	"github.com/rhodriguerrier/lil_image_processor/channelCombiner"
	"github.com/rhodriguerrier/lil_image_processor/channelSplitter"
	"github.com/rhodriguerrier/lil_image_processor/convolution"
)

type GaussianOperator struct {
	Kernel [][]int
}

func convolveChannel(channelMap [][]uint8, gaussianKernel [][]int) [][]int {
	convolutionKernel := &convolution.Kernel{Operator: gaussianKernel}
	return convolutionKernel.Convolve(channelMap)
}

func getGaussianDenominator(kernel [][]int) int {
	var weightSum int
	for y := range kernel {
		for x := range kernel[y] {
			weightSum += kernel[y][x]
		}
	}
	return weightSum
}

func applyGaussianDenominator(denominator int, pixelMap [][]int) [][]int {
	outputBlurMap := make([][]int, len(pixelMap))
	for y := range pixelMap {
		outputBlurMap[y] = make([]int, len(pixelMap[y]))
		for x := range pixelMap[y] {
			outputBlurMap[y][x] = pixelMap[y][x] / denominator
		}
	}
	return outputBlurMap
}

func (g *GaussianOperator) CalculateBlur(img image.Image) [][]color.RGBA {
	colourMaps := channelSplitter.SplitChannels(img)
	gaussianDenominator := getGaussianDenominator(g.Kernel)
	v := reflect.ValueOf(*colourMaps)
	var convolvedChannels [][][]int
	for i := 0; i < v.NumField(); i++ {
		convolvedChannels = append(
			convolvedChannels,
			applyGaussianDenominator(
				gaussianDenominator,
				convolveChannel(v.Field(i).Interface().([][]uint8), g.Kernel),
			),
		)
	}
	return channelCombiner.CombineRGB(convolvedChannels)

}
