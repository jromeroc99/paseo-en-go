//go:build !solution

// Package images draws a picture by implementing the image.Image interface.
//
// Write this type and these methods in this file:
//
//	type Image struct {
//		Width, Height int
//	}
//
//	func (m Image) ColorModel() color.Model
//	func (m Image) Bounds() image.Rectangle
//	func (m Image) At(x, y int) color.Color
//
// Image is a picture of Width by Height pixels.
//
// ColorModel returns the colour model of the image: color.RGBAModel.
//
// Bounds returns the rectangle that holds the image, starting at (0, 0).
// Example: Image{Width: 10, Height: 20}.Bounds() returns
// image.Rect(0, 0, 10, 20).
//
// At returns the colour of the pixel at (x, y).
// Example: At(10, 20) returns color.RGBA{10, 20, 255, 255}.
//
// Hint: import "image" and "image/color". At returns
// color.RGBA{uint8(x), uint8(y), 255, 255}.
package images
