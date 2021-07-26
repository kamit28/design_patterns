package creational

// Contractor
type Contractor struct {
	Name, Position string
	DayRate        int
}

//Person
type Person struct {
	StreetAddress, City, State, PostCode string

	CompanyName, Position, Department string
	AnnualIncome                      float32
}

// User
type User struct {
	name     string
	age      int
	email    string
	password uint32
	status   bool // will be defaulted to false, ans will be changed to true once email is verified
}

// pizza type
type PizzaSize int

const (
	Small PizzaSize = iota
	Medium
	Large
)

type PizzaBase int

const (
	Thin PizzaBase = iota
	Traditional
	DeepDish
)

type Pizza struct {
	name     string
	size     PizzaSize
	base     PizzaBase
	sauce    string
	toppings []string
}

// OrderItem type
type OrderItem struct {
	ItemName string
	Price    float32
	Quantity int
}

// Address type
type Address struct {
	Street, City, PostCode string
}

// Order type
type Order struct {
	Name            string
	OrderTotal      float32
	Items           []OrderItem
	ShippingAddress *Address
}

//Employee type
type RoleType int

const (
	Developer RoleType = iota
	Manager
)

type Employee struct {
	Name         string
	Role         RoleType
	AnnualIncome int
}

// Candidate
type Candidate struct {
	Name         string
	Party        PartyName
	Constituency Constituency
}

type PartyName int

const (
	Labour PartyName = iota
	Liberal
	Greens
)

type Constituency struct {
	Name  string
	State string
}
