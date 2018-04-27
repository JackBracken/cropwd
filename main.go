package main

import (
	"bytes"
	"fmt"
	"os"
)

func split(path string) []string {
	var r []string
	var buffer bytes.Buffer

	for i, char := range path {
		if i == 0 {
			continue
		}
		if char == os.PathSeparator || i == len(path)-1 {
			r = append(r, buffer.String())
			buffer.Reset()
		} else {
			buffer.WriteRune(char)
		}
	}
	return r
}

func prefix(str string, len int) string {
	return str[0:len]
}

func cropwd() string {
	var buffer bytes.Buffer
	wd, _ := os.Getwd()
	dirs := split(wd)

	buffer.WriteRune(os.PathSeparator)
	for i, dir := range dirs {
		if i == len(dirs)-1 {
			buffer.WriteString(dir)
		} else {
			buffer.WriteString(prefix(dir, 1))
			buffer.WriteRune(os.PathSeparator)
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println(cropwd())
}
