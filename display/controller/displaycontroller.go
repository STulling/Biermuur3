package controller

import (
	"STulling/biermuur/display"
	"STulling/biermuur/display/effectlib"
	"STulling/biermuur/math"
)

var (
	callbacks = map[string]func(float64, float64){
		"wave":      effectlib.Wave,
		"debugwave": effectlib.DebugWave,
		"slowwave":  effectlib.SlowWave,
		"sparkle":   effectlib.Sparkle,
		"mond":      effectlib.Mond,
		"fill":      effectlib.Fill,
		"diamond":   effectlib.Ruit,
		"circle":    effectlib.Cirkel,
		"bars":      effectlib.Simple,
		"clear":     effectlib.Clear,
		"snake":     effectlib.Snake,
	}
	callback = callbacks["wave"]
)

func SetCallback(name string) {
	callback = callbacks[name]
}

func RunDisplayPipe(channel chan []int16) {
	go display.Init()
	for {
		if display.IsRunning() {
			break
		}
	}
	for {
		block := <-channel
		rms, tone := math.ProcessBlock(block)
		display.Primary = effectlib.Wheel(uint8(tone * 255))
		callback(rms, tone)
	}
}
