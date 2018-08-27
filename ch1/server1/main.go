package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lHandler)

	err := http.ListenAndServe("localhost:8001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

func lHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if size := query.Get("size"); size != "" {
		sizeInt, err := strconv.Atoi(size)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := lissajous(w, sizeInt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	if err := lissajous(w, 100); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func lissajous(out io.Writer, size int) error {
	const (
		cycles  = 5
		res     = 0.001
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
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
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
