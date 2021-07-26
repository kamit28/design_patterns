package creational

import (
	"bytes"
	"encoding/gob"
)

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.Street,
		a.City,
		a.PostCode,
	}
}

func (o *Order) DeepCopy() *Order {
	p := *o
	p.ShippingAddress = o.ShippingAddress.DeepCopy()
	copy(p.Items, o.Items)
	return &p
}

func (o *Order) DeepCopySerialization() *Order {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(o)

	d := gob.NewDecoder(&b)
	repeatOrder := Order{}
	_ = d.Decode(&repeatOrder)
	return &repeatOrder
}
