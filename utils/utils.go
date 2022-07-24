package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReplaceLine(filename string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, `"${`) {
			lines[i] = strings.Replace(line, `"${`, "", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `}`) {
			lines[i] = strings.Replace(line, `}"`, "", -1)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func ReplaceLineInDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		if file.IsDir() {
			ReplaceLineInDir(dir + "/" + file.Name())
		} else {
			ReplaceLine(dir + "/" + file.Name())
		}
	}
}

func ReplaceLineInDirOnly(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			ReplaceLine(dir + "/" + file.Name())
		}
	}
}
