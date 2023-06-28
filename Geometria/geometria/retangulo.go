package geometria

type Retangulo struct {
	Geo    Geometria
	Base   float32
	Altura float32
}

func (r Retangulo) Area() float32 {
	return r.Base * r.Altura
}

func (r Retangulo) Perimetro() float32 {
	return r.Base + r.Altura
}
