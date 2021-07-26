package creational

import "testing"

func TestPrototypeFactory(t *testing.T) {
	c := NewLiberalCandiadate("Amit", "Toongabbie")
	if c.Name != "Amit" {
		t.Errorf("expected Amit, but got %s\n", c.Name)
	}

	if c.Constituency.Name != "Toongabbie" {
		t.Errorf("expected Toongabbie, but got %s\n", c.Constituency.Name)
	}

	c = NewLabourCandiadate("Sumeet", "Armidale")
	if c.Name != "Sumeet" {
		t.Errorf("expected Sumeet, but got %s\n", c.Name)
	}

	if c.Constituency.Name != "Armidale" {
		t.Errorf("expected Armidale, but got %s\n", c.Constituency.Name)
	}

	c = NewGreensCandiadate("Anshu", "Gladstone")
	if c.Name != "Anshu" {
		t.Errorf("expected Anshu, but got %s\n", c.Name)
	}

	if c.Constituency.Name != "Gladstone" {
		t.Errorf("expected Gladstone, but got %s\n", c.Constituency.Name)
	}
}
