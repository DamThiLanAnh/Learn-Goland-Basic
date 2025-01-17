package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
		fmt.Println("end open file")
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil

}
func main() {
	openFile("file.txt")
}
