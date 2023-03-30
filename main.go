package main

import (
	"URLopener/compile"
	"URLopener/redact"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	compilator()
	// time.Sleep(100 * time.Second)
	time.Sleep(3 * time.Second)
	defer tryDefer()

}
func tryDefer() {

	redact.RedactSec()

	err := os.Rename("launcher.exe", compile.Name+".exe")
	if err != nil {
		log.Fatal(err)
	}
}
func compilator() {
	cmd := "cmd"
	compile.Compile()
	// Name := compile.Name
	redact.Redact()

	// resName = strings.
	// lol := "охуеть блин"
	// resName := fmt.Sprintf("%s", lol)
	// launcher := []string{"/c", "go build", "launcher/launcher.go", "&&", "ren", "launcher.exe", Name + ".exe"}
	launcher := []string{"/c", "go build", "launcher/launcher.go"}
	// renamer := []string{"/c", "rename", "launcher.exe", "fuck shit" + ".exe"}
	err := exec.Command(cmd, launcher...).Start()
	if err != nil {
		fmt.Println(err, "Ошибка создания файла")
	}

	// fmt.Println(resName)
}
