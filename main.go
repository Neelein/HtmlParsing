package main

import (
	"bufio"
	"htmlParser/elements"
	"htmlParser/operators"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	fileName := strings.TrimSpace(userInput)

	styleAttrubute := map[string]string{"rel": "stylesheet", "type": "text/css", "href": "index.css"}
	html := elements.CreateElement("html", "", "", nil)
	link := elements.CreateElement("link", "", "", styleAttrubute)
	html.AddChildElement(&link)
	div := elements.CreateElement("div", "test", "background", nil)
	div2 := elements.CreateElement("div", "test2", "", nil)
	div3 := elements.CreateElement("div", "test3", "", nil)
	p := elements.CreateElement("p", "test3", "", nil)
	div.AddChildElement(&div2)
	div.AddChildElement(&div3)
	div2.AddChildElement(&p)
	html.AddChildElement(&div)
	div.RemoveChildElement(0)
	htmlContent := html.JsonString()
	operators.CreateHtmlFile(fileName, htmlContent)
}
