package creational

import (
	"reflect"
	"testing"
)

func TestNewPersonBuilder(t *testing.T) {
	b := NewPersonBuilder()

	if reflect.TypeOf(b) != reflect.TypeOf(&PersonBuilder{}) {
		t.Errorf("expected *PersonBuilder but was %T\n", b)
	}

	pab := b.LivesAt()
	if reflect.TypeOf(pab) != reflect.TypeOf(&PersonAddressBuilder{}) {
		t.Errorf("expected *PersonAddressBuilder but was %T\n", pab)
	}

	p := pab.
		Street("27 Lamonerie St").
		City("Toongabbie").
		State("New South Wales").
		PostCode("2146").Build()
	if reflect.TypeOf(p) != reflect.TypeOf(&Person{}) {
		t.Errorf("expected *Person but was %T\n", p)
	}

	pjb := b.WorksAt()
	if reflect.TypeOf(pjb) != reflect.TypeOf(&PersonJobBuilder{}) {
		t.Errorf("expected *PersonJobBuilder but was %T\n", pjb)
	}

	p = pjb.
		Company("ANZ Banking Corp").
		Department("Digital Banking").
		Position("Senior Software engineer").
		AnnualIncome(178000.00).Build()
	if reflect.TypeOf(p) != reflect.TypeOf(&Person{}) {
		t.Errorf("expected *Person but was %T\n", p)
	}

	if p.Position != "Senior Software engineer" {
		t.Errorf("expected Senior Software engineer, but was %s\n", p.Position)
	}
}
