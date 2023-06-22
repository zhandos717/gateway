package window

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

type IpAddres struct {
	ID   uint `gorm:"primaryKey"`
	Ip   string
	Port uint32
}

func Run() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if db.AutoMigrate(&IpAddres{}) != nil {
		return
	}

	ipAddres := IpAddres{}

	if result := db.Where("id = ?", 1).First(&ipAddres); result.Error != nil {
		ipAddres = IpAddres{Ip: "192.168.0.1", Port: 8081}
	}

	a := app.New()

	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("Первое приложнеие")
	w.Resize(fyne.NewSize(350, 180))
	w.CenterOnScreen()

	ip := widget.NewEntry()
	port := widget.NewEntry()

	ip.SetText(ipAddres.Ip)
	port.SetText(fmt.Sprintf("%v", ipAddres.Port))

	port.SetOnValidationChanged(func(err error) {

	})

	form := &widget.Form{
		Items: []*widget.FormItem{{Text: "Ip-Адресс: ", Widget: ip}},

		OnSubmit: func() {

			port, err := strconv.Atoi(port.Text)

			if err != nil {
				dialog.ShowInformation("Ошибка", "Поле порт должно быть строкой", w)
				return
			}

			db.Save(&IpAddres{ID: 1, Ip: ip.Text, Port: uint32(port)})
		},
	}

	form.SubmitText = "Сохранить"

	form.Append("Порт: ", port)

	w.SetContent(container.NewVBox(
		form,
	))

	w.ShowAndRun()
}
