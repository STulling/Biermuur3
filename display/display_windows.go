package display

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/faiface/gui/win"
	"github.com/faiface/mainthread"
)

const (
	CellSize = 20
	Width    = 20
	Height   = 18
	LedCount = Width * Height
)

var (
	w         *win.Win
	Primary   uint32 = RGBToColor(0, 255, 0)
	Secondary uint32 = RGBToColor(0, 0, 0)
	running   bool   = false
)

func IsRunning() bool {
	return running
}

func Render() {}

func SetPixelColor(x int, y int, colorz uint32) {
	if x < 0 || y < 0 {
		return
	}
	if x >= Width || y >= Height {
		return
	}
	w.Draw() <- func(drw draw.Image) image.Rectangle {
		offset := image.Point{x * CellSize, y * CellSize}
		r := image.Rectangle{offset, offset.Add(image.Point{CellSize, CellSize})}
		draw.Draw(drw, r, &image.Uniform{ColorToRGB(colorz)}, image.ZP, draw.Src)
		return r
	}
}

func SetLedColor(i int, color uint32) {
	SetPixelColor(i%Width, i/Width, color)
}

func SetStrip(color uint32) {
	w.Draw() <- func(drw draw.Image) image.Rectangle {
		r := image.Rectangle{image.Point{0, 0}, image.Point{CellSize * Width, CellSize * Height}}
		draw.Draw(drw, r, &image.Uniform{ColorToRGB(color)}, image.ZP, draw.Src)
		return r
	}
}

func Clear() {
	SetStrip(RGBToColor(0, 0, 0))
}

func RGBToColor(r uint8, g uint8, b uint8) uint32 {
	return uint32(uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | 0xff)
}

func ColorToRGB(rgba uint32) color.RGBA {
	a := uint8(rgba & 0xff)
	b := uint8((rgba >> 8) & 0xff)
	g := uint8((rgba >> 16) & 0xff)
	r := uint8((rgba >> 24) & 0xff)
	return color.RGBA{r, g, b, a}
}

func run() {
	window, err := win.New(win.Title("biermuur/gui"), win.Size(CellSize*Width, CellSize*Height))
	if err != nil {
		panic(err)
	}
	w = window

	running = true
	for event := range w.Events() {
		switch event.(type) {
		case win.WiClose:
			close(w.Draw())
		}
	}
}

func Init() {
	mainthread.Run(run)
}
