package channelCombiner

import (
	"testing"
	"math/rand"
	"time"
)

func TestCustomUint8Conv(t *testing.T) {
	var tooLow, acceptable, tooHigh int = -5, 123, 5049
	if toUint8(tooLow) != uint8(0) {
		t.Fatal("Negative integers should be set to zero")
	} else if toUint8(acceptable) != uint8(acceptable) {
		t.Fatal("Integers within 0 and 255 should remain the same")
	} else if toUint8(tooHigh) != uint8(255) {
		t.Fatal("Integers greater than uint8 max should be clamped to 255")
	}
}

func TestCombineChannelsSize(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	xDim, yDim := 3, 2
	listOfChannels := make([][][]int, 3)
	for i := range listOfChannels {
		listOfChannels[i] = make([][]int, yDim)
		for y := range listOfChannels[i] {
			listOfChannels[i][y] = make([]int, xDim)
			for x := range listOfChannels[i][y] {
				listOfChannels[i][y][x] = rand.Intn(256)
			}
		}
	}
	combinedChannels := CombineRGB(listOfChannels)
	if len(combinedChannels) != yDim {
		t.Fatal("Y Dimension not preserved after combination")
	} else if len(combinedChannels[0]) != xDim {
		t.Fatal("X dimension not preserved after combination")
	}
}

func TestCombineChannelsValue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	xDim, yDim := 3, 2
	listOfChannels := make([][][]int, 3)
	for i := range listOfChannels {
		listOfChannels[i] = make([][]int, yDim)
		for y := range listOfChannels[i] {
			listOfChannels[i][y] = make([]int, xDim)
			for x := range listOfChannels[i][y] {
				listOfChannels[i][y][x] = rand.Intn(256)
			}
		}
	}
	randPosX, randPosY := rand.Intn(xDim), rand.Intn(yDim)
	origRedVal := uint8(listOfChannels[0][randPosY][randPosX])
	origGreenVal := uint8(listOfChannels[1][randPosY][randPosX])
	origBlueVal := uint8(listOfChannels[2][randPosY][randPosX])
	combinedChannels := CombineRGB(listOfChannels)
	if origRedVal != combinedChannels[randPosY][randPosX].R {
		t.Fatal("Red channel not properly preserved")
	} else if origGreenVal != combinedChannels[randPosY][randPosX].G {
		t.Fatal("Green channel not properly preserved")
	} else if origBlueVal != combinedChannels[randPosY][randPosX].B {
		t.Fatal("Blue channel not properly preserved")
	}
}
