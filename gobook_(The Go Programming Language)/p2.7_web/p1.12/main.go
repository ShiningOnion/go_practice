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
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gif", outputGif)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func outputGif(w http.ResponseWriter, r *http.Request) {

	cycles := 5.0 // number of complete x oscillator revolutions
	res := 0.001  // angular resolution
	size := 100   // image canvas covers [-size..+size]
	nframes := 64 // number of animation frames
	delay := 8    // delay between frames in 10ms units

	if keys, ok := r.URL.Query()["cycles"]; ok {
		if value, err := strconv.Atoi(keys[0]); err == nil {
			cycles = float64(value)
		} else {
			log.Fatal(err)
		}
	}

	if keys, ok := r.URL.Query()["res"]; ok {
		if s, err := strconv.ParseFloat(keys[0], 64); err == nil {
			res = s
		} else {
			log.Fatal(err)
		}
	}

	if keys, ok := r.URL.Query()["size"]; ok {
		if value, err := strconv.Atoi(keys[0]); err == nil {
			size = value
		} else {
			log.Fatal(err)
		}

	}

	if keys, ok := r.URL.Query()["nframes"]; ok {
		if value, err := strconv.Atoi(keys[0]); err == nil {
			nframes = value
		} else {
			log.Fatal(err)
		}
	}

	if keys, ok := r.URL.Query()["delay"]; ok {

		if value, err := strconv.Atoi(keys[0]); err == nil {
			delay = value
		} else {
			log.Fatal(err)
		}
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles float64, res float64, size int, nframes int, delay int) {
	const (
	// cycles  = 5     // number of complete x oscillator revolutions
	// res = 0.001 // angular resolution
	// size    = 100   // image canvas covers [-size..+size]
	// nframes = 64 // number of animation frames
	// delay   = 8  // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
