package elements

import (
	"fmt"
	"strings"
)

type Element struct {
	elementType  string
	childElement []*Element
	innerText    string
	class        string
	attribute    map[string]string
}

func CreateElement(elementType string, innerText string, class string, attribute map[string]string) Element {
	element := Element{
		elementType:  elementType,
		innerText:    innerText,
		class:        class,
		childElement: nil,
		attribute:    attribute,
	}

	return element
}

func (e *Element) AddChildElement(element *Element) {
	(*e).childElement = append((*e).childElement, element)
}

func (e *Element) DisPlayFile() {
	result := parsingToHtml(e)
	fmt.Print(result)
}

func (e *Element) JsonString() string {
	return parsingToHtml(e)
}

func parsingToHtml(element *Element) string {
	var stringBuilder strings.Builder
	for key, valu := range (*element).attribute {
		stringBuilder.WriteString(fmt.Sprintf("%s='%s'", key, valu))
	}

	attributeString := stringBuilder.String()

	elementstring := fmt.Sprintf("<%s class='%s' %s>%s", (*element).elementType, (*element).class, attributeString, (*element).innerText)

	if (*element).childElement == nil {
		elementstring += fmt.Sprintf("</%s>", (*element).elementType)
		return elementstring
	} else {
		for _, element := range (*element).childElement {
			elementstring += parsingToHtml(element)
		}

	}
	elementstring += fmt.Sprintf("</%s>", (*element).elementType)
	return elementstring
}
