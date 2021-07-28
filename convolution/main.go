package convolution

import ()

type Kernel struct {
	Operator [][]int
}

func emptyCopy(yLen, xLen int) [][]int {
	outputShell := make([][]int, yLen)
	for i := range outputShell {
		outputShell[i] = make([]int, xLen)
	}
	return outputShell
}

func getNeighbourHood(sqKernelLen int, matrixPoints [][]uint8, pixelPosX, pixelPosY int) [][]int {
	bufferSize := sqKernelLen/2
	neighbourHoodPoints := emptyCopy(sqKernelLen, sqKernelLen)
	yCount, xCount := 0, 0
	for y := pixelPosY - bufferSize; y <= pixelPosY + bufferSize; y++ {
		for x := pixelPosX - bufferSize; x <= pixelPosX + bufferSize; x++ {
			neighbourHoodPoints[yCount][xCount] = int(matrixPoints[y][x])
			xCount++
		}
		xCount = 0
		yCount++
	}
	return neighbourHoodPoints
}

func getConvolvedPixel(pointNeighbours [][]int, operatorKernel [][]int) int {
	var convolutionSum int
	for y := range operatorKernel {
		for x := range operatorKernel[y] {
			convolutionSum += pointNeighbours[y][x] * operatorKernel[y][x]
		}
	}
	return convolutionSum
}

func (k *Kernel) Convolve(inputMatrix [][]uint8) [][]int {
	outputMatrix := emptyCopy(len(inputMatrix), len(inputMatrix[0]))
	for y := range outputMatrix {
		for x := range outputMatrix[y] {
			if y == 0 || y == len(outputMatrix) - 1 {
				outputMatrix[y][x] = 0
			} else if x == 0 || x == len(outputMatrix[y]) - 1 {
				outputMatrix[y][x] = 0
			} else {
				toConvolve := getNeighbourHood(len(k.Operator), inputMatrix, x, y)
				outputMatrix[y][x] = getConvolvedPixel(toConvolve, k.Operator)
			}
		}
	}
	return outputMatrix
}
