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

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) GetName() string {
	return p.name
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
	//relationships Relationships

	// abstraction to depend upon
	browser RelationshipBrowser
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for _, relation := range r.relations {
		if relation.from.name == name &&
			relation.relationship == Parent {
			result = append(result, relation.to)
		}
	}
	return result
}

func NewResearch(relationships *Relationships) Research {
	return Research{relationships}
}

func (r *Research) InvestigateParentToChild(name string) {
	results := r.browser.FindAllChildrenOf(name)

	for _, result := range results {
		fmt.Printf("%s has a child named %s\n", name, result.name)
	}
}
