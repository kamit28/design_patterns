package creational

import (
	"fmt"
	"strings"
)

const (
	indentSize int = 2
)

type HtmlElement struct {
	Name, Text string
	Elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.Name))

	if len(e.Text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.Text)
		sb.WriteString("\n")

	}

	for _, el := range e.Elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.Name))

	return sb.String()
}

type HtmlBuilder struct {
	RootName string
	Root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		RootName: rootName,
		Root:     HtmlElement{rootName, "", []HtmlElement{}},
	}
}

func (b *HtmlBuilder) String() string {
	return b.Root.String()
}

func (b *HtmlBuilder) AddChild(childName string, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}

	b.Root.Elements = append(b.Root.Elements, e)
	return b
}
