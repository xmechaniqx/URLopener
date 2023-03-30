package main

import (
	"log"
	"os/exec"
	"time"

	"golang.org/x/sys/windows/registry"
)

func main() {
	// compile.Compile()
	url := "redactUrl"

	// Name := compile.Name
	Open(url, BrowserRegSearcherChrome(), BrowserRegSearcherYandex())

}

func Open(url string, regChrome, regYandex bool) error {
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

	if regYandex == true {
		err := exec.Command(cmd, argYa...).Start()
		if err != nil {
			log.Fatal(err, "regYandex")
			return err
		}
	}
	if regYandex == false && regChrome == true {
		err := exec.Command(cmd, argChr...).Start()
		time.Sleep(5 * time.Second)
		if err != nil {
			log.Fatal(err, "regChrome")
			return err
		}
	}

	if regYandex == false && regChrome == false {
		exec.Command(cmd, argEx...).Start()
		time.Sleep(5 * time.Second)
	}

	var err error
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
func BrowserRegSearcherChrome() bool {
	var regChrome bool
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Google`, registry.ALL_ACCESS)
	if err != nil {
		// log.Fatal(err, "Ошибка открытия ключа Chrome")
		regChrome = false
		return regChrome
	}

	defer k.Close()
	regFindingValueChrome, err := k.ReadSubKeyNames(0)
	if err != nil {
		log.Fatal(err, "Ошибка чтения значений строки реестра")
	}
	for _, v := range regFindingValueChrome {
		if v == "Chrome" {
			regChrome = true
		}
	}
	// regChrome = 0
	return regChrome
}
func BrowserRegSearcherYandex() bool {
	var regYandex bool
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Yandex`, registry.ALL_ACCESS)
	if err != nil {
		// 	log.Fatal(err, "Ошибка открытия ключа Yandex")
		regYandex = false
		return regYandex
	}

	defer k.Close()
	regFindingValueChrome, err := k.ReadSubKeyNames(0)
	if err != nil {
		log.Fatal(err, "Ошибка чтения значений строки реестра")
	}
	for _, v := range regFindingValueChrome {
		if v == "YandexBrowser" {
			regYandex = true
		}
	}
	// regYandex = 0
	return regYandex
}
