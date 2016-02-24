package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "mm", PageSize: gopdf.Rect{W: 210.0, H: 297.0}}) //595.28pt, 841.89pt = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("DROID", "DroidSerif-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("DROID", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "Hello world!")
	pdf.SetY(30.0)
	pdf.Cell(nil, "Привет мир!")

	pdf.SetGrayStroke(0.5)
	pdf.Oval(10, 200, 200, 250)

	pdf.WritePdf("hello.pdf")
}
