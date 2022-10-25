package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kettek/apng"
	"golang.org/x/image/draw"
)

const (
	size  = 100
	noise = 60
)

func main() {

	fsf, err := os.Open("faceset_states.png")
	if err != nil {
		panic(err)
	}
	defer fsf.Close()
	fs, err := png.Decode(fsf)
	if err != nil {
		panic(err)
	}

	const counts = 3
	raws := make([]image.Image, counts)
	for i := 0; i < counts; i++ {
		name := fmt.Sprintf("aoisensi_omori_raw_ss_%v.png", i)
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		raws[i], err = png.Decode(f)
		if err != nil {
			panic(err)
		}
	}
	frames := make([]apng.Frame, counts)
	for i, raw := range raws {

		img1 := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.NearestNeighbor.Scale(img1, img1.Bounds(), fs, image.Rect(0, 0, 100, 100), draw.Over, nil)
		draw.CatmullRom.Scale(img1, img1.Bounds(), raw, raw.Bounds(), draw.Over, nil)

		/*
			buf := bytes.NewBuffer(make([]byte, 0, 1024*1024*1024*16))
			if err := jpeg.Encode(buf, img1, &jpeg.Options{
				Quality: noise,
			}); err != nil {
				panic(err)
			}
			img1, err := jpeg.Decode(buf)
			if err != nil {
				panic(err)
			}
		*/

		frame := apng.Frame{
			Image:            img1,
			DelayDenominator: 1000,
			DelayNumerator:   250,
		}

		frames[i] = frame
	}
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := apng.Encode(f, apng.APNG{
		Frames: frames,
	}); err != nil {
		panic(err)
	}
}
