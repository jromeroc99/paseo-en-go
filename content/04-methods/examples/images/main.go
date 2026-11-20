package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	m.Set(10, 20, color.RGBA{0, 173, 216, 255})

	fmt.Println(m.Bounds())
	fmt.Println(m.At(10, 20))
	fmt.Println(m.At(0, 0))
}
