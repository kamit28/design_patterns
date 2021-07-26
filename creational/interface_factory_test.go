package creational

import (
	"reflect"
	"testing"
)

func TestInterfaceFactory(t *testing.T) {
	r := NewRectancgle(5.0, 3.0)

	if reflect.TypeOf(r) != reflect.TypeOf(&rectangle{}) {
		t.Errorf("expected type of rectangle, but got %T", r)
	}
	a := r.Area()
	if a != 15 {
		t.Errorf("expected area to be 15 but was %f", a)
	}

	p := r.Perimeter()
	if p != 16 {
		t.Errorf("expected perimeter o be 16 but was %f", p)
	}
}
