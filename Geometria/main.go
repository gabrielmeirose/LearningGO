package main

import (
	"fmt"
	geo "geo/geometria"
)

func showData(x float32, y float32) string {
	return fmt.Sprintf("%f | %f", x, y)
}

func main() {
	fmt.Println("Starting Program...")

	r := geo.Retangulo{
		Geo:    geo.Geometria{Cor: "Vermelho"},
		Base:   20,
		Altura: 10,
	}

	c := geo.Circulo{
		Geo:  geo.Geometria{Cor: "Vermelho"},
		Raio: 10,
	}

	t := geo.Triangulo{
		Geo:    geo.Geometria{Cor: "Vermelho"},
		Base:   20,
		Altura: 10,
		Lado_x: 15,
		Lado_y: 12,
	}

	// Calling methods
	fmt.Println("Rectangle")
	fmt.Println(showData(r.Area(), r.Perimetro()))
	fmt.Println("Circle")
	fmt.Println(showData(c.Area(), c.Perimetro()))
	fmt.Println("Triangle")
	fmt.Println(showData(t.Area(), t.Perimetro()))

}
