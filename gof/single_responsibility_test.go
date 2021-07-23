package gof

import (
	"testing"
)

func TestAddEntry(t *testing.T) {
	j := Journal{}
	entryCount = 0
	lenEntries := j.AddEntry("Hello, world")
	newLen := j.AddEntry("How are you")
	if newLen != lenEntries+1 {
		t.Errorf("length of entries should have been %d, but was %d\n", lenEntries+1, newLen)
	}
	expectedResult := "2: How are you"
	actualVal := j.entries[1]
	if actualVal != expectedResult {
		t.Errorf("expected result was %s but was %s", expectedResult, actualVal)
	}
}

func TestRemoveEntry(t *testing.T) {
	j := Journal{}
	entryCount = 0
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	lenEntries := len(j.entries)
	newLen, _ := j.RemoveEntry(3)
	if newLen != lenEntries-1 {
		t.Errorf("length of entries should have been %d, but was %d\n", lenEntries-1, newLen)
	}
}

func TestRemoveEntry_InvalidIndex(t *testing.T) {
	j := Journal{}
	entryCount = 0
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	j.AddEntry("Hello")
	lenEntries := len(j.entries)
	newLen, err := j.RemoveEntry(7)
	if err == nil {
		t.Errorf("test for RemoveEntry with index 7 should have returned error index out of range, but did not")
	}
	if newLen != lenEntries {
		t.Errorf("length of entries should have remained %d, but was changed to %d", lenEntries, newLen)
	}
}
