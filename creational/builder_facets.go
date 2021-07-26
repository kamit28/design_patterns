package creational

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) LivesAt() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) WorksAt() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonAddressBuilder) Street(street string) *PersonAddressBuilder {
	b.person.StreetAddress = street
	return b
}

func (b *PersonAddressBuilder) City(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) State(state string) *PersonAddressBuilder {
	b.person.State = state
	return b
}

func (b *PersonAddressBuilder) PostCode(postCode string) *PersonAddressBuilder {
	b.person.PostCode = postCode
	return b
}

func (pjb *PersonJobBuilder) Company(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) Department(department string) *PersonJobBuilder {
	pjb.person.Department = department
	return pjb
}

func (pjb *PersonJobBuilder) Position(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) AnnualIncome(annualIncome float32) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}
