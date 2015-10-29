package main

import (
	"log"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
)

func main() {
	var err error

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")

	number := "20151029"
	// Code128 barcode
	bcode, err := code128.Encode(number)
	if err != nil {
		log.Fatal("error: Barcode ", err)
	}
	key := barcode.Register(bcode)
	barcode.Barcode(pdf, key, 10, 50, 100, 15, false)
	pdf.SetDrawColor(255, 0, 0)
	pdf.Rect(10, 50, 100, 15, "D")
	pdf.SetXY(10, 65)
	pdf.SetFont("Arial", "", 12)
	pdf.WriteAligned(100, 5, number, "C")

	// Code128 barcode short
	key128 := barcode.RegisterCode128(pdf, number)
	barcode.Barcode(pdf, key128, 10, 80, 100, 15, false)
	pdf.Rect(10, 80, 100, 15, "D")

	// QR-code
	keyqr := barcode.RegisterQR(pdf, "tuxotronic.org", qr.H, qr.Unicode)
	barcode.Barcode(pdf, keyqr, 10, 100, 40, 40, false)

	err = pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal("error: Write PDF ", err)
	}
}
