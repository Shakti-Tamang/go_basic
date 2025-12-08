package interfaceprac

type Shape interface {
	Area() float32

	Peremeter() float32
}

type Parameter struct {
	Length float32

	Breadth float32
}

func (p Parameter) Area() float32 {

	area := p.Length * p.Breadth
	return area
}

func (p Parameter) Peremeter() float32 {

	peri := 2 * (p.Length + p.Breadth)

	return peri
}
