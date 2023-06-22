package window

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zhandos717/gateway/app/database"
	"github.com/zhandos717/gateway/app/models"
	"net"
	"strconv"
)

func Run() {

	db := database.DB

	ipAddres := models.IpAddres{}

	if result := db.Where("id = ?", 1).First(&ipAddres); result.Error != nil {
		db.Save(&models.IpAddres{ID: 1, Ip: ipAddres.Ip, Port: ipAddres.Port})

		ipAddres = models.IpAddres{Ip: "192.168.0.1", Port: 8081}
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

			if !isValidIP4(ip.Text) {
				dialog.ShowInformation("Ошибка", "Поле Ip не верный", w)
				return
			}

			if err != nil {
				dialog.ShowInformation("Ошибка", "Поле порт должно быть строкой", w)
				return
			}

			db.Save(&models.IpAddres{ID: 1, Ip: ip.Text, Port: uint32(port)})
		},
	}

	form.SubmitText = "Сохранить"

	form.Append("Порт: ", port)

	w.SetContent(container.NewVBox(
		form,
	))

	w.ShowAndRun()
}

func isValidIP4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	ip4 := parsedIP.To4()
	if ip4 == nil {
		return false
	}

	return true
}
