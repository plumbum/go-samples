package main

import (
	"github.com/andlabs/ui"
	"reflect"
	"fmt"
	"time"
)

var w ui.Window

type MyTable struct {
	Num int
	Text string
}

func initGUI() {

	lb := ui.NewLabel("My label")

	tf := ui.NewTextField()
	tf.SetText("Hello world")
	bt := ui.NewButton("Add")

	hstack := ui.NewHorizontalStack(tf, bt)
	hstack.SetStretchy(0)
	hstack.SetPadded(true)

	idx := 1
	lines := []MyTable{}

	table := ui.NewTable(reflect.TypeOf(MyTable{}))
	table.OnSelected(func() {
		fmt.Println("Select line:", table.Selected())
		table.RLock() // Блокируем таблицу на время чтения
		selIdx := table.Selected()
		d := table.Data().(*[]MyTable)
		tf.SetText((*d)[selIdx].Text)
		table.RUnlock()
	})

	bt.OnClicked(func() {
		table.Lock() // Блокируем таблицу на время обновления строк
		lines = append(lines, MyTable{
			Num: idx,
			Text: tf.Text(),
		})
		d := table.Data().(*[]MyTable)
		*d = lines
		table.Unlock()
		idx++
	})

	lbTime := ui.NewLabel("")

	btExit := ui.NewButton("Exit")

	hstack2 := ui.NewHorizontalStack(btExit)

	stack := ui.NewVerticalStack(lb, hstack, table, lbTime, hstack2)
	stack.SetPadded(true)
	stack.SetStretchy(2)

	w = ui.NewWindow("My window", 640, 480, stack)
	w.SetMargined(true)
	w.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	w.Show()

	tick := time.NewTicker(time.Second)
	go func() {
		for t := range tick.C {
			table.Lock() // Блокируем таблицу на время обновления строк
			lines = append(lines, MyTable{
				Num: 0,
				Text: t.String(),
			})
			d := table.Data().(*[]MyTable)
			*d = lines
			table.Unlock()
		}
	}()

	btExit.OnClicked(func () {
		tick.Stop()
		ui.Stop()
	})

}

func main() {
	go ui.Do(initGUI)
	err := ui.Go()
	if err != nil {
		panic(err)
	}
}
