package hsvToRgb

import (
	"math"
)

type RGB struct {
	Red float64
	Green float64
	Blue float64
}

type HSV struct {
	Hue float64
	Saturation float64
	Value float64
}

func (h *HSV) HsvToRgb() *RGB {
	c := h.Value * h.Saturation
	m := h.Value - c
	x := c * (1 - math.Abs(math.Mod((h.Hue/60), 2.0) - 1))
	var rDash, gDash, bDash float64
	if (h.Hue >= 0.0) && (h.Hue < 60.0) {
		rDash, gDash, bDash = c, x, 0.0
	} else if (h.Hue >= 60.0) && (h.Hue < 120.0) {
		rDash, gDash, bDash = x, c, 0.0
	} else if (h.Hue >= 120.0) && (h.Hue < 180.0) {
		rDash, gDash, bDash = 0.0, c, x
	} else if (h.Hue >= 180.0) && (h.Hue < 240.0) {
		rDash, gDash, bDash = 0.0, x, c
	} else if (h.Hue >= 240.0) && (h.Hue < 300.0) {
		rDash, gDash, bDash = x, 0.0, c
	} else if (h.Hue >= 300.0) && (h.Hue < 360.0) {
		rDash, gDash, bDash = c, 0.0, x
	}
	return &RGB{
		Red: (rDash + m) * 255,
		Green: (gDash + m) * 255,
		Blue: (bDash + m) * 255,
	}
}
