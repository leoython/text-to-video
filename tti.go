package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/leoython/text-to-video/internal"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
)

const (
	defaultSize float64 = 45
	defaultSpacing float64 = 1.5
)

type Imager struct {}

func(i *Imager) genBaseImage(filename, text string) error {
	m := image.NewRGBA(image.Rect(0, 0, 1600, 900))
	draw.Draw(m, m.Bounds(), image.Black, image.ZP, draw.Src)
	
	if err := internal.CreateFolderIfNotExists("image"); err != nil {
		return err
	}

	// Read the font data.
	fontBytes, err := ioutil.ReadFile("./front/SourceHanSans-Bold.ttf")
	if err != nil {
		return err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(defaultSize)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(image.White)

	pt := freetype.Pt(350, 200+int(c.PointToFixed(defaultSize)>>6))
	if _, err := c.DrawString(text, pt); err != nil {
		return err
	}

	fd, err := os.Create(fmt.Sprintf("image/%s.jpeg", filename))
	if err != nil {
		return err
	}
	err = jpeg.Encode(fd, m, nil)
	if err != nil {
		return err
	}
	return nil
}