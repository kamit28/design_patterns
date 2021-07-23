package main

func main() {
	r := Research{}

	parent := Person{"Binoy"}
	child1 := Person{"Anshu"}
	child2 := Person{"Amit"}
	child3 := Person{"Sumit"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)
	relationships.AddParentAndChild(&parent, &child3)

	r.InvestigateParentToChild(parent.name)
}
