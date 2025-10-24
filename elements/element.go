package elements

import "fmt"

type Element struct {
	elementType  string
	childElement []*Element
	innerText    string
	class        string
}

func CreateElement(elementType string, innerText string, class string) Element {
	element := Element{
		elementType:  elementType,
		innerText:    innerText,
		class:        class,
		childElement: nil,
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
	elementstring := fmt.Sprintf("<%s>%s", (*element).elementType, (*element).innerText)

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
