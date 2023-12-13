package complex_number_test

import (
	"math"
	"math/cmplx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/quantum-computing-exercises/quantum_computing_computer_scientists/1_complex_number"
)

var _ = Describe("ComplexNumber", func() {
	It("should add complex numbers", func() {
		complexFormatter := NewComplexFormatter('f', 0)

		sum1 := NewComplexNumber(1, 2).Add(NewComplexNumber(3, 4))
		actualSum1 := (1 + 2i) + (3 + 4i)
		Expect(complexFormatter.Format(sum1.Complex128())).To(Equal("(4+6i)"))
		Expect(sum1.Complex128()).To(Equal(actualSum1))

		sum2 := NewComplexNumber(-8, 6).Add(NewComplexNumber(5, -3))
		actualSum2 := (-8 + 6i) + (5 - 3i)
		Expect(complexFormatter.Format(sum2.Complex128())).To(Equal("(-3+3i)"))
		Expect(sum2.Complex128()).To(Equal(actualSum2))

		sum3 := NewComplexNumber(4, -7).Add(NewComplexNumber(3, 2))
		actualSum3 := (4 - 7i) + (3 + 2i)
		Expect(complexFormatter.Format(sum3.Complex128())).To(Equal("(7-5i)"))
		Expect(sum3.Complex128()).To(Equal(actualSum3))
	})

	It("should multiply complex numbers", func() {
		complexFormatter := NewComplexFormatter('f', 0)

		product1 := NewComplexNumber(2, 3).Multiply(NewComplexNumber(4, -5))
		actualProduct1 := (2 + 3i) * (4 - 5i)
		Expect(complexFormatter.Format(product1.Complex128())).To(Equal("(23+2i)"))
		Expect(product1.Complex128()).To(Equal(actualProduct1))

		product2 := NewComplexNumber(1, 2).Multiply(NewComplexNumber(3, -4))
		actualProduct2 := (1 + 2i) * (3 - 4i)
		Expect(complexFormatter.Format(product2.Complex128())).To(Equal("(11+2i)"))
		Expect(product2.Complex128()).To(Equal(actualProduct2))

		product3 := NewComplexNumber(-2, 1).Multiply(NewComplexNumber(3, -2))
		actualProduct3 := (-2 + 1i) * (3 - 2i)
		Expect(complexFormatter.Format(product3.Complex128())).To(Equal("(-4+7i)"))
		Expect(product3.Complex128()).To(Equal(actualProduct3))
	})

	It("should subtract complex numbers", func() {
		complexFormatter := NewComplexFormatter('f', 0)

		difference1 := NewComplexNumber(5, 2).Subtract(NewComplexNumber(3, 4))
		actualDifference1 := (5 + 2i) - (3 + 4i)
		Expect(complexFormatter.Format(difference1.Complex128())).To(Equal("(2-2i)"))
		Expect(difference1.Complex128()).To(Equal(actualDifference1))

		difference2 := NewComplexNumber(-2, -1).Subtract(NewComplexNumber(3, 2))
		actualDifference2 := (-2 - 1i) - (3 + 2i)
		Expect(complexFormatter.Format(difference2.Complex128())).To(Equal("(-5-3i)"))
		Expect(difference2.Complex128()).To(Equal(actualDifference2))

		difference3 := NewComplexNumber(1, 2).Subtract(NewComplexNumber(4, -5))
		actualDifference3 := (1 + 2i) - (4 - 5i)
		Expect(complexFormatter.Format(difference3.Complex128())).To(Equal("(-3+7i)"))
		Expect(difference3.Complex128()).To(Equal(actualDifference3))
	})

	It("should divide complex numbers", func() {
		complexFormatter := NewComplexFormatter('f', 2)

		quotient1 := NewComplexNumber(2, 3).Divide(NewComplexNumber(1, -2))
		actualQuotient1 := (2 + 3i) / (1 - 2i)
		Expect(complexFormatter.Format(quotient1.Complex128())).To(Equal("(-0.80+1.40i)"))
		Expect(quotient1.Complex128()).To(Equal(actualQuotient1))

		quotient2 := NewComplexNumber(4, -1).Divide(NewComplexNumber(2, 1))
		actualQuotient2 := (4 - 1i) / (2 + 1i)
		Expect(complexFormatter.Format(quotient2.Complex128())).To(Equal("(1.40-1.20i)"))
		Expect(quotient2.Complex128()).To(Equal(actualQuotient2))

		quotient3 := NewComplexNumber(3, 2).Divide(NewComplexNumber(1, 1))
		actualQuotient3 := (3 + 2i) / (1 + 1i)
		Expect(complexFormatter.Format(quotient3.Complex128())).To(Equal("(2.50-0.50i)"))
		Expect(quotient3.Complex128()).To(Equal(actualQuotient3))
	})

	It("should get conjugate of complex numbers", func() {
		complexFormatter := NewComplexFormatter('f', 0)

		conjugate1 := NewComplexNumber(3, 4).Conjugate()
		actualConjugate := cmplx.Conj(3 + 4i)
		Expect(complexFormatter.Format(conjugate1.Complex128())).To(Equal("(3-4i)"))
		Expect(conjugate1.Complex128()).To(Equal(actualConjugate))

		conjugate2 := NewComplexNumber(-2, -1).Conjugate()
		actualConjugate2 := cmplx.Conj(-2 - 1i)
		Expect(complexFormatter.Format(conjugate2.Complex128())).To(Equal("(-2+1i)"))
		Expect(conjugate2.Complex128()).To(Equal(actualConjugate2))

		conjugate3 := NewComplexNumber(1, -2).Conjugate()
		actualConjugate3 := cmplx.Conj(1 - 2i)
		Expect(complexFormatter.Format(conjugate3.Complex128())).To(Equal("(1+2i)"))
		Expect(conjugate3.Complex128()).To(Equal(actualConjugate3))
	})

	It("should get modulus of complex numbers", func() {
		modulus1 := NewComplexNumber(3, 4).Modulus()
		actualModulus1 := cmplx.Abs(3 + 4i)
		Expect(modulus1).To(Equal(float64(5)))
		Expect(modulus1).To(Equal(actualModulus1))

		modulus2 := NewComplexNumber(-2, -1).Modulus()
		actualModulus2 := cmplx.Abs(-2 - 1i)
		Expect(modulus2).To(Equal(math.Sqrt(5)))
		Expect(modulus2).To(Equal(actualModulus2))

		modulus3 := NewComplexNumber(1, -2).Modulus()
		actualModulus3 := cmplx.Abs(1 - 2i)
		Expect(modulus3).To(Equal(math.Sqrt(5)))
		Expect(modulus3).To(Equal(actualModulus3))
	})

	It("should get polar representation of complex numbers", func() {
		magnitude1, argument1 := NewComplexNumber(1, 1).Polar()
		actualMagnitude1, actualArgument1 := cmplx.Polar(1 + 1i)
		Expect(magnitude1).To(Equal(actualMagnitude1))
		Expect(argument1).To(Equal(actualArgument1))

		magnitude2, argument2 := NewComplexNumber(-2, -2).Polar()
		actualMagnitude2, actualArgument2 := cmplx.Polar(-2 - 2i)
		Expect(magnitude2).To(Equal(actualMagnitude2))
		Expect(argument2).To(Equal(actualArgument2))

		magnitude3, argument3 := NewComplexNumber(4, 3).Polar()
		actualMagnitude3, actualArgument3 := cmplx.Polar(4 + 3i)
		Expect(magnitude3).To(Equal(actualMagnitude3))
		Expect(argument3).To(Equal(actualArgument3))
	})

	It("should get complex number from polar representation", func() {
		complexNumber1 := NewComplexNumberFromPolar(3, math.Pi/4)
		actualComplexNumber1 := cmplx.Rect(3, math.Pi/4)
		Expect(complexNumber1.Complex128()).To(Equal(actualComplexNumber1))

		complexNumber2 := NewComplexNumberFromPolar(2, -(math.Pi / 3))
		actualComplexNumber2 := cmplx.Rect(2, -(math.Pi / 3))
		Expect(complexNumber2.Complex128()).To(Equal(actualComplexNumber2))

		complexNumber3 := NewComplexNumberFromPolar(1, math.Pi)
		actualComplexNumber3 := cmplx.Rect(1, math.Pi)
		Expect(complexNumber3.Complex128()).To(Equal(actualComplexNumber3))
	})
})
