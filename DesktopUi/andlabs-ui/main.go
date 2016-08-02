package main

import (
	"github.com/andlabs/ui"
	"time"
	"log"
	"strconv"
	"math/rand"
	"net/http"
	"image"
	"image/draw"
	_ "image/png"
	_ "image/jpeg"
)

func LoadLoremImage() (*image.RGBA, error) {
	//  Try to load image
	client := http.Client{}
	reqImg, err := client.Get("http://lorempixel.com/400/300/")
	if err != nil {
		return nil, err
	}
	source, format, err := image.Decode(reqImg.Body)
	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	log.Println("Image format ", format)
	return rgba, nil
}

func initGUI() {

	ff := ui.ListFontFamilies()
	for i := 0; i < ff.NumFamilies(); i++ {
		log.Printf("%3d. Font family '%s'\n", i + 1, ff.Family(i))
	}

	window := ui.NewWindow("Привет мир!", 800, 480, false)
	window.SetMargined(true)

	progress := ui.NewProgressBar()
	labelTime := ui.NewLabel("")

	labelInfo := ui.NewLabel("Info")
	labelInfoButtonHandler := func(b *ui.Button) {
		labelInfo.SetText("Click button " + b.Text())
	}

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	hbox.Append(func() *ui.Box {
		box := ui.NewVerticalBox()
		box.SetPadded(true)

		box.Append(func() *ui.Button {
			button := ui.NewButton("Button 1")
			button.OnClicked(labelInfoButtonHandler)
			return button
		}(), false)
		box.Append(func() *ui.Button {
			button := ui.NewButton("Button 2")
			button.OnClicked(labelInfoButtonHandler)
			return button
		}(), false)
		box.Append(func() *ui.Button {
			button := ui.NewButton("Button 3")
			button.OnClicked(labelInfoButtonHandler)
			return button
		}(), false)
		box.Append(ui.NewHorizontalSeparator(), false)

		label := ui.NewLabel("It's all")
		box.Append(label, true)

		box.Append(func() *ui.Button {
			button := ui.NewButton("Exit")
			button.OnClicked(func(*ui.Button) {
				ui.Quit()
			})
			return button
		}(), false)

		return box
	}(), false)

	areaHandler := NewHistogramAreaHandler(20)
	area := ui.NewArea(areaHandler)

	hbox.Append(func() *ui.Box {
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		box.Append(labelInfo, false)
		box.Append(labelTime, false)

		tab := ui.NewTab()
		tab.Append("Histogram demo", area)
		tab.Append("Controls demo", func() *ui.Box {
			box := ui.NewVerticalBox()
			box.SetPadded(true)

			box.Append(ui.NewEntry(), false)
			box.Append(ui.NewCheckbox("Check it"), false)
			box.Append(func() *ui.RadioButtons {
				radio := ui.NewRadioButtons()
				radio.Append("Radio button 1")
				radio.Append("Radio button 2")
				radio.Append("Radio button 3")
				return radio
			}(), false)
			box.Append(func() *ui.Group {
				combo := ui.NewCombobox()
				combo.Append("First")
				combo.Append("Second")
				combo.Append("Third")
				combo.Append("Fourth")
				combo.OnSelected(func(cb *ui.Combobox) {
					ui.MsgBoxError(window, "OnSelected", "Line #" + strconv.Itoa(cb.Selected() + 1))
				})
				group := ui.NewGroup("Can't get text, only index")
				group.SetChild(combo)
				return group
			}(), false)
			box.Append(ui.NewSlider(0, 100), false)
			box.Append(ui.NewSpinbox(0, 10), false)
			box.Append(ui.NewDatePicker(), false)
			box.Append(ui.NewDateTimePicker(), false)

			return box
		}())
		tab.Append("Tab 3", ui.NewLabel("At tab 3"))
		box.Append(tab, true)

		box.Append(progress, false)

		return box
	}(), true)

	window.SetChild(hbox)

	window.OnClosing(func(*ui.Window) bool {
		log.Println("Window close")
		ui.Quit()
		return true
	})
	window.Show()

	progressCounter := 0
	progressTicker := time.NewTicker(time.Millisecond * 50)
	go func() {
		for _ = range progressTicker.C {
			// Что бы записать значение в виджет используем потокобезопасный вызов
			ui.QueueMain(func() {
				progress.SetValue(progressCounter)
			})
			progressCounter++
			if progressCounter > 100 {
				progressCounter = 0
			}
		}
	}()

	timeTicker := time.NewTicker(time.Millisecond * 10)
	go func() {
		for t := range timeTicker.C {
			// Что бы записать значение в виджет используем потокобезопасный вызов
			ui.QueueMain(func() {
				labelTime.SetText(t.Format(time.StampMilli))
			})
		}
	}()

	hystogrammTicker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for _ = range hystogrammTicker.C {
			// Что бы записать значение в виджет используем потокобезопасный вызов
			ui.QueueMain(func() {
				areaHandler.Push(rand.Intn(100))
				area.QueueRedrawAll()
			})
		}
	}()

	log.Println("InitGUI done")
}

func main() {
	err := ui.Main(initGUI)
	if err != nil {
		panic(err)
	}
	log.Println("The end")
}
