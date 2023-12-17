package ComplexNumber

import (
	"math"
)

type ComplexNumber struct {
	real      float64
	imaginary float64
}

func New(real, imaginary float64) ComplexNumber {
	return ComplexNumber{
		real:      real,
		imaginary: imaginary,
	}
}

// Programming Drill 1.3.1 Write a program that converts a complex number from its Cartesian representation to its polar representation and vice versa.
func FromPolar(magnitude, argument float64) ComplexNumber {
	return New(
		magnitude*math.Cos(argument),
		magnitude*math.Sin(argument),
	)
}

// Programming Drill 1.1.1 Write a program that accepts two complex numbers and outputs their sum and their product.
func (complexNumber ComplexNumber) Add(addend ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      complexNumber.real + addend.real,
		imaginary: complexNumber.imaginary + addend.imaginary,
	}
}

// Programming Drill 1.1.1 Write a program that accepts two complex numbers and outputs their sum and their product.
// Programming Drill 1.3.3 Expand your program. Add functions for multiplication, division, and returning the polar coordinates of a number.
func (complexNumber ComplexNumber) Multiply(multiplier ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      (complexNumber.real * multiplier.real) - (complexNumber.imaginary * multiplier.imaginary),
		imaginary: (complexNumber.real * multiplier.imaginary) + (complexNumber.imaginary * multiplier.real),
	}
}

// Programming Drill 1.2.1 Take the program that you wrote in the last programming drill and make it also perform subtraction and division of complex numbers. In addition, let the user enter a complex number and have the computer return its modulus and conjugate.
func (complexNumber ComplexNumber) Subtract(subtrahend ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      complexNumber.real - subtrahend.real,
		imaginary: complexNumber.imaginary - subtrahend.imaginary,
	}
}

// Programming Drill 1.2.1 Take the program that you wrote in the last programming drill and make it also perform subtraction and division of complex numbers. In addition, let the user enter a complex number and have the computer return its modulus and conjugate.
// Programming Drill 1.3.3 Expand your program. Add functions for multiplication, division, and returning the polar coordinates of a number.
func (complexNumber ComplexNumber) Divide(divisor ComplexNumber) ComplexNumber {
	numerator := complexNumber.Multiply(divisor.Conjugate())
	denominator := divisor.Multiply(divisor.Conjugate())

	return ComplexNumber{
		real:      numerator.real / denominator.real,
		imaginary: numerator.imaginary / denominator.real,
	}
}

// Programming Drill 1.2.1 Take the program that you wrote in the last programming drill and make it also perform subtraction and division of complex numbers. In addition, let the user enter a complex number and have the computer return its modulus and conjugate.
func (complexNumber ComplexNumber) Conjugate() ComplexNumber {
	return ComplexNumber{
		real:      complexNumber.real,
		imaginary: -complexNumber.imaginary,
	}
}

// Programming Drill 1.2.1 Take the program that you wrote in the last programming drill and make it also perform subtraction and division of complex numbers. In addition, let the user enter a complex number and have the computer return its modulus and conjugate.
func (complexNumber ComplexNumber) Modulus() float64 {
	return math.Sqrt(math.Pow(complexNumber.real, 2) + math.Pow(complexNumber.imaginary, 2))
}

// Programming Drill 1.3.1 Write a program that converts a complex number from its Cartesian representation to its polar representation and vice versa.
// Programming Drill 1.3.3 Expand your program. Add functions for multiplication, division, and returning the polar coordinates of a number.
func (complexNumber ComplexNumber) Polar() (float64, float64) {
	magnitude := complexNumber.Modulus()
	argument := math.Atan2(complexNumber.imaginary, complexNumber.real)

	return magnitude, argument
}

func (complexNumber ComplexNumber) Complex128() complex128 {
	return complex(complexNumber.real, complexNumber.imaginary)
}
