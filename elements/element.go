package elements

import (
	"errors"
	"fmt"
	"strings"
)

type Element struct {
	elementType  string
	childElement []*Element
	innerText    string
	class        string
	attribute    map[string]string
	id           string
}

func CreateElement(elementType string, innerText string, class string, attribute map[string]string, id string) Element {
	element := Element{
		elementType:  elementType,
		innerText:    innerText,
		class:        class,
		childElement: nil,
		attribute:    attribute,
		id:           id,
	}

	return element
}

func (e *Element) AddChildElement(element *Element) error {
	if checkIdExist((*element).id, e) != nil {
		errorString := fmt.Sprintf("Id %s is exist!!", (*element).id)
		return errors.New(errorString)
	}
	(*e).childElement = append((*e).childElement, element)
	return nil
}

func (e *Element) DisPlayFile() {
	result := parsingToHtml(e)
	fmt.Print(result)
}

func (e *Element) JsonString() string {
	return parsingToHtml(e)
}

func (e *Element) RemoveChildElement(index int) {
	(*e).childElement = append((*e).childElement[:index], (*e).childElement[index+1:]...)
}

func parsingToHtml(element *Element) string {

	var stringBuilder strings.Builder
	elementstring := ""
	calssString := ""
	for key, valu := range (*element).attribute {
		stringBuilder.WriteString(fmt.Sprintf("%s='%s'", key, valu))
	}

	if (*element).class != "" {
		calssString = fmt.Sprintf("class='%s'", (*element).class)
	}

	attributeString := stringBuilder.String()

	if (*element).elementType == "link" {
		elementstring = fmt.Sprintf("<%s %s %s/>", (*element).elementType, calssString, attributeString)
	} else {
		elementstring = fmt.Sprintf("<%s %s %s>%s", (*element).elementType, calssString, attributeString, (*element).innerText)
	}

	if (*element).childElement == nil {
		if (*element).elementType != "link" {
			elementstring += fmt.Sprintf("</%s>", (*element).elementType)
		}
		return elementstring
	} else {
		for _, element := range (*element).childElement {
			elementstring += parsingToHtml(element)
		}

	}
	elementstring += fmt.Sprintf("</%s>", (*element).elementType)
	return elementstring
}

func (e *Element) GetElementById(id string) (*Element, error) {
	element := findElementById(id, e)
	if element != nil {
		return element, nil
	} else {
		return nil, errors.New("Cant find the element")
	}
}

func findElementById(id string, element *Element) *Element {
	var elements []*Element
	for _, ele := range (*element).childElement {
		if (*ele).id == id {
			return ele
		}
		if len(ele.childElement) > 0 {
			elements = append(elements, ele)
		}
	}

	if len(elements) > 0 {
		for _, ele := range elements {
			if findElementById(id, ele) != nil {
				return findElementById(id, ele)
			}
		}
	}
	return nil
}

func checkIdExist(id string, element *Element) *Element {
	if id == "" {
		return nil
	}

	if findElementById(id, element) != nil {
		return element
	}

	return nil
}
