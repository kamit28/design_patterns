package creational

type Rectangle interface {
	Area() float32
	Perimeter() float32
}

type rectangle struct {
	length, width float32
}

func (r *rectangle) Area() float32 {
	return r.length * r.width
}

func (r *rectangle) Perimeter() float32 {
	return 2 * (r.length + r.width)
}

func NewRectancgle(l, w float32) Rectangle {
	return &rectangle{l, w}
}
