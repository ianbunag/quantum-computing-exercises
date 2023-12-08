package complex_number_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/quantum-computing-exercises/quantum_computing_computer_scientists/1_complex_number"
)

var _ = Describe("ComplexNumber", func() {
	Describe("Add", func() {
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
	})
})
