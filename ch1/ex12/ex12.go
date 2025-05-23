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
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0, 239, 74, 255},    // green
	color.RGBA{115, 104, 251, 255}, // purple
	color.RGBA{249, 156, 40, 255},  // orange
}

const (
	whiteIndex  = 0 // first color in palette
	greenIndex  = 1 // next color in palette
	purpleIndex = 2
	orangeIndex = 3
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		cycles_string := query.Get("cycles")
		if len(cycles_string) > 0 {
			cycles, err := strconv.Atoi(cycles_string)

			if err != nil {
				fmt.Fprintf(w, "cycles needs to be a number\n")
			} else {
				lissajous(w, float64(cycles))
			}
		} else {
			lissajous(w, 4)
		}
	})

	// http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles float64) {
	if cycles == 0 {
		cycles = 4 // default
	}
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		var index uint8
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			cycle_part := math.Mod(t, 2*math.Pi)
			if cycle_part > math.Pi*2/3 && cycle_part < math.Pi*4/3 {
				index = purpleIndex
			} else if cycle_part > math.Pi*4/3 {
				index = orangeIndex
			} else {
				index = greenIndex
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
