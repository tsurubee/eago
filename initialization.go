package eago

import (
	"math/rand"
	"time"
	"image"
)

func InitFloatVector(n int, upperLimit, lowerLimit float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	vector := make([]float64, n)
	for i := range vector {
		vector[i] = lowerLimit + rand.Float64()*(upperLimit-lowerLimit)
	}
	return vector
}

func InitRGBAImage(img *image.RGBA) image.RGBA {
	rand.Seed(time.Now().UnixNano())
	pix := make([]uint8, len(img.Pix))
	rand.Read(pix)
	return image.RGBA{
		Pix:    pix,
		Stride: img.Stride,
		Rect:   img.Rect,
	}
}
