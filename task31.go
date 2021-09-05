package main

import (
	"fmt"
	"math"
)

type Point struct { // объявляем структурный тип Point, представляющий ТОЧКУ
	x, y int // ТОЧКА описывается 2 параметрами (координатами на плоскости, в случае если рассматриваем плоскость)
}

func newPoint(x, y int) Point { // конструктор - возвращает инстанс типа Point
	return Point{x, y}
}
func (p *Point) showDistanceTo(anotherPoint Point) float64 {
	return math.Sqrt(math.Pow(float64(anotherPoint.x-p.x), 2) + math.Pow(float64(anotherPoint.y-p.y), 2))
	// расстояние равно квадратному корню из суммы квадратов разностей координат по каждой оси
}

func main() {
	p1 := newPoint(4, 8)
	p2 := newPoint(2, 6)
	fmt.Println(p1.showDistanceTo(p2), p2.showDistanceTo(p1)) // Выведем расстояние от 1-ой до 2-ой и от 2-ой до 1-ой
}
