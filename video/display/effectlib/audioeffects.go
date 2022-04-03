package effectlib

import (
	"math"
	"math/rand"
	"time"

	"STulling/video/display"
)

var (
	t      = 0.
	xArray = make([]int, display.Width)
	snake  = make([][2]uint32, display.Width)
)

func Wave(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	dt := 0.1 * (1 + 3*pitch)
	t += dt
	for x := 0; x < display.Width; x++ {
		xVal := 3. * math.Pi * float64(x) / (display.Width - 1)
		xArray[x] = int(rms*display.Height/2*math.Sin(xVal+t) + display.Height/2)
	}
	for x := 0; x < display.Width; x++ {
		display.SetPixelColor(x, xArray[x], display.Primary)
		display.SetPixelColor(x, xArray[x]-1, display.Primary)
	}
	display.Render()
}

func SlowWave(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	dt := 0.1
	t += dt
	for x := 0; x < display.Width; x++ {
		xVal := 3. * math.Pi * float64(x) / (display.Width - 1)
		xArray[x] = int(rms*display.Height/2*math.Sin(xVal+t) + display.Height/2)
	}
	for x := 0; x < display.Width; x++ {
		display.SetPixelColor(x, xArray[x], display.Primary)
		display.SetPixelColor(x, xArray[x]-1, display.Primary)
	}
	display.Render()
}

func DebugWave(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	rms = 0.8
	dt := 0.3 //* (1 + 3*pitch)
	t += dt
	for x := 0; x < display.Width; x++ {
		x_val := 3. * math.Pi * float64(x) / (display.Width - 1)
		xArray[x] = int(rms*display.Height/2*math.Sin(x_val+t) + display.Height/2)
	}
	for x := 0; x < display.Width; x++ {
		display.SetPixelColor(x, xArray[x], display.RGBToColor(0, 255, 0))
		display.SetPixelColor(x, xArray[x]-1, display.RGBToColor(0, 255, 0))
	}
	display.Render()
}

func Snake(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	color := display.Primary
	height := uint32(pitch * display.Height)
	snake = append(snake[1:], [2]uint32{height, color})
	for i, data := range snake {
		display.SetPixelColor(i, int(data[0]), data[1])
		display.SetPixelColor(i, int(data[0])+1, data[1])
	}
	display.Render()
}

func Simple(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	amount := int(rms * display.LedCount)
	for i := 0; i < amount; i++ {
		display.SetLedColor(i, display.Primary)
	}
	display.Render()
}

func Sparkle(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	for i := 0; i < display.LedCount; i++ {
		if rand.Float64() < rms {
			display.SetLedColor(i, display.Primary)
		}
	}
	display.Render()
}

func Cirkel(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	xMid := display.Width/2 - 0.5
	yMid := display.Height/2 - 0.5
	radius := rms * 10
	for y := 0.; y < display.Height; y++ {
		for x := 0.; x < display.Width; x++ {
			afstand := math.Sqrt(math.Pow(y-yMid, 2) + math.Pow(x-xMid, 2))
			if afstand < radius {
				display.SetPixelColor(int(x), int(y), display.Primary)
			}
		}
	}
	display.Render()
}

func Ruit(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	xMid := display.Width/2 - 0.5
	yMid := display.Height/2 - 0.5
	i := rms * math.Max(display.Width, display.Height)
	for y := 0.; y < display.Height; y++ {
		for x := 0.; x < display.Width; x++ {
			if math.Abs(x-xMid) < i && math.Abs(y-yMid) < i && math.Abs(x-xMid)+math.Abs(y-yMid) < i {
				display.SetPixelColor(int(x), int(y), display.Primary)
			}
		}
	}
	display.Render()
}

func Fill(rms float64, pitch float64) {
	display.SetStrip(display.Primary)
	display.Render()
}

func Mond(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	c := float64(display.Width / 2)
	a := (rms * display.Height / 2) / math.Pow(((1-0.7*pitch)*display.Width)/2, 2)
	b := (1 + rms) * display.Height / 2
	for x := 0.; x < display.Width; x++ {
		y := int(-(a * math.Pow(x-c, 2)) + b)
		if y >= display.Height/2 {
			display.SetPixelColor(int(x), y, display.Primary)
			display.SetPixelColor(int(x), display.Height-y, display.Primary)
		}
	}
	display.Render()
}

var numbers_pixels = [10][5][3]bool{
	{
		{true, true, true},
		{true, false, true},
		{true, false, true},
		{true, false, true},
		{true, true, true},
	}, {
		{true, true, false},
		{false, true, false},
		{false, true, false},
		{false, true, false},
		{true, true, true},
	}, {
		{true, true, true},
		{false, false, true},
		{true, true, true},
		{true, false, false},
		{true, true, true},
	}, {
		{true, true, true},
		{false, false, true},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	}, {
		{true, false, true},
		{true, false, true},
		{true, true, true},
		{false, false, true},
		{false, false, true},
	}, {
		{true, true, true},
		{true, false, false},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	}, {
		{true, true, true},
		{true, false, false},
		{true, true, true},
		{true, false, true},
		{true, true, true},
	}, {
		{true, true, true},
		{false, false, true},
		{false, false, true},
		{false, false, true},
		{false, false, true},
	}, {
		{true, true, true},
		{true, false, true},
		{true, true, true},
		{true, false, true},
		{true, true, true},
	}, {
		{true, true, true},
		{true, false, true},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	},
}

var offset = [4][2]int{
	{1, 6},
	{5, 6},
	{12, 6},
	{16, 6},
}

func Clock(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	hm := time.Now().Format("1504")

	//draw numbers
	for c, offset := range offset {
		index := hm[c] - 48 // ascii to int
		number_pixels := numbers_pixels[index]
		for y, row := range number_pixels {
			for x, cell := range row {
				if cell {
					display.SetPixelColor(offset[0]+x, offset[1]+y, display.Primary)
				}
			}
		}
	}

	display.SetPixelColor(9, 6, display.Primary)
	display.SetPixelColor(10, 6, display.Primary)
	display.SetPixelColor(9, 7, display.Primary)
	display.SetPixelColor(10, 7, display.Primary)
	display.SetPixelColor(9, 9, display.Primary)
	display.SetPixelColor(10, 9, display.Primary)
	display.SetPixelColor(9, 10, display.Primary)
	display.SetPixelColor(10, 10, display.Primary)
	display.Render()
}

func Clear(rms float64, pitch float64) {
	display.Clear()
	display.Render()
}
