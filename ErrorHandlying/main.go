package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	text, err := readFile("test.txtr")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(text)
	}
}

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		//diff error to print with errors.New
		//return "", errors.New("unable to open the file ")

		//diff error to print with fmt.Error
		return "", fmt.Errorf("unable to open the file %s", fileName)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.New("unable to read the file")
	}
	return string(data), nil
}
