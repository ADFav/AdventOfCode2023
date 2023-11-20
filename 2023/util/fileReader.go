package util

import (
	"bufio"
	"fmt"
	"os"
)

func FileReader(dayNum string, year string) []string {
	fileName := year + "/input/" + dayNum + ".txt"
	readFile, err := os.Open(fileName)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var result []string
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	return result
}
