package gof

import "fmt"

// Depencendy Inversion Principle
// High-Level Modules should not depend upon Low-Level Modules
// Both should depend on abstractions.

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low-Level Module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

func (r *Relationships) AddSibling(s1, s2 *Person) {
	r.relations = append(r.relations, Info{s1, Sibling, s2})
	r.relations = append(r.relations, Info{s2, Sibling, s1})
}

// High-Level Module
type Research struct {
	// This breaks DIP as it depends on low-level module
	relationships Relationships
}

func (r *Research) InvestigateParentToChild(name string) {
	relations := r.relationships.relations

	for _, relation := range relations {
		if relation.from.name == name &&
			relation.relationship == Parent {
			fmt.Printf("%s is a child of %s", relation.to.name, name)
		}
	}
}
