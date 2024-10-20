package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 
	width, height := 400, 400
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, white)
		}
	}

	// 
	skinColor := color.RGBA{255, 220, 180, 255} // 
	for x := 100; x < 300; x++ {
		for y := 100; y < 300; y++ {
			if (x-200)*(x-200)+(y-200)*(y-200) <= 100*100 {
				img.Set(x, y, skinColor)
			}
		}
	}

	// 
	eyeColor := color.RGBA{0, 0, 0, 255} // 
	for x := 150; x < 170; x++ {
		for y := 150; y < 170; y++ {
			if (x-160)*(x-160)+(y-160)*(y-160) <= 10*10 {
				img.Set(x, y, eyeColor)
			}
		}
	}
	for x := 230; x < 250; x++ {
		for y := 150; y < 170; y++ {
			if (x-240)*(x-240)+(y-160)*(y-160) <= 10*10 {
				img.Set(x, y, eyeColor)
			}
		}
	}

	// 
	for x := 170; x < 230; x++ {
		img.Set(x, 250, eyeColor)
	}

	// 
	f, _ := os.Create("miku_simple.png")
	defer f.Close()
	png.Encode(f, img)
}