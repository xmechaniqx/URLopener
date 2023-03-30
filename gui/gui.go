package gui

import (
	"strings"
	"time"

	"github.com/tadvi/winc"
)

func Gui() (url, name string) {
	// var url, name string
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(350, 300) // (width, height)
	mainWindow.SetText("URL Opener")

	edt := winc.NewEdit(mainWindow)
	edt.SetPos(80, 40)
	edt.SetText("Введите адрес ссылки")
	drn := winc.NewEdit(mainWindow)
	drn.SetPos(80, 80)
	drn.SetText("Введите будущее название файла")
	// Most Controls have default size unless SetSize is called.
	// edt.SetText("Введите адрес ссылки")

	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Генерировать")
	btn.SetPos(120, 120) // (x, y)
	btn.SetSize(100, 70) // (width, height)\
	btn.OnClick().Bind(func(e *winc.Event) {
		url = (edt.Text())
		url = strings.Replace(url, "https://", "", -1)
		var arrString []string

		arrString = append(arrString, (drn.Text()))
		name = strings.Join(arrString, " ")

		// name = strconv.Quote(name)
		// name = name + ".exe"
		// name = (drn.Text())
		winc.Exit()
		// fmt.Println(name)
	})

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop() // Must call to start event loop.
	time.Sleep(3 * time.Second)
	return url, name
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
