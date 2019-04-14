// Lissajous generates GIF animations of tandom Lissajous figures.
package main

import (
	"os"
	"crypto/rand"
	"math"
	"image/gif"
	"io"
	"image/color"
	"image"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100  // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames 
		delay = 8 // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect :
	}
}