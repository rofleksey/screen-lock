package main

import (
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"log"
	"os"
	"screenLock/util"
	"screenLock/widgets"
	"strconv"
	"time"
)

var passFlag = flag.String("p", "", "password")
var focusFlag = flag.Bool("f", false, "keep focus")
var msgFlag = flag.String("m", "enter password", "startup message")
var bgColorFlag = flag.String("bg", "FF0000FF", "background color (rgba hex)")
var textColorFlag = flag.String("fg", "00000011", "text color (rgba hex)")

var currentPassword = ""

func main() {
	flag.Parse()

	passNum, err := strconv.Atoi(*passFlag)
	if *passFlag == "" || err != nil || passNum < 0 {
		log.Println("Invalid password format, should be a number")
		flag.Usage()
		os.Exit(1)
	}

	if *passFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	bgColor, err := util.ParseColor(*bgColorFlag)
	if err != nil {
		log.Printf("Failed to parse background color: %v", err)
		flag.Usage()
		os.Exit(1)
	}

	textColor, err := util.ParseColor(*textColorFlag)
	if err != nil {
		log.Printf("Failed to parse text color: %v", err)
		flag.Usage()
		os.Exit(1)
	}

	a := app.New()
	w := a.NewWindow("Touchverse Screen Lock")
	w.SetFullScreen(true)

	if *focusFlag {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				w.RequestFocus()
			}
		}()
	}

	topLabel := widgets.NewPasswordLabel(*msgFlag, textColor)

	digits := make([]fyne.CanvasObject, 0, 9)

	for i := range 9 {
		digitStr := fmt.Sprint(i + 1)

		digit := widgets.NewDigit(digitStr, textColor, func() {
			currentPassword += digitStr

			if currentPassword == *passFlag {
				os.Exit(0)
			}

			if len(currentPassword) > len(*passFlag) {
				currentPassword = ""
			}

			topLabel.SetText(currentPassword)
		})

		digits = append(digits, digit)
	}

	digitsContent := container.New(layout.NewGridLayout(3), digits...)
	mainColumn := container.New(layout.NewBorderLayout(topLabel, nil, nil, nil), topLabel, digitsContent)

	background := canvas.NewRectangle(bgColor)
	rootContent := container.NewStack(background, mainColumn)

	w.SetContent(rootContent)
	w.ShowAndRun()
}
