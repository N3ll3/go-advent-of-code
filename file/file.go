package file

import (
	"bufio"
	"fmt"
	"os"
)



func Readfile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("opening file error", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines

}
