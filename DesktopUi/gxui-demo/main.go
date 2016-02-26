package main

import (
	"time"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"fmt"
	"os"
	"image"
	"image/draw"
	_ "image/jpeg"
)

/*
func main() {
}
*/

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	font, err := driver.CreateFont(gxfont.Default, 75)
	if err != nil {
		panic(err)
	}

	window := theme.CreateWindow(800, 480, "Привет")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray50))
	window.SetBorderPen(gxui.Pen{Width:5, Color: gxui.Yellow})

	f, err := os.Open("wallpaper.jpg")
	if err != nil {
		fmt.Printf("Failed to open image %v\n", err)
		os.Exit(1)
	}

	source, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Failed to open image %v\n", err)
		os.Exit(2)
	}

	wallpaper := theme.CreateImage()
	window.AddChild(wallpaper)

	// Copy the image to a RGBA format before handing to a gxui.Texture
	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	wallpaper.SetTexture(texture)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	window.AddChild(layout)

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Здравствуй мир")
	label.OnMouseMove(func(e gxui.MouseEvent) {
		fmt.Printf("X=%d; Y=%d\n", e.Point.X, e.Point.Y)
	})
	layout.AddChild(label)

	lTimer := theme.CreateLabel()
	lTimer.SetFont(font)
	lTimer.SetColor(gxui.Green30)
	layout.AddChild(lTimer)

	button := theme.CreateButton()
	button.SetText("Exit")
	button.SetPadding(math.Spacing{20, 10, 20, 10})
	button.SetMargin(math.Spacing{20, 10, 20, 10})
	button.OnClick(func(e gxui.MouseEvent) {
		window.Close()
	})
	layout.AddChild(button)

	ticker := time.NewTicker(time.Millisecond * 30)
	go func() {
		phase := float32(0)
		for _ = range ticker.C {
			c := gxui.Color{
				R: 0.75 + 0.25 * math.Cosf((phase + 0.000) * math.TwoPi),
				G: 0.75 + 0.25 * math.Cosf((phase + 0.333) * math.TwoPi),
				B: 0.75 + 0.25 * math.Cosf((phase + 0.666) * math.TwoPi),
				A: 0.50 + 0.50 * math.Cosf(phase * 10),
			}
			phase += 0.01
			driver.Call(func() {
				label.SetColor(c)
			})
		}
	}()

	ticker2 := time.NewTicker(time.Millisecond * 100)
	go func() {
		for t := range ticker.C {
			driver.Call(func() {
				lTimer.SetText(t.Format(time.Stamp))
			})
		}
	}()

	window.OnClose(ticker.Stop)
	window.OnClose(ticker2.Stop)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}