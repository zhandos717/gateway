package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Первое приложнеие")

	label := widget.NewLabel("Привет")

	w.SetContent(container.NewVBox(
		label, widget.NewButton("hi", func() {
			label.SetText("world")
		})))

	w.ShowAndRun()
}
