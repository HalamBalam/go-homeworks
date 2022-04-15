package hw

import (
	"errors"
	"math"
)

// Point Добавлен новый тип, соответствующий сущности "точка" в системе координат.
// Тип Geom удален, т.к. ему не соответсвует ни одна из общепринятых сущностей.
type Point struct {
	X, Y float64
}

// CalculateDistance В функцию передаются 2 точки, между которыми нужно посчитать расстояние.
// Вместо вывода ошибки в лог, возвращается объект error.
func CalculateDistance(p1 Point, p2 Point) (distance float64, err error) {

	if p1.correct() && p2.correct() {
		distance = math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
	} else {
		err = errors.New("координаты не могут быть меньше нуля")
	}

	// возврат расстояния между точками
	return distance, err
}

// По условиям задачи, координаты не могут быть меньше 0.
func (p *Point) correct() bool {
	return p.X >= 0 && p.Y >= 0
}
