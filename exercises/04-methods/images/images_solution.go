//go:build solution

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
	return color.RGBAModel
}

// Bounds returns the rectangle that holds the image, starting at (0, 0).
func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Width, m.Height)
}

// At returns the colour of the pixel at (x, y).
func (m Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}
