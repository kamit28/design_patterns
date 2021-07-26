package creational

import (
	"reflect"
	"testing"
)

func TestFunctionalBuilder(t *testing.T) {
	b := PizzaBuilder{}

	pizza := b.Named("Meat Lovers").
		OfSize(Medium).
		WithBase(DeepDish).
		WithSauce("Basil Pesto").
		AddTopping("Pepperoni").
		AddTopping("Chicken mince").
		AddTopping("Pulled pork").
		AddTopping("Jalapinos").
		AddTopping("Chopped Olives").
		AddTopping("Red Peppers").
		Build()

	if reflect.TypeOf(pizza) != reflect.TypeOf(&Pizza{}) {
		t.Errorf("expected type was %T, but was %T", reflect.TypeOf(&Pizza{}), reflect.TypeOf(pizza))
	}

	if pizza.base != DeepDish {
		t.Errorf("expected pizza base was %v but was %v", DeepDish, pizza.base)
	}

	if pizza.size != Medium {
		t.Errorf("expected pizza size was %v but was %v", Medium, pizza.size)
	}

	if pizza.sauce != "Basil Pesto" {
		t.Errorf("expected pizza base was Basil Pesto but was %v", pizza.sauce)
	}

	if len(pizza.toppings) != 6 {
		t.Errorf("expected 5 toppings, but got %d", len(pizza.toppings))
	}
}
