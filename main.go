package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IpAddres struct {
	gorm.Model
	Ip          string
	Port        uint16
	Description string
}

func main() {

	windowApp()
}

func windowApp() {

	//var ipAddress []IpAddres

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&IpAddres{})
	db.Create(&IpAddres{Ip: "192.168.0.1", Port: 8081})

	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Первое приложнеие")
	w.Resize(fyne.NewSize(320, 100))
	w.CenterOnScreen()

	var ipAddres IpAddres

	db.First(&ipAddres, 1)

	ip := widget.NewEntry()
	port := widget.NewEntry()

	ip.SetText(ipAddres.Ip)
	port.SetText(fmt.Sprintf("%s", ipAddres.Port))

	form := &widget.Form{
		Items: []*widget.FormItem{{Text: "Ip-Адресс", Widget: ip}},
		OnSubmit: func() {
			// optional, handle form submission
			log.Println("Ip-Адресс:", ip.Text)
			log.Println("Порт:", port.Text)
		},
	}

	form.SubmitText = "Сохранить"

	form.Append("Порт", port)

	w.SetContent(container.NewVBox(
		form,
	))

	w.ShowAndRun()
}
