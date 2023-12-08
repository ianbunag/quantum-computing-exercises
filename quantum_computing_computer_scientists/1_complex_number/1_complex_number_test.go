package complex_number_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/quantum-computing-exercises/quantum_computing_computer_scientists/1_complex_number"
)

var _ = Describe("ComplexNumber", func() {
	Describe("Add", func() {
		It("should add complex numbers", func() {
			complexSum1 := New(1, 2).Add(New(3, 4))
			actualSum1 := (1 + 2i) + (3 + 4i)
			Expect(FormatComplexNumber(complexSum1)).To(Equal("4 + 6i"))
			Expect(FormatComplexNumber(complexSum1)).To(Equal(FormatComplex128(actualSum1)))

			complexSum2 := New(-8, 6).Add(New(5, -3))
			actualSum2 := (-8 + 6i) + (5 - 3i)
			Expect(FormatComplexNumber(complexSum2)).To(Equal("-3 + 3i"))
			Expect(FormatComplexNumber(complexSum2)).To(Equal(FormatComplex128(actualSum2)))

			complexSum3 := New(4, -7).Add(New(3, 2))
			actualSum3 := (4 - 7i) + (3 + 2i)
			Expect(FormatComplexNumber(complexSum3)).To(Equal("7 - 5i"))
			Expect(FormatComplexNumber(complexSum3)).To(Equal(FormatComplex128(actualSum3)))
		})
	})
})
