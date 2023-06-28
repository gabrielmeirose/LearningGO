package geometria

import "math"

type Triangulo struct {
	Geo    Geometria
	Base   float32
	Altura float32
	Lado_x float32
	Lado_y float32
}

func (t Triangulo) Area() float32 {
	return 0.5 * (t.Base * t.Lado_x) * float32(math.Sin(float64(t.Lado_y)))
}

func (t Triangulo) Perimetro() float32 {
	return t.Base + t.Lado_x + t.Lado_y
}
