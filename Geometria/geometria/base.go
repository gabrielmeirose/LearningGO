package geometria

type Geometria struct {
	Cor string
}

type InterfaceGeo interface {
	Area() float32
	Perimetro() float32
}
