package creational

import "testing"

func TestNewEmployee(t *testing.T) {
	emp, err := NewEmployee(Developer)
	if err != nil {
		t.Errorf("Error was not expected")
	}

	if emp.Role != Developer {
		t.Errorf("expected role was %d, but was %d", Developer, emp.Role)
	}

	if emp.Name != "" {
		t.Errorf("expected name to be empty")
	}

	emp.Name = "Amit"

	if emp.Name != "Amit" {
		t.Errorf("expected Name to be Amit, but was %s\n", emp.Name)
	}

	emp, err = NewEmployee(Manager)
	if err != nil {
		t.Errorf("Error was not expected")
	}

	if emp.Role != Manager {
		t.Errorf("expected role was %d, but was %d", Manager, emp.Role)
	}

	emp, err = NewEmployee(3)
	if err == nil {
		t.Errorf("Error was expected, but was nil")
	}

	if emp != nil {
		t.Errorf("expected nil but was not")
	}
}
