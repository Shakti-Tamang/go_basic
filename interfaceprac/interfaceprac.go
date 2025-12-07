package interfaceprac

type shape interface {
	Area(a float32, b float32) float32

	Peremeter(a float32, b float32) float32
}

func Area(a float32, b float32) float32 {

	area := a * b
	return area
}

func Peremeter(a float32, b float32) float32 {

	peri := 2 * (a + b)

	return peri
}
