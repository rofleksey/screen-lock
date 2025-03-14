package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type PasswordLabel struct {
	widget.BaseWidget
	text *canvas.Text
}

func NewPasswordLabel(startText string, color color.RGBA) *PasswordLabel {
	text := canvas.NewText(startText, color)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 50

	label := &PasswordLabel{
		text: text,
	}
	label.ExtendBaseWidget(label)

	return label
}

func (w *PasswordLabel) SetText(newText string) {
	w.text.Text = newText
	w.text.Refresh()
}

func (w *PasswordLabel) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(w.text)
}
