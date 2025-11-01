package operators

import (
	"fmt"
	"os"
)

func CreateHtmlFile(fileName string, fileContent string) {
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

func reRender(fileName string, fileContent string) {
	os.Truncate(fileName, 0)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("File open fail")
		return
	}

	defer file.Close()

	_, err = file.WriteString(fileContent)

	if err != nil {
		fmt.Println("writing fail!!")
		return
	}

	fmt.Println("Render")
}
