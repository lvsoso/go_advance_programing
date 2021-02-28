package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")

	display := widget.NewEntry()
	display.MultiLine = true
	myWindow.SetContent(display)
	myWindow.ShowAndRun()
}