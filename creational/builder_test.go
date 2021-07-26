package creational

import (
	"reflect"
	"testing"
)

func TestNewHtmlBuilder(t *testing.T) {
	b := NewHtmlBuilder("html")

	if reflect.TypeOf(b) != reflect.TypeOf(&HtmlBuilder{}) {
		t.Errorf("expect type was *HtmlBuilder, but was %T", b)
	}

	if b.RootName != "html" {
		t.Errorf("expected root element name was html, but was %s", b.RootName)
	}

	if len(b.Root.Elements) != 0 {
		t.Errorf("expect child elements size was 0, but was %d", len(b.Root.Elements))
	}
}

func TestString(t *testing.T) {
	b := NewHtmlBuilder("ul")

	b.AddChild("li", "hello")
	var expected = `<ul>
  <li>
    hello
  </li>
</ul>
`
	var result = b.String()
	if result != expected {
		t.Errorf("expected output was\n %s, but was\n %s\n", expected, result)
	}
}

func TestStringHtmlElement(t *testing.T) {
	b := HtmlElement{"tag", "My Tag", []HtmlElement{}}
	b.Elements = append(b.Elements, HtmlElement{"sub", "Sub Tag", []HtmlElement{}})

	expected := `<tag>
  My Tag
  <sub>
    Sub Tag
  </sub>
</tag>
`
	result := b.String()
	if result != expected {
		t.Errorf("expected output was\n %s, but was\n %s\n", expected, result)
	}
}

func TestAddChild(t *testing.T) {
	b := NewHtmlBuilder("root")

	if b.RootName != "root" {
		t.Errorf("expected root element name was root, but was %s", b.RootName)
	}

	if len(b.Root.Elements) != 0 {
		t.Errorf("expect child elements size was 0, but was %d", len(b.Root.Elements))
	}

	b.AddChild("child1", "First Child")
	if len(b.Root.Elements) != 1 {
		t.Errorf("expect child elements size was 1, but was %d", len(b.Root.Elements))
	}

	if b.Root.Elements[0].Name != "child1" {
		t.Errorf("expected child1, but was %s\n", b.Root.Elements[0].Name)
	}

	if b.Root.Elements[0].Text != "First Child" {
		t.Errorf("expected First Child, but was %s\n", b.Root.Elements[0].Text)
	}
}
