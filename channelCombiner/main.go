package channelCombiner

import (
	"image/color"
)

func toUint8(input int) uint8 {
	if input < 0 {
		return uint8(0)
	} else if input > 255 {
		return uint8(255)
	}
	return uint8(input)
}

func CombineRGB(listOfChannels [][][]int) [][]color.RGBA {
	combinedMap := make([][]color.RGBA, len(listOfChannels[0]))
	for y := range combinedMap {
		combinedMap[y] = make([]color.RGBA, len(listOfChannels[0][0]))
		for x := range combinedMap[y] {
			combinedMap[y][x] = color.RGBA{
				R: toUint8(listOfChannels[0][y][x]),
				G: toUint8(listOfChannels[1][y][x]),
				B: toUint8(listOfChannels[2][y][x]),
				A: uint8(255),
			}
		}
	}
	return combinedMap
}
