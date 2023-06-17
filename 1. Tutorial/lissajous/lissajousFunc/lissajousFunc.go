package lissajousFunc

import 
(
"io"
"image/color"
"math/rand"
"image"
"math"
"image/gif"
"fmt"
)

type LissajousOpts struct  {
	Cycles float64
	Res float64
	Size int
	Nframes int
	Delay int
}

var opts LissajousOpts
var palette = []color.Color{color.White}

const (
	whiteIndex = 0
	blackIndex = 1
)

func fillPalette(){
	for i := uint8(0); i < 15; i++ {
		for j := uint8(0); j < 15; j++ {
			palette = append(palette, color.RGBA{i * 16, 0x00, j * 16, 0xff})
		}
	}
}

func setOpts(funcOpts LissajousOpts){
	if funcOpts.Cycles + 1 <= 1e-6 {
		opts.Cycles = 5.0
	} else {
		opts.Cycles = funcOpts.Cycles
	}
	
	if funcOpts.Res + 1 <= 1e-6 {
		opts.Res = 0.001
	} else {
		opts.Res = funcOpts.Res
	}

	if funcOpts.Size  == -1 {
		opts.Size = 250
	} else {
		opts.Size = funcOpts.Size
	}
	if funcOpts.Nframes == -1 {
		opts.Nframes = 64
	} else {
		opts.Nframes = funcOpts.Nframes
	}
	if funcOpts.Delay == -1 {
		opts.Delay = 8
	} else {
		opts.Delay = funcOpts.Delay
	}
}

func Lissajous(out io.Writer, passedOpts LissajousOpts){
	setOpts(passedOpts)
	if len(palette) <= 1 {
		fillPalette()
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: opts.Nframes}
	phase := 0.0
	for i := 0; i < opts.Nframes; i++{
		rect := image.Rect(0, 0, 2*opts.Size+1, 2*opts.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < opts.Cycles*math.Pi; t += opts.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(opts.Size+int(x*float64(opts.Size)+0.5), opts.Size+int(y*float64(opts.Size)+0.5), uint8(int(t * 10) % 255))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, opts.Delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Println(err)
	}
}