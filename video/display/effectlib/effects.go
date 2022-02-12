package effectlib

import (
	"math/rand"
	"time"

	"STulling/video/display"
)

func RandomRGB() uint32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return display.RGBToColor(uint8(r1.Int()), uint8(r1.Int()), uint8(r1.Int()))
}

func Wheel(pos uint8) uint32 {
	if pos < 85 {
		return display.RGBToColor(pos*3, 255-pos*3, 0)
	}
	if pos < 170 {
		pos -= 85
		return display.RGBToColor(255-pos*3, 0, pos*3)
	}
	pos -= 170
	return display.RGBToColor(0, pos*3, 255-pos*3)
}
