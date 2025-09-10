package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func MergeImages(inputDir, outputFile string) error {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return err
	}

	var images []image.Image
	maxWidth := 0

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(f.Name()))
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			continue
		}

		path := filepath.Join(inputDir, f.Name())
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		var img image.Image
		if ext == ".png" {
			img, err = png.Decode(file)
		} else {
			img, err = jpeg.Decode(file)
		}
		file.Close()

		if err != nil {
			return err
		}

		if img.Bounds().Dx() > maxWidth {
			maxWidth = img.Bounds().Dx()
		}
		images = append(images, img)
	}

	if len(images) == 0 {
		return fmt.Errorf("no images found in %s", inputDir)
	}

	var resized []image.Image
	totalHeight := 0
	for _, img := range images {
		newImg := resize.Resize(uint(maxWidth), 0, img, resize.Lanczos3)
		resized = append(resized, newImg)
		totalHeight += newImg.Bounds().Dy()
	}

	out := image.NewRGBA(image.Rect(0, 0, maxWidth, totalHeight))

	y := 0
	for _, img := range resized {
		draw.Draw(out, image.Rect(0, y, maxWidth, y+img.Bounds().Dy()), img, image.Point{}, draw.Over)
		y += img.Bounds().Dy()
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, out)
}
