package controller

import (
	"STulling/video/display"
	"STulling/video/display/effectlib"
	"bufio"
	"fmt"
	"os"
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
	rms      float64
	tone     float64
)

func SetCallback(name string) {
	callback = callbacks[name]
}

func RunDisplayPipe() {
	go display.Init()
	reader := bufio.NewReader(os.Stdin)
	for {
		if display.IsRunning() {
			break
		}
	}
	for {
		line, _ := reader.ReadString(';')
		fmt.Sscanf(line, "%f, %f;", &rms, &tone)
		display.Primary = effectlib.Wheel(uint8(tone * 255))
		callback(rms, tone)
	}
}
