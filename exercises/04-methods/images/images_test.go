package images

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

// Image must satisfy image.Image; if it does not, this line does not compile.
var _ image.Image = Image{}

func TestColorModel(t *testing.T) {
	m := Image{Width: 4, Height: 4}
	if got := m.ColorModel(); got != color.RGBAModel {
		t.Errorf("ColorModel() = %v, want color.RGBAModel", got)
	}
}

func TestBounds(t *testing.T) {
	tests := []struct {
		m    Image
		want image.Rectangle
	}{
		{m: Image{Width: 256, Height: 256}, want: image.Rect(0, 0, 256, 256)},
		{m: Image{Width: 10, Height: 20}, want: image.Rect(0, 0, 10, 20)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.m.Bounds(); got != tt.want {
				t.Errorf("%#v.Bounds() = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}

func TestAt(t *testing.T) {
	m := Image{Width: 256, Height: 256}
	tests := []struct {
		x, y int
		want color.Color
	}{
		{x: 0, y: 0, want: color.RGBA{0, 0, 255, 255}},
		{x: 10, y: 20, want: color.RGBA{10, 20, 255, 255}},
		{x: 255, y: 128, want: color.RGBA{255, 128, 255, 255}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := m.At(tt.x, tt.y); got != tt.want {
				t.Errorf("At(%d, %d) = %v, want %v", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestEncodePNG(t *testing.T) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, Image{Width: 16, Height: 16}); err != nil {
		t.Errorf("png.Encode(Image{Width: 16, Height: 16}) = %v, want <nil>", err)
	}
}
