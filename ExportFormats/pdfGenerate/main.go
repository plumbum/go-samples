package main

import (
	"github.com/jung-kurt/gofpdf"
	"log"
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

	var x1, y1 float64 = 100, 250
	var xc, yc float64 = 25, 225
	var x2, y2 float64 = 200, 150
	pdf.SetFillColor(192, 192, 192)
	pdf.SetDrawColor(0, 192, 0)
	pdf.Curve(x1, y1, xc, yc, x2, y2, "DF")
	pdf.SetDrawColor(192, 0, 0)
	pdf.Line(x1, y1, xc, yc)
	pdf.Line(x2, y2, xc, yc)

	// pdf.Arc(60, 60, 70, 70, 1.0, 0.5, 2.0, "DF")
	poly := []gofpdf.PointType{
		{150, 10},
		{200, 10},
		{150, 110},
		{200, 110},
	}
	pdf.SetDrawColor(0, 192, 0)
	pdf.Beziergon(poly, "DF");
	pdf.SetDrawColor(192, 0, 0)
	pdf.Polygon(poly, "D")

	err = pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal("error: Write PDF ", err)
	}
}
