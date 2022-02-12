package display

import (
	"fmt"

	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	brightness = 255
	gpioPin    = 21
	freq       = 800000
	Width      = 20
	Height     = 18
	LedCount   = Width * Height
)

var (
	strip     ws
	Primary   uint32 = RGBToColor(0, 255, 0)
	Secondary uint32 = RGBToColor(0, 0, 0)
	running   bool   = false
)

type ws struct {
	ws2811 *ws281x.WS2811
}

func (ws *ws) init() error {
	err := ws.ws2811.Init()
	if err != nil {
		return err
	}

	return nil
}

func (ws *ws) close() {
	ws.ws2811.Fini()
}

func IsRunning() bool {
	return running
}

func Render() {
	strip.ws2811.Render()
}

func SetPixelColor(x int, y int, color uint32) {
	if x < 0 || y < 0 {
		return
	}
	if x >= Width || y >= Height {
		return
	}
	if y%2 == 1 {
		x = Width - 1 - x
	}
	strip.ws2811.Leds(0)[x+y*Width] = color
}

func SetLedColor(i int, color uint32) {
	strip.ws2811.Leds(0)[i] = color
}

func SetStrip(color uint32) {
	for i := 0; i < LedCount; i++ {
		strip.ws2811.Leds(0)[i] = color
	}
}

func Clear() {
	for i := 0; i < LedCount; i++ {
		strip.ws2811.Leds(0)[i] = 0
	}
}

func RGBToColor(r uint8, g uint8, b uint8) uint32 {
	return uint32(uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}

func Init() {
	opt := ws281x.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = LedCount
	opt.Channels[0].GpioPin = gpioPin
	opt.Frequency = freq

	ws2811, err := ws281x.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}

	strip = ws{
		ws2811: ws2811,
	}

	fmt.Println("Led strip hardware information: " + fmt.Sprint(ws281x.HwDetect()))

	err = strip.init()
	if err != nil {
		panic(err)
	}
	Render()
	running = true
}
