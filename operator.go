package main

import (
	"fmt"
	"os"
)

func createHtmlFile(fileName string, fileContent string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error create file")
		return
	}

	defer file.Close()

	_, err = file.WriteString(fileContent)

	if err != nil {
		fmt.Println("Error writing file content")
		return
	}

	fmt.Println("create file")
}
