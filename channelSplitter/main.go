package channelSplitter

import (
	"image"
)

type ColourChannels struct {
	Red [][]uint8
	Green [][]uint8
	Blue [][]uint8
}

func SplitChannels(img image.Image) *ColourChannels {
	bounds := img.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	outputRed := make([][]uint8, maxY)
	outputGreen := make([][]uint8, maxY)
	outputBlue := make([][]uint8, maxY)
	for y := 0; y < maxY; y++ {
		outputRed[y] = make([]uint8, maxX)
		outputGreen[y] = make([]uint8, maxX)
		outputBlue[y] = make([]uint8, maxX)
		for x := 0; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r, g, b = r>>8, g>>8, b>>8
			outputRed[y][x] = uint8(r)
			outputGreen[y][x] = uint8(g)
			outputBlue[y][x] = uint8(b)
		}
	}
	return &ColourChannels{Red: outputRed, Green: outputGreen, Blue: outputBlue}
}
