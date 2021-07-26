package creational

type pizzaMod func(*Pizza)

type PizzaBuilder struct {
	actions []pizzaMod
}

func (b *PizzaBuilder) Named(name string) *PizzaBuilder {
	b.actions = append(b.actions, func(p *Pizza) {
		p.name = name
	})
	return b
}

func (b *PizzaBuilder) OfSize(size PizzaSize) *PizzaBuilder {
	b.actions = append(b.actions, func(p *Pizza) {
		p.size = size
	})
	return b
}

func (b *PizzaBuilder) WithBase(pizzaBase PizzaBase) *PizzaBuilder {
	b.actions = append(b.actions, func(p *Pizza) {
		p.base = pizzaBase
	})
	return b
}

func (b *PizzaBuilder) WithSauce(sauce string) *PizzaBuilder {
	b.actions = append(b.actions, func(p *Pizza) {
		p.sauce = sauce
	})
	return b
}

func (b *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	b.actions = append(b.actions, func(p *Pizza) {
		p.toppings = append(p.toppings, topping)
	})
	return b
}

func (b *PizzaBuilder) Build() *Pizza {
	p := Pizza{}

	for _, action := range b.actions {
		action(&p)
	}
	return &p
}
