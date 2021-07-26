package creational

type Contractor struct {
	Name, Position string
	DayRate        int
}

// functional way
func NewContractorFactory(position string, dayRate int) func(name string) *Contractor {
	return func(name string) *Contractor {
		return &Contractor{name, position, dayRate}
	}
}

// Traditional way
type Contractorfactory struct {
	Position string
	DayRate  int
}

func (f *Contractorfactory) Create(name string) *Contractor {
	return &Contractor{name, f.Position, f.DayRate}
}

func GetContractorFactory(position string, dayRate int) *Contractorfactory {
	return &Contractorfactory{position, dayRate}
}
