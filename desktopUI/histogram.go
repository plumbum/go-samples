package main
import (
"math/rand"
"time"
"github.com/andlabs/ui"
"strconv"
"log"
)

type HistogramAreaHandler struct {
	Data []int
}

func NewHistogramAreaHandler(bars int) *HistogramAreaHandler {
	a := new(HistogramAreaHandler)

	rand.Seed(time.Now().Unix())
	a.Data = make([]int, bars)
	for i, _ := range a.Data {
		a.Data[i] = rand.Intn(100)
	}
	return a
}

func (a *HistogramAreaHandler) Push(value int) {
	a.Data = append(a.Data[1:], value)
}

func (a HistogramAreaHandler) Draw(area *ui.Area, dp *ui.AreaDrawParams) {

	fntNum := ui.LoadClosestFont(&ui.FontDescriptor{Family:"Droid Sans", Size: 10, Weight: ui.TextWeightBold})
	fntVolume := ui.LoadClosestFont(&ui.FontDescriptor{Family:"Droid Sans", Size: 12, Italic: ui.TextItalicItalic})

	brush := &ui.Brush{Type: ui.Solid, R:0.1, G:0.5, B:0.7, A:1  }

	columnWidth := dp.AreaWidth / float64(len(a.Data))
	barWidth := columnWidth * 0.8

	for i, volume := range a.Data {

		x1 := (float64(i) +0.1) * columnWidth;
		x2 := x1 + barWidth
		y1 := dp.AreaHeight - 16;
		y2 := y1 - (dp.AreaHeight - 32.0) * float64(volume) / 100.0;

		layoutNum := ui.NewTextLayout(strconv.Itoa(i), fntNum, columnWidth)
		w1, _ := layoutNum.Extents()
		dp.Context.Text(x1 + (barWidth - w1) / 2, y1, layoutNum)

		layoutVolume := ui.NewTextLayout(strconv.Itoa(volume), fntVolume, columnWidth)
		w2, h2 := layoutVolume.Extents()
		dp.Context.Text(x1 + (barWidth - w2) / 2, y2 - h2, layoutVolume)

		dp.Context.Save()
		p := ui.NewPath(ui.Winding)
		p.NewFigure(x1, y1)
		p.LineTo(x2, y1)
		p.LineTo(x2, y2)
		p.LineTo(x1, y2)
		p.CloseFigure()
		p.End()
		dp.Context.Fill(p, brush)
		dp.Context.Restore()
	}

}

func (a *HistogramAreaHandler) MouseEvent(area *ui.Area, me *ui.AreaMouseEvent) {
	if me.Up == 1 {
		log.Printf("Mouse pressed at [%d:%d]\n", int(me.X), int(me.Y))
		a.Push(rand.Intn(100))
		area.QueueRedrawAll()
	}
}

func (a HistogramAreaHandler) MouseCrossed(area *ui.Area, left bool) {
	log.Printf("Mouse crossed area %v\n", left)
}

func (a HistogramAreaHandler) KeyEvent(area *ui.Area, ke *ui.AreaKeyEvent) bool {
	if ke.Up {
		log.Printf("Key pressed %v\n", ke.Key)
	}
	return false
}

func (a HistogramAreaHandler) DragBroken(area *ui.Area) {
	log.Println("Drag broken")
}

