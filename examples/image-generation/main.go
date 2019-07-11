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
	"time"
)

var targetPath = "./gopher.png"

type Image image.RGBA

func (I Image) Initialization() eago.Genome {
	return Image(eago.InitRGBAImage(loadImage(targetPath)))
}

func (I Image) Fitness() float64 {
	return imageSimilarity(Image(copyImage(loadImage(targetPath))), I)
}

func (I Image) Mutation() {
	mutationRate := 0.0005
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(I.Pix); i++ {
		if rand.Float64() < mutationRate {
			I.Pix[i] = uint8(rand.Intn(255))
		}
	}
}

func (I Image) Crossover(X eago.Genome) eago.Genome {
	rand.Seed(time.Now().UnixNano())
	pix := make([]uint8, len(I.Pix))
	child := &image.RGBA{
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
	return Image(copyImage(child))
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
	return difference * difference
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

func NRGBAtoRGBA(img *image.NRGBA) *image.RGBA {
	return &image.RGBA{
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
	return NRGBAtoRGBA(img.(*image.NRGBA))
}

func saveImage(filePath string, rgba image.RGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(imgFile, rgba.SubImage(rgba.Rect))
}

func printImage(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Printf("\x1b]1337;File=inline=1:%s\a\n", imgBase64Str)
}

func main() {
	var img Image
	target := loadImage(targetPath)
	printImage(target.SubImage(target.Rect))
	start := time.Now()

	ga := eago.NewGA(eago.GAConfig{
		PopulationSize: 600,
		NGenerations:   20000,
		CrossoverRate:  0.9,
		MutationRate:   0.9,
		ParallelEval:   true,
	})
	ga.Selector = eago.Tournament{
		NContestants: 40,
	}
	ga.PrintCallBack = func() {
		if ga.Generations % 100 == 0 {
			sofar := time.Since(start)
			fmt.Printf("Generation %3d | Elapsed time: %s | Fitness=%.3f\n", ga.Generations, sofar, ga.BestIndividual.Fitness)
			ga.BestIndividual.Chromosome.(Image).PrintImage()
			saveImage(fmt.Sprintf("./results/gopher_%v.png", ga.Generations), createRGBAImage(ga.BestIndividual.Chromosome.(Image)))
		}
	}
	if err := ga.Minimize(img); err != nil {
		log.Fatal(err)
	}
}
