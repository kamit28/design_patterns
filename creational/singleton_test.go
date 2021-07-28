package creational

import (
	"reflect"
	"testing"
)

func TestSingleton(t *testing.T) {
	db := GetDatabase()

	db2 := GetDatabase()
	// Testing singleton behaviour
	if db != db2 {
		t.Error("expected both to point to same instance, but they are pointing to different instances")
	}

	// Testing functionality
	if reflect.TypeOf(db) != reflect.TypeOf(&database{}) {
		t.Errorf("expected type was *database, but was %T\n", db)
	}

	p := db.GetPopulation("Mumbai")
	if p != 20185064 {
		t.Errorf("expected population was 20185064, but got %d\n", p)
	}
}
