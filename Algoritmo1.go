//Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como parâmetro e
//calcule a área do círculo (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi.

package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{
		radius: 10,
	}
	fmt.Println(area(circle))
}

type Circle struct {
	radius float64
}

func area(circle Circle) float64 {
	calculate := math.Pi * (circle.radius * circle.radius)
	return calculate
}
