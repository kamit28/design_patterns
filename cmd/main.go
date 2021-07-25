package main

import (
	"fmt"

	"github.com/kamit28/design_patterns/creational"
	"github.com/kamit28/design_patterns/gof"
)

func main() {
	TestDependencyInversion()
	TestBuilder()
	TestBuilderFacets()
	TestFunctionalBuilder()
	TestFactoryMethod()
}

func TestDependencyInversion() {
	parent := gof.Person{}
	parent.SetName("Vasudev")
	child1 := gof.Person{}
	child1.SetName("Balram")
	child2 := gof.Person{}
	child2.SetName("Krishna")
	child3 := gof.Person{}
	child3.SetName("Subhadra")

	relationships := gof.Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)
	relationships.AddParentAndChild(&parent, &child3)
	r := gof.NewResearch(&relationships)
	r.InvestigateParentToChild("Vasudev")
}

func TestBuilder() {
	b := creational.NewHtmlBuilder("ul")

	b.AddChild("li", "hello").AddChild("li", "World")
	fmt.Println(b.String())
}

func TestBuilderFacets() {
	b := creational.NewPersonBuilder()

	person := b.LivesAt().
		Street("27 Lamonerie St").
		City("Toongabbie").
		State("New South Wales").
		PostCode("2146").
		WorksAt().
		Company("ANZ Banking Corp").
		Department("Digital Banking").
		Position("Senior Software engineer").
		AnnualIncome(178000.00).Build()

	fmt.Println(person)
}

func TestFunctionalBuilder() {
	b := creational.PizzaBuilder{}

	pizza := b.Named("Meat Lovers").
		OfSize(creational.Medium).
		WithBase(creational.DeepDish).
		WithSauce("Basil Pesto").
		AddTopping("Pepperoni").
		AddTopping("Chicken mince").
		AddTopping("Pulled pork").
		AddTopping("Jalapinos").
		AddTopping("Chopped Olives").
		AddTopping("Red Peppers").
		Build()

	fmt.Println(pizza)
}

func TestFactoryMethod() {
	user, err := creational.NewUser("Amit", 41, "amit@myemail.com", "myPassword!baby1")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
