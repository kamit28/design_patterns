package creational

import (
	"reflect"
	"testing"
)

func TestNewContractorFactory(t *testing.T) {
	f := NewContractorFactory("Senior Java Developer", 850)
	c := f("Amit")

	if reflect.TypeOf(f).Kind() != reflect.Func {
		t.Errorf("expected func but got %T\n", f)
	}

	if reflect.TypeOf(c) != reflect.TypeOf(&Contractor{}) {
		t.Errorf("expected *Contractor but got %T\n", c)
	}

	if c.DayRate != 850 {
		t.Errorf("expected 850, but got %d\n", c.DayRate)
	}

	if c.Name != "Amit" {
		t.Errorf("expected Amit, but got %s\n", c.Name)
	}

	if c.Position != "Senior Java Developer" {
		t.Errorf("expected Senior Java Developer, but got %s\n", c.Position)
	}
}

func TestGetContractorFactory(t *testing.T) {
	f := GetContractorFactory("Developer", 900)

	if reflect.TypeOf(f) != reflect.TypeOf(&Contractorfactory{}) {
		t.Errorf("expected *ContractorFactory but got %T\n", f)
	}

	if f.DayRate != 900 {
		t.Errorf("expected 900, but got %d\n", f.DayRate)
	}

	if f.Position != "Developer" {
		t.Errorf("expected Developer, but got %s\n", f.Position)
	}
}

func TestCreate(t *testing.T) {
	f := GetContractorFactory("Developer", 900)
	c := f.Create("Amit")

	if reflect.TypeOf(c) != reflect.TypeOf(&Contractor{}) {
		t.Errorf("expected *Contractor but got %T\n", c)
	}

	if c.Name != "Amit" {
		t.Errorf("expected Amit, but got %s\n", c.Name)
	}
}
