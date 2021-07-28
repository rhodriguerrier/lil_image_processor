package hsvToRgb

import (
	"testing"
)

func TestHsvConversion(t *testing.T) {
	testRedHSV := &HSV{Hue: 0.0, Saturation: 1.0, Value: 1.0}	// Known HSV value for RED (255, 0, 0)
	testRedRGB := &RGB{Red: 255.0, Green: 0.0, Blue: 0.0}

	testGreenHSV := &HSV{Hue: 120.0, Saturation: 1.0, Value: 1.0}	// Known HSV value for GREEN (0, 255, 0)
	testGreenRGB := &RGB{Red: 0.0, Green: 255.0, Blue: 0.0}

	testBlueHSV := &HSV{Hue: 240.0, Saturation: 1.0, Value: 1.0}	// Known HSV value for BLUE (0, 0, 255)
	testBlueRGB := &RGB{Red: 0.0, Green: 0.0, Blue: 255.0}

	if testRedHSV.HsvToRgb().Red != testRedRGB.Red || testRedHSV.HsvToRgb().Green != testRedRGB.Green || testRedHSV.HsvToRgb().Blue != testRedRGB.Blue {
		t.Fatal("HSV calculation incorrect")
	} else if testGreenHSV.HsvToRgb().Red != testGreenRGB.Red || testGreenHSV.HsvToRgb().Green != testGreenRGB.Green || testGreenHSV.HsvToRgb().Blue != testGreenRGB.Blue {
		t.Fatal("HSV calculation incorrect")
	} else if testBlueHSV.HsvToRgb().Red != testBlueRGB.Red || testBlueHSV.HsvToRgb().Green != testBlueRGB.Green || testBlueHSV.HsvToRgb().Blue != testBlueRGB.Blue {
		t.Fatal("HSV calculation incorrect")
	}
}
