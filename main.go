package main

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/rhodriguerrier/lil_image_processor/gaussianBlur"
	"github.com/rhodriguerrier/lil_image_processor/grayscale"
	"github.com/rhodriguerrier/lil_image_processor/hsvToRgb"
	"github.com/rhodriguerrier/lil_image_processor/sharpen"
	"github.com/rhodriguerrier/lil_image_processor/sobel"
)

type imgProcessFunc func(image.Image) *image.RGBA

func grayScaleImg(img image.Image) *image.RGBA {
	grayImg := grayscale.ImageToGray(img)
	width, height := grayImg.Bounds().Max.X, grayImg.Bounds().Max.Y
	outputImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grayPixelVal := grayImg.GrayAt(x, y).Y
			outputImg.SetRGBA(x, y, color.RGBA{R: grayPixelVal, G: grayPixelVal, B: grayPixelVal, A: uint8(255)})
		}
	}
	return outputImg
}

func sobelGrayImg(img image.Image) *image.RGBA {
	grayscaledImg := grayscale.ImageToGray(img)
	grayPixelData := grayscale.GetGrayValues(grayscaledImg)
	sobelOperators := sobel.GetKernels()
	sobelPixelMap, _ := sobelOperators.CalculateSobel(grayPixelData.GrayValues)
	outputEdgeImg := image.NewRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(sobelPixelMap[0]), len(sobelPixelMap)},
		},
	)
	for y := range sobelPixelMap {
		for x := range sobelPixelMap[y] {
			outputEdgeImg.SetRGBA(x, y, color.RGBA{R: sobelPixelMap[y][x], G: sobelPixelMap[y][x], B: sobelPixelMap[y][x], A: uint8(255)})
		}
	}
	return outputEdgeImg
}

func sobelThetaImg(img image.Image) *image.RGBA {
	grayscaledImg := grayscale.ImageToGray(img)
	grayPixelData := grayscale.GetGrayValues(grayscaledImg)
	sobelOperators := sobel.GetKernels()
	sobelPixelMap, sobelThetaMap := sobelOperators.CalculateSobel(grayPixelData.GrayValues)
	outputEdgeImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{len(sobelPixelMap[0]), len(sobelPixelMap)}})
	for y := range sobelPixelMap {
		for x := range sobelPixelMap[y] {
			newGrayRatio := float64(sobelPixelMap[y][x]) / 255.0
			newHSV := hsvToRgb.HSV{Hue: sobelThetaMap[y][x], Saturation: 1.0, Value: newGrayRatio}
			rgbConvert := newHSV.HsvToRgb()
			outputEdgeImg.SetRGBA(
				x,
				y,
				color.RGBA{R: uint8(rgbConvert.Red), G: uint8(rgbConvert.Green), B: uint8(rgbConvert.Blue), A: uint8(255)},
			)
		}
	}
	return outputEdgeImg
}

func sharpenImg(img image.Image) *image.RGBA {
	sharpenOperator := &sharpen.SharpenOperator{
		Kernel: [][]int{
			{0, -1, 0},
			{-1, 5, -1},
			{0, -1, 0},
		},
	}
	sharpenPixelMap := sharpenOperator.CalculateSharpen(img)
	sharpenOutputImg := image.NewRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(sharpenPixelMap[0]), len(sharpenPixelMap)},
		},
	)
	for y := range sharpenPixelMap {
		for x := range sharpenPixelMap[y] {
			sharpenOutputImg.SetRGBA(x, y, sharpenPixelMap[y][x])
		}
	}
	return sharpenOutputImg
}

func blurImg(img image.Image) *image.RGBA {
	gaussianOperator := &gaussianBlur.GaussianOperator{
		Kernel: [][]int{
			{1, 2, 1},
			{2, 4, 2},
			{1, 2, 1},
		},
	}
	blurPixelMap := gaussianOperator.CalculateBlur(img)
	blurOutputImg := image.NewRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(blurPixelMap[0]), len(blurPixelMap)},
		},
	)
	for y := range blurPixelMap {
		for x := range blurPixelMap[y] {
			blurOutputImg.SetRGBA(x, y, blurPixelMap[y][x])
		}
	}
	return blurOutputImg
}

func main() {
	inputFileName := flag.String("file", "", "input file")
	editMode := flag.String("editMode", "g", "edit mode (e.g. g - grayscale, s - sobel, sc - coloured sobel orientation)")
	outputFileName := flag.String("outFile", "outputImg.jpg", "filename and extension string of output")
	flag.Parse()

	file, err := os.Open(*inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitOutPath := strings.Split(*outputFileName, ".")
	outputFileExt := splitOutPath[len(splitOutPath)-1]
	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	processesMap := map[string]imgProcessFunc {
		"g": grayScaleImg,
		"s": sobelGrayImg,
		"sc": sobelThetaImg,
		"sh": sharpenImg,
		"b": blurImg,
	}

	if outputFileExt == "jpg" {
		jpeg.Encode(outputFile, processesMap[*editMode](loadedImg), nil)
	} else if outputFileExt == "png" {
		png.Encode(outputFile, processesMap[*editMode](loadedImg))
	} else {
		log.Fatal("Unsupported output file extension provided")
	}

}
