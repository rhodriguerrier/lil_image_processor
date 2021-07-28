package main

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/rhodriguerrier/lil_image_processor/gaussianBlur"
	"github.com/rhodriguerrier/lil_image_processor/grayscale"
	"github.com/rhodriguerrier/lil_image_processor/hsvToRgb"
	"github.com/rhodriguerrier/lil_image_processor/sharpen"
	"github.com/rhodriguerrier/lil_image_processor/sobel"
)

func grayScaleImg(img image.Image) *image.Gray {
	return grayscale.ImageToGray(img)
}

func sobelGrayImg(img image.Image) *image.Gray {
	grayscaledImg := grayscale.ImageToGray(img)
	grayPixelData := grayscale.GetGrayValues(grayscaledImg)
	sobelOperators := sobel.GetKernels()
	sobelPixelMap, _ := sobelOperators.CalculateSobel(grayPixelData.GrayValues)
	outputEdgeImg := image.NewGray(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(sobelPixelMap[0]), len(sobelPixelMap)},
		},
	)
	for y := range sobelPixelMap {
		for x := range sobelPixelMap[y] {
			outputEdgeImg.SetGray(x, y, color.Gray{Y: sobelPixelMap[y][x]})
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

	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	if *editMode == "g" {
		jpeg.Encode(outputFile, grayScaleImg(loadedImg), nil)
	} else if *editMode == "s" {
		jpeg.Encode(outputFile, sobelGrayImg(loadedImg), nil)
	} else if *editMode == "sc" {
		jpeg.Encode(outputFile, sobelThetaImg(loadedImg), nil)
	} else if *editMode == "sh" {
		jpeg.Encode(outputFile, sharpenImg(loadedImg), nil)
	} else if *editMode == "b" {
		jpeg.Encode(outputFile, blurImg(loadedImg), nil)
	} else {
		log.Fatal("Invalid edit move provided")
	}
}
