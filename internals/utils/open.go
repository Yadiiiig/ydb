package utils

import "os"

func OpenData(f string) *os.File {
	file, err := os.OpenFile(f, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}
