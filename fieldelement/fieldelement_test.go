package fieldelement

import (
	"testing"
)

func TestNewFieldElement(t *testing.T) {
	tests := []struct {
		num   int64
		prime int64
		panic bool
	}{
		{num: 2, prime: 7, panic: false},
		{num: 0, prime: 7, panic: false},
		{num: 6, prime: 7, panic: false},
		{num: 7, prime: 7, panic: true},
		{num: -1, prime: 7, panic: true},
	}

	for _, tt := range tests {
		if tt.panic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic for num %d and prime %d, but did not panic", tt.num, tt.prime)
				}
			}()
		}

		fe := NewFieldElement(tt.num, tt.prime)
		if !tt.panic && (fe.num != tt.num || fe.prime != tt.prime) {
			t.Errorf("expected FieldElement with num %d and prime %d, got num %d and prime %d", tt.num, tt.prime, fe.num, fe.prime)
		}
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		num      int64
		prime    int64
		exponent int64
		expected int64
	}{
		{num: 3, prime: 13, exponent: 3, expected: 1},
		{num: 7, prime: 13, exponent: -3, expected: 8},
	}

	for _, tt := range tests {
		fe := NewFieldElement(tt.num, tt.prime)
		result := fe.Pow(tt.exponent)
		if result.num != tt.expected {
			t.Errorf("Pow(%d, %d) with prime %d: expected %d, got %d", tt.num, tt.exponent, tt.prime, tt.expected, result.num)
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		num         int64
		prime       int64
		denominator int64
		expected    int64
	}{
		{num: 2, prime: 19, denominator: 7, expected: 3},
		{num: 7, prime: 19, denominator: 5, expected: 9},
	}

	for _, tt := range tests {
		numerator := NewFieldElement(tt.num, tt.prime)
		denominator := NewFieldElement(tt.denominator, tt.prime)

		result := numerator.Div(denominator)
		if result.num != tt.expected {
			t.Errorf("Div(%d, %d) with prime %d: expected %d, got %d", tt.num, tt.denominator, tt.prime, tt.expected, result.num)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		num1     int64
		num2     int64
		prime    int64
		expected int64
	}{
		{num1: 2, num2: 3, prime: 19, expected: 6},
		{num1: 7, num2: 3, prime: 19, expected: 2},
		{num1: 10, num2: 12, prime: 19, expected: 6},
	}

	for _, tt := range tests {
		fe1 := NewFieldElement(tt.num1, tt.prime)
		fe2 := NewFieldElement(tt.num2, tt.prime)
		result := fe1.Mul(fe2)
		if result.num != tt.expected {
			t.Errorf("Mul(%d, %d) with prime %d: expected %d, got %d", tt.num1, tt.num2, tt.prime, tt.expected, result.num)
		}
	}
}

func TestMulThreeElements(t *testing.T) {
	tests := []struct {
		num1     int64
		num2     int64
		num3     int64
		prime    int64
		expected int64
	}{
		{num1: 2, num2: 3, num3: 4, prime: 19, expected: 5},
		{num1: 7, num2: 3, num3: 5, prime: 19, expected: 10},
		{num1: 95, num2: 45, num3: 31, prime: 97, expected: 23},
	}

	for _, tt := range tests {
		fe1 := NewFieldElement(tt.num1, tt.prime)
		fe2 := NewFieldElement(tt.num2, tt.prime)
		fe3 := NewFieldElement(tt.num3, tt.prime)
		result := fe1.Mul(fe2).Mul(fe3)
		if result.num != tt.expected {
			t.Errorf("Mul(%d, %d, %d) with prime %d: expected %d, got %d", tt.num1, tt.num2, tt.num3, tt.prime, tt.expected, result.num)
		}
	}
}
