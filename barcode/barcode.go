package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/code128"
	"golang.org/x/image/draw"
	"os"
	"image"
	"image/png"
	"image/color"
	"time"
)

func main() {
	s := time.Now().Format("2006-01-02 15:04")
	writePng("qrcode.png", addBounds(createQrCode(s, image.Pt(8,8))))
	writePng("code128.png", addBounds(createCode128(s, image.Pt(2, 64))))
}

func createCode128(data string, scale image.Point) image.Image {
	bar, err := code128.Encode(data)
	chk(err)
	imgWidth := bar.Bounds().Dx()*scale.X
	imgHeight := bar.Bounds().Dy()*scale.Y
	barScale, err := barcode.Scale(bar, imgWidth, imgHeight)
	chk(err)
	return barScale
}

func createQrCode(data string, scale image.Point) image.Image {
	bar, err := qr.Encode(data, qr.M, qr.Auto)
	chk(err)
	imgWidth := bar.Bounds().Dx()*scale.X
	imgHeight := bar.Bounds().Dy()*scale.Y
	barScale, err := barcode.Scale(bar, imgWidth, imgHeight)
	chk(err)
	return barScale
}

func addBounds(img image.Image) image.Image {
	imgWidth := img.Bounds().Dx()+16
	imgHeight := img.Bounds().Dy()+16
	newImg := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)
	draw.Copy(newImg, image.Pt(8, 8), img, img.Bounds(), draw.Src, nil)
	return newImg

}

func writePng(fileName string, img image.Image) {
	var err error
	wr, err := os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY, 0666)
	chk(err)
	defer wr.Close()
	err = png.Encode(wr, img)
	chk(err)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
