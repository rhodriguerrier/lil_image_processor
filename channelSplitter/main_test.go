package channelSplitter

import (
	"testing"
	"os"
	"image"
	_ "image/jpeg"
	"math/rand"
	"time"
)

func TestChannelSplitterShape(t *testing.T) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Failed to open file, cannot continue on to testing")
	}
	defer file.Close()

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Failed to load image, cannot continue on to testing")
	}
	xBound, yBound := loadedImg.Bounds().Max.X, loadedImg.Bounds().Max.Y
	splitChannelStruct := SplitChannels(loadedImg)
	if len(splitChannelStruct.Red) != yBound {
		t.Fatal("Image y dimension not preserved in red channel")
	} else if len(splitChannelStruct.Green) != yBound {
		t.Fatal("Image y dimension not preserved in green channel")
	} else if len(splitChannelStruct.Blue) != yBound {
		t.Fatal("Image y dimension not preserved in blue channel")
	} else if len(splitChannelStruct.Red[0]) != xBound {
		t.Fatal("Image x dimension not preserved in red channel")
	} else if len(splitChannelStruct.Green[0]) != xBound {
		t.Fatal("Image x dimension not preserved in green channel")
	} else if len(splitChannelStruct.Blue[0]) != xBound {
		t.Fatal("Image x dimension not preserved in blue channel")
	}
}

func TestChannelSplitterValue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Failed to open file, cannot continue on to testing")
	}
	defer file.Close()

	loadedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Failed to load image, cannot continue on to testing")
	}
	xBound, yBound := loadedImg.Bounds().Max.X, loadedImg.Bounds().Max.Y
	randPosX, randPosY := rand.Intn(xBound), rand.Intn(yBound)
	r, g, b, _ := loadedImg.At(randPosX, randPosY).RGBA()
	r, g, b = r>>8, g>>8, b>>8
	splitChannelStruct := SplitChannels(loadedImg)
	if splitChannelStruct.Red[randPosY][randPosX] != uint8(r) {
		t.Fatal("Red channel value not preserved when splitting")
	} else if splitChannelStruct.Green[randPosY][randPosX] != uint8(g) {
		t.Fatal("Green channel value not preserved when splitting")
	} else if splitChannelStruct.Blue[randPosY][randPosX] != uint8(b) {
		t.Fatal("Blue channel value not preserved when splitting")
	}



}
