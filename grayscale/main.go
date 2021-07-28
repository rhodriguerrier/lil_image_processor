package grayscale

import (
	"image"
	_ "image/jpeg"
)

type GrayImgData struct {
	GrayValues [][]uint8
}

func ImageToGray(colourImg image.Image) *image.Gray {
	bounds := colourImg.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	grayImg := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grayImg.Set(x, y, colourImg.At(x, y))
		}
	}

	return grayImg
}

func GetGrayValues(grayImg *image.Gray) *GrayImgData {
	bounds := grayImg.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	grayImgArr := make([][]uint8, height)
	for i := range grayImgArr {
		grayImgArr[i] = make([]uint8, width)
		for j := range grayImgArr[i] {
			grayImgArr[i][j] = grayImg.GrayAt(j, i).Y
		}
	}
	return &GrayImgData{GrayValues: grayImgArr}

}

