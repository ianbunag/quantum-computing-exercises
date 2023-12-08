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
func (complexNumber ComplexNumber) Add(addend ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      complexNumber.real + addend.real,
		imaginary: complexNumber.imaginary + addend.imaginary,
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
// Programming Drill 1.1.1 Write a program that accepts two complex numbers and outputs their sum and their product.
func (complexNumber ComplexNumber) Multiply(multiplier ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      (complexNumber.real * multiplier.real) - (complexNumber.imaginary * multiplier.imaginary),
		imaginary: (complexNumber.real * multiplier.imaginary) + (complexNumber.imaginary * multiplier.real),
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
// Programming Drill 1.2.1 Take the program that you wrote in the last programming drill and make it also perform subtraction and division of complex numbers. In addition, let the user enter a complex number and have the computer return its modulus and conjugate.
func (complexNumber ComplexNumber) Subtract(subtrahend ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      complexNumber.real - subtrahend.real,
		imaginary: complexNumber.imaginary - subtrahend.imaginary,
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
// Programming Drill 1.2.1 Take the program that you wrote in the last programming drill and make it also perform subtraction and division of complex numbers. In addition, let the user enter a complex number and have the computer return its modulus and conjugate.
func (complexNumber ComplexNumber) Divide(divisor ComplexNumber) ComplexNumber {
	numerator := complexNumber.Multiply(divisor.Conjugate())
	denominator := divisor.Multiply(divisor.Conjugate())

	return ComplexNumber{
		real:      numerator.real / denominator.real,
		imaginary: numerator.imaginary / denominator.real,
	}
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
// Programming Drill 1.2.1 Take the program that you wrote in the last programming drill and make it also perform subtraction and division of complex numbers. In addition, let the user enter a complex number and have the computer return its modulus and conjugate.
func (complexNumber ComplexNumber) Conjugate() ComplexNumber {
	return ComplexNumber{
		real:      complexNumber.real,
		imaginary: -complexNumber.imaginary,
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
