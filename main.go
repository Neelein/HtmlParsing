package main

import (
	"htmlParser/elements"
)

func main() {

	const fileName = "index.html"
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
	htmlContent := html.JsonString()
	createHtmlFile(fileName, htmlContent)
}
