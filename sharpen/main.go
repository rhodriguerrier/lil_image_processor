package sharpen

import (
	"image"
	"image/color"
	"reflect"

	"github.com/rhodriguerrier/lil_image_processor/channelCombiner"
	"github.com/rhodriguerrier/lil_image_processor/channelSplitter"
	"github.com/rhodriguerrier/lil_image_processor/convolution"
)

type SharpenOperator struct {
	Kernel [][]int
}

func convolveChannel(channelMap [][]uint8, sharpKernel [][]int) [][]int {
	convolutionKernel := &convolution.Kernel{Operator: sharpKernel}
	return convolutionKernel.Convolve(channelMap)
}

func (s *SharpenOperator) CalculateSharpen(img image.Image) [][]color.RGBA {
	colourMaps := channelSplitter.SplitChannels(img)
	v := reflect.ValueOf(*colourMaps)
	var convolvedChannels [][][]int
	for i := 0; i < v.NumField(); i++ {
		convolvedChannels = append(
			convolvedChannels,
			convolveChannel(v.Field(i).Interface().([][]uint8), s.Kernel),
		)
	}
	return channelCombiner.CombineRGB(convolvedChannels)
}
