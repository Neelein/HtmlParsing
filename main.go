package main

import (
	"htmlParser/elements"
)

func main() {
	const fileName = "index.html"
	div := elements.CreateElement("div", "test", "")
	div2 := elements.CreateElement("div", "test2", "")
	div3 := elements.CreateElement("div", "test3", "")
	p := elements.CreateElement("p", "test3", "")
	div.AddChildElement(&div2)
	div.AddChildElement(&div3)
	div2.AddChildElement(&p)
	htmlContent := div.JsonString()
	createHtmlFile(fileName, htmlContent)
}
