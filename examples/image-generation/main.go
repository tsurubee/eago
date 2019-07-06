package main

import (
	"github.com/tsurubee/eago"
	"os"
	"log"
	"image"
	"bytes"
	"image/png"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
)

type Image image.RGBA

func (I Image) Initialization() eago.Genome {
	return Image(eago.InitRGBAImage(loadImage("./monalisa.png")))
}

func (I Image) Fitness() float64 {
	return imageSimilarity(Image(copyImage(loadImage("./monalisa.png"))), I)
}

func (I Image) Mutation() {
	//eago.AddNormalFloat(I, 0.5)
}

func (I Image) Crossover(X eago.Genome) eago.Genome {
	pix := make([]uint8, len(I.Pix))
	child := image.RGBA{
			Pix:    pix,
			Stride: I.Stride,
			Rect:   I.Rect,
	}
	mid := rand.Intn(len(I.Pix))
	for i := 0; i < len(I.Pix); i++ {
		if i > mid {
			child.Pix[i] = I.Pix[i]
		} else {
			child.Pix[i] = X.(Image).Pix[i]
		}

	}
	return Image(child)
}

func (I Image) PrintImage()  {
	img := createRGBAImage(I)
	printImage(img.SubImage(img.Rect))
}

func imageSimilarity(i1, i2 Image) float64 {
	sim := 0.
	for i := 0; i < len(i1.Pix); i++ {
		sim += euclideanDistance(i1.Pix[i], i2.Pix[i])
	}
	return math.Sqrt(sim)
}

func euclideanDistance(x, y uint8) float64 {
	difference := float64(x - y)
	return math.Sqrt(difference * difference)
}

func copyImage(img *image.RGBA) image.RGBA {
	return image.RGBA{
		Pix:    img.Pix,
		Stride: img.Stride,
		Rect:   img.Rect,
	}
}

func createRGBAImage(img Image) image.RGBA {
	return image.RGBA{
		Pix:    img.Pix,
		Stride: img.Stride,
		Rect:   img.Rect,
	}
}

func loadImage(filePath string) *image.RGBA {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	return img.(*image.RGBA)
}

func printImage(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Printf("\x1b]1337;File=inline=1:%s\a\n", imgBase64Str)
}

func main() {
	var img Image
	ga := eago.NewGA(eago.GAConfig{
		PopulationSize: 100,
		NGenerations:   10,
		CrossoverRate:  0.8,
		MutationRate:   0.01,
		ParallelEval:   true,
	})
	ga.PrintCallBack = func() {
		ga.BestIndividual.Chromosome.(Image).PrintImage()
	}
	if err := ga.Minimize(img); err != nil {
		log.Fatal(err)
	}
}
