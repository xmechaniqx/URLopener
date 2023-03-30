package redact

import (
	"URLopener/compile"
	"bufio"
	"os"
	"strings"
)

func SearchAndReplace(filename, search, replace string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), search, replace, -1)
		lines = append(lines, line)
	}
	return ReplaceText(filename, lines)
}

func ReplaceText(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 644)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return err
}

func Redact() {
	filename := "C:/Users/a57.rantonenko/go/src/URLopener/launcher/launcher.go"
	search := "redactUrl"
	replace := compile.Url
	SearchAndReplace(filename, search, replace)
	// fmt.Print("done1")
}

func RedactSec() {
	filename := "C:/Users/a57.rantonenko/go/src/URLopener/launcher/launcher.go"
	search := compile.Url
	replace := "redactUrl"
	SearchAndReplace(filename, search, replace)
	// fmt.Print("done2")
}
