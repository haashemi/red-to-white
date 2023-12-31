package main

import (
	"image"
	"image/color"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/haashemi/painter"
)

func main() {
	// just to make sure that this directory exists
	os.Mkdir("export", os.ModePerm)

	err := filepath.WalkDir("input", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		rawImage, err := painter.Decode(f)
		if err != nil {
			return err
		}

		newImage := image.NewNRGBA(image.Rect(0, 0, rawImage.Bounds().Dx(), rawImage.Bounds().Dy()))

		for x := 0; x < rawImage.Bounds().Dx(); x++ {
			for y := 0; y < rawImage.Bounds().Dy(); y++ {
				rawColor := rawImage.At(x, y).(color.RGBA)

				newColor := color.NRGBA{255, 255, 255, rawColor.R}

				newImage.SetNRGBA(x, y, newColor)
			}
		}

		exportPath := strings.Replace(path, "input", "export", 1)
		return painter.SavePNG(newImage, exportPath)
	})

	log.Fatalln("Application closed.", err)
}
