package ellipticcurve

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y, A, B *float64
}

func NewPoint(x, y, a, b *float64) (*Point, error) {
	p := Point{X: x, Y: y, A: a, B: b}

	// Accommodates points at infinity. nil indicates infinity.
	if x == nil && y == nil {
		return nil, nil
	}

	if math.Pow(*y, 2) != math.Pow(*x, 3)+((*a)*(*x))+(*b) {
		return nil, fmt.Errorf("(%v, %v) is not on the curve", x, y)
	}

	return &p, nil
}

func (p *Point) areInSameCurve(other *Point) bool {
	return *p.A == *other.A && *p.B == *other.B
}

func (p *Point) Equal(other *Point) bool {
	return *p.X == *other.X && *p.Y == *other.Y && p.areInSameCurve(other)
}

func (p *Point) NotEqual(other *Point) bool {
	return !p.Equal(other)
}

// // Add adds two points on the elliptic curve
// func (p *Point) Add(other *Point) (*Point, error) {
// 	if !p.areInSameCurve(other) {
// 		return nil, fmt.Errorf("points %v, %v are not on the same curve", p, other)
// 	}

// 	if p.X == nil {
// 		return other, nil
// 	}

// 	if other.X == nil {
// 		return p, nil
// 	}

// 	// Points are additive inverses
// 	if *p.X == *other.X && *p.Y != *other.Y {
// 		return NewPoint(nil, nil, p.A, p.B)
// 	}

// 	if *p.X != *other.X {
// 		s := (*other.Y - *p.Y) / (*other.X - *p.X)
// 		x := s*s - *p.X - *other.X
// 		y := s*(*p.X-x) - *p.Y
// 		return NewPoint(&x, &y, p.A, p.B)
// 	}

// 	if p.Equal(other) {
// 		// y coordinate == 0
// 		if *p.Y == 0 {
// 			return NewPoint(nil, nil, p.A, p.B)
// 		}

// 		s := (3**p.X**2 + *p.A) / (2 * *p.Y)
// 		x := s*s - 2**p.X
// 		y := s*(*p.X-x) - *p.Y
// 		return NewPoint(&x, &y, p.A, p.B)
// 	}

// 	return nil, fmt.Errorf("unhandled case in point addition")
// }

// // Mul multiplies a point by a coefficient
// func (p *Point) Mul(coefficient int) (*Point, error) {
// 	coef := coefficient
// 	current := p
// 	result, _ := NewPoint(nil, nil, p.A, p.B)
// 	for coef != 0 {
// 		if coef&1 != 0 {
// 			var err error
// 			result, err = result.Add(current)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 		var err error
// 		current, err = current.Add(current)
// 		if err != nil {
// 			return nil, err
// 		}
// 		coef >>= 1
// 	}
// 	return result, nil
// }
