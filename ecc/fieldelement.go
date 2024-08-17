package ecc

import (
	"fmt"
	"math"
)

type FieldElement struct {
	num   int64
	prime int64
}

func NewFieldElement(num, prime int64) *FieldElement {
	if num >= prime || num < 0 {
		panic(fmt.Sprintf("num %d not in field range 0 to %d", num, prime-1))
	}
	return &FieldElement{num, prime}
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", f.prime, f.num)
}

func (f *FieldElement) Equals(other *FieldElement) bool {
	return f.prime == other.prime && f.num == other.num
}

func (f *FieldElement) Add(other *FieldElement) *FieldElement {
	if f.prime != other.prime {
		panic("cannot add two numbers in different Fields")
	}
	num := (f.num + other.num) % f.prime
	return NewFieldElement(num, f.prime)
}

func (f *FieldElement) Sub(other *FieldElement) *FieldElement {
	if f.prime != other.prime {
		panic("cannot subtract two numbers in different Fields")
	}
	num := (f.num - other.num) % f.prime
	return NewFieldElement(num, f.prime)
}

func (f *FieldElement) Mul(other *FieldElement) *FieldElement {
	if f.prime != other.prime {
		panic("cannot multiply two numbers in different Fields")
	}
	num := (f.num * other.num) % f.prime
	return NewFieldElement(num, f.prime)
}

func (f *FieldElement) Pow(exponent int64) *FieldElement {
	n := exponent
	for n < 0 {
		n += f.prime - 1
	}
	num := int64(math.Pow(float64(f.num), float64(n))) % f.prime
	return NewFieldElement(num, f.prime)
}

func (f *FieldElement) Div(other *FieldElement) *FieldElement {
	if f.prime != other.prime {
		panic("cannot divide two numbers in different Fields")
	}
	num := (f.num * other.Pow(f.prime-2).num) % f.prime
	return NewFieldElement(num, f.prime)
}
