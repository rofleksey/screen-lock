package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Digit struct {
	widget.BaseWidget
	text     string
	color    color.RGBA
	callback func()
}

func NewDigit(text string, color color.RGBA, callback func()) *Digit {
	digit := &Digit{
		text:     text,
		color:    color,
		callback: callback,
	}
	digit.ExtendBaseWidget(digit)
	return digit
}

func (w *Digit) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(w.text, w.color)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 50

	return widget.NewSimpleRenderer(text)
}

func (w *Digit) Tapped(_ *fyne.PointEvent) {
	w.callback()
}
