package rotate

import (
	"image"
	"image/color"
	"math"
)

func degToRads(degrees int) float64 {
	return float64(degrees) * (math.Pi / 180.0)
}

func rotatePixel(currentX, currentY, centreX, centreY int, rads float64) (float64, float64) {
	newX := (math.Cos(rads) * float64(currentX - centreX)) + (math.Sin(rads) * float64(currentY - centreY)) + float64(centreX)
	newY := (-math.Sin(rads) * float64(currentX - centreX)) + (math.Cos(rads) * float64(currentY - centreY)) + float64(centreY)
	return math.Round(newX), math.Round(newY)
}

func RotateImage(img image.Image, degrees int) *image.RGBA {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	centreX, centreY := width/2, height/2

	rotatedImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newX, newY := rotatePixel(x, y, centreX, centreY, degToRads(degrees))
			r, g, b, _ := img.At(x, y).RGBA()
			r, g, b = r>>8, g>>8, b>>8
			rotatedImg.SetRGBA(int(newX), int(newY), color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(255)})
		}
	}
	return rotatedImg
}
