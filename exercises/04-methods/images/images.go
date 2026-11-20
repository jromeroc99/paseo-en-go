//go:build !solution

// Package images draws a picture by implementing the image.Image interface.
package images

import (
	"image"
	"image/color"
)

// Image is a picture of Width by Height pixels.
type Image struct {
	Width, Height int
}

// ColorModel returns the colour model of the image: color.RGBAModel.
func (m Image) ColorModel() color.Model {
	// TODO: return color.RGBAModel.
	return nil
}

// Bounds returns the rectangle that holds the image, starting at (0, 0).
// Example: Image{Width: 10, Height: 20}.Bounds() returns image.Rect(0, 0, 10, 20).
func (m Image) Bounds() image.Rectangle {
	// TODO: use image.Rect with the width and the height.
	return image.Rectangle{}
}

// At returns the colour of the pixel at (x, y).
// Example: At(10, 20) returns color.RGBA{10, 20, 255, 255}.
func (m Image) At(x, y int) color.Color {
	// TODO: return color.RGBA{uint8(x), uint8(y), 255, 255}.
	return nil
}
