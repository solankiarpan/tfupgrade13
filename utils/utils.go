package utils

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func ReplaceUsingRegexp(regex string, replace string, filename string) {

	re := regexp.MustCompile(regex)
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		lines[i] = re.ReplaceAllString(line, replace)
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}

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
		if strings.Contains(line, `}"`) {
			lines[i] = strings.Replace(line, `}"`, "", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `“$`) {
			continue
		} else {
			lines[i] = strings.Replace(line, `“`, "", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `“$`) {
			continue
		} else {
			lines[i] = strings.Replace(line, `”`, "", -1)
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func ReplaceLineInDir(dir string, regex string, replace string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		if file.IsDir() {
			ReplaceLineInDir(dir+"/"+file.Name(), regex, replace)
		} else {
			ReplaceUsingRegexp(regex, replace, dir+"/"+file.Name())
		}
	}
}

func ReplaceLineInDirOnly(dir string, regex string, replace string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			ReplaceUsingRegexp(regex, replace, dir+"/"+file.Name())
		}
	}
}
