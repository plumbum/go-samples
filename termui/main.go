package main

import ui "github.com/gizak/termui"

func main() {

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	w11 := ui.NewPar("Hello world")
	w11.Height = 10
	w11.Border.Label = "Hello"
	w11.Border.LabelFgColor = ui.ColorGreen

	w12 := ui.NewPar("first")
	w12.Height = 20

	w2 := ui.NewPar("second")
	w2.Height = 20

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 0, w11),
			ui.NewCol(6, 0, w12)),
		ui.NewRow(
			ui.NewCol(12, 0, w2)))

	ui.Body.Align()

	ui.Render(ui.Body)

	<- ui.EventCh()

}
