package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Первое приложнеие")
	label := widget.NewLabel("Введите IP терминала ")

	w.Resize(fyne.NewSize(400, 320))

	entry := widget.NewEntry()

	btn := widget.NewButton("Сохранить", func() {
		ip := entry.Text

		fmt.Println(ip)
	})

	w.SetContent(container.NewVBox(
		entry,
		label,
		btn,
	))

	w.ShowAndRun()
}

func btnHandler() {

}
