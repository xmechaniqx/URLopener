package main

import (
	"log"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func main() {
	// compile.Compile()
	url := "redactUrl"

	// Name := compile.Name
	Open(url, BrowserRegSearcherChrome(), BrowserRegSearcherYandex())

}

func Open(url string, regChrome, regYandex int) error {
	var cmd string
	var argEx, argYa, argChr []string

	yandex := "browser"
	chrome := "chrome"
	cmd = "cmd"
	// args = []string{"/c", yandex, url}
	argYa = []string{"/c", "start", yandex, url}
	argChr = []string{"/c", "start", chrome, url}
	argEx = []string{"/c", "start", url}
	// args = append(args, url)

	if regYandex == 1 {
		err := exec.Command(cmd, argYa...).Start()
		if err != nil {
			log.Fatal(err, "regYandex")
			return err
		}
	}
	if regYandex == 0 && regChrome == 1 {
		err := exec.Command(cmd, argChr...).Start()
		if err != nil {
			log.Fatal(err, "regChrome")
			return err
		}
	}

	if regYandex == 0 && regChrome == 0 {
		exec.Command(cmd, argEx...).Start()
	}

	var err error
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
func BrowserRegSearcherChrome() int {
	var regChrome int
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Google`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err, "Ошибка открытия ключа")
	}
	defer k.Close()
	regFindingValueChrome, err := k.ReadSubKeyNames(0)
	if err != nil {
		log.Fatal(err, "Ошибка чтения значений строки реестра")
	}
	for _, v := range regFindingValueChrome {
		if v == "Chrome" {
			regChrome = 1
		}
	}
	// regChrome = 0
	return regChrome
}
func BrowserRegSearcherYandex() int {
	var regYandex int
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Yandex`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err, "Ошибка открытия ключа")
	}
	defer k.Close()
	regFindingValueChrome, err := k.ReadSubKeyNames(0)
	if err != nil {
		log.Fatal(err, "Ошибка чтения значений строки реестра")
	}
	for _, v := range regFindingValueChrome {
		if v == "YandexBrowser" {
			regYandex = 1
		}
	}
	// regYandex = 0
	return regYandex
}
