package complex_number

import "fmt"

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

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func (to ComplexNumber) Add(from ComplexNumber) ComplexNumber {
	return ComplexNumber{
		real:      to.real + from.real,
		imaginary: to.imaginary + from.imaginary,
	}
}

func (complex ComplexNumber) String() string {
	return FormatComplexNumber(complex)
}

func simpleFormat(real, imaginary float64) string {
	intReal, intImaginary := int(real), int(imaginary)

	if intReal == 0 {
		return fmt.Sprintf("%di", intImaginary)
	}

	if intImaginary == 0 {
		return fmt.Sprintf("%d", intReal)
	}

	if intImaginary < 0 {
		return fmt.Sprintf("%d - %di", intReal, intImaginary*-1)
	}

	return fmt.Sprintf("%d + %di", intReal, intImaginary)
}

// Format ComplexNumber in the simple form of a + bi.
func FormatComplexNumber(complex ComplexNumber) string {
	return simpleFormat(complex.real, complex.imaginary)
}

// Format complex128 in the simple form of a + bi.
func FormatComplex128(complex complex128) string {
	return simpleFormat(real(complex), imag(complex))
}
