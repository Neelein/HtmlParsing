package main

import (
	"bufio"
	"fmt"
	"htmlParser/elements"
	"htmlParser/operators"
	"os"
	"strings"
)

func main() {

	var errorstring strings.Builder
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	fileName := strings.TrimSpace(userInput)

	styleAttrubute := map[string]string{"rel": "stylesheet", "type": "text/css", "href": "index.css"}
	html := elements.CreateElement("html", "", "", nil, "")
	link := elements.CreateElement("link", "", "", styleAttrubute, "")
	html.AddChildElement(&link)
	div := elements.CreateElement("div", "test", "background", nil, "test")
	div2 := elements.CreateElement("div", "test2", "", nil, "test1")
	div3 := elements.CreateElement("div", "test3", "", nil, "test1")
	p := elements.CreateElement("p", "test3", "", nil, "")
	err := div.AddChildElement(&div2)

	if err != nil {
		errorstring.WriteString(fmt.Sprintf("%s\n", err))
	}

	err = div.AddChildElement(&div3)

	if err != nil {
		errorstring.WriteString(fmt.Sprintf("%s\n", err))
	}

	err = div2.AddChildElement(&p)

	if err != nil {
		errorstring.WriteString(fmt.Sprintf("%s\n", err))
	}

	err = html.AddChildElement(&div)

	if err != nil {
		errorstring.WriteString(fmt.Sprintf("%s\n", err))
	}

	if errorstring.String() != "" {
		fmt.Println(errorstring.String())
		return
	}

	htmlContent := html.JsonString()
	operators.CreateHtmlFile(fileName, htmlContent)
	//element, err := html.GetElementById("test1")
}
