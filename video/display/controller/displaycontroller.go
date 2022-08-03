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
		"clock":     effectlib.Clock,
	}
	activations = map[string]func(float64) float64{
		"linear":              effectlib.Linear,
		"sine":                effectlib.Sine,
		"smoothstep":          effectlib.Smoothstep,
		"smootherstep":        effectlib.Smootherstep,
		"quadratic":           effectlib.Quadratic,
		"cubic":               effectlib.Cubic,
		"truncatedlinear":     effectlib.Truncatedlinear,
		"moretruncatedlinear": effectlib.MoreTruncatedlinear,
	}
	activation = activations["linear"]
	callback   = callbacks["clear"]
	maxRMS     = 1.0
	rms        float64
	tone       float64
)

func SetCallback(name string) {
	callback = callbacks[name]
}

func SetActivation(name string) {
	activation = activations[name]
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

		maxRMS = maxRMS * 0.999
		if rms > maxRMS {
			maxRMS = rms
		}
		rms = rms / maxRMS

		display.Primary = effectlib.Wheel(uint8(tone * 255))
		callback(activation(rms), tone)
	}
}
