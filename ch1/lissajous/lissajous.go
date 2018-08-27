package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	if err := lissajous(os.Stdout); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func lissajous(out io.Writer) error {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, fillPalette())
		colorIndex := uint8(rand.Intn(len(img.Palette)))
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		return err
	}

	return nil
}

func fillPalette() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 1},   //black
		color.RGBA{0, 0, 255, 1}, //blue
		color.RGBA{255, 0, 0, 1}, //red
		color.RGBA{0, 255, 0, 1}, //green
	}
}
