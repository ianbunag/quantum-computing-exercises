package ComplexVectorSpaces_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ComplexVectorSpaces "github.com/yvnbunag/quantum-computing-exercises/quantum_computing_computer_scientists/3_ComplexVectorSpaces"
)

var _ = Describe("3ComplexVectorSpaces", func() {
	It("should add vectors", func() {
		vector1 := ComplexVectorSpaces.Vector{1 + 2i, -3 + 4i}
		vector2 := ComplexVectorSpaces.Vector{5 + 6i, 7 + -8i}
		sum := vector1.Add(vector2)

		Expect(sum).To(Equal(ComplexVectorSpaces.Vector{6 + 8i, 4 + -4i}))
	})

	It("should invert vectors", func() {
		vector := ComplexVectorSpaces.Vector{1 + 2i, -3 + 4i}
		inverse := vector.Inverse()

		Expect(inverse).To(Equal(ComplexVectorSpaces.Vector{-1 + -2i, 3 + -4i}))
		Expect(vector.Add(inverse)).To(Equal(ComplexVectorSpaces.Vector{0 + 0i, 0 + 0i}))
	})
})
