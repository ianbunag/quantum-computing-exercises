package complex_number

import (
	"strconv"
)

type ComplexNumber struct {
	real      float64
	imaginary float64
}

func NewComplexNumber(real, imaginary float64) ComplexNumber {
	return ComplexNumber{
		real:      real,
		imaginary: imaginary,
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
// Programming Drill 1.1.1 Write a program that accepts two complex numbers and outputs their sum and their product.
func (to ComplexNumber) Add(from ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      to.real + from.real,
		imaginary: to.imaginary + from.imaginary,
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
// Programming Drill 1.1.1 Write a program that accepts two complex numbers and outputs their sum and their product.
func (to ComplexNumber) Multiply(from ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      (to.real * from.real) - (to.imaginary * from.imaginary),
		imaginary: (to.real * from.imaginary) + (to.imaginary * from.real),
	}
}

func (complexNumber ComplexNumber) Complex128() complex128 {
	return complex(complexNumber.real, complexNumber.imaginary)
}

type ComplexFormatter struct {
	fmt     byte
	prec    int
	bitSize int
}

func NewComplexFormatter(fmt byte, prec int) ComplexFormatter {
	return ComplexFormatter{
		fmt:     fmt,
		prec:    prec,
		bitSize: 128,
	}
}

func (formatter ComplexFormatter) Format(complexNumber complex128) string {
	return strconv.FormatComplex(
		complexNumber,
		formatter.fmt,
		formatter.prec,
		formatter.bitSize,
	)
}
