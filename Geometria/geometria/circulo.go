package geometria

import "math"

type Circulo struct {
	Geo  Geometria
	Raio float32
}

func (c Circulo) Area() float32 {
	return float32(math.Pi * math.Pow(float64(c.Raio), 2.0))
}

func (c Circulo) Perimetro() float32 {
	return math.Pi * c.Raio * 2
}
