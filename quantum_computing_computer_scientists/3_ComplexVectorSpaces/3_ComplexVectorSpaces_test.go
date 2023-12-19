package ComplexVectorSpaces_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ComplexVectorSpaces "github.com/yvnbunag/quantum-computing-exercises/quantum_computing_computer_scientists/3_ComplexVectorSpaces"
)

var _ = Describe("3ComplexVectorSpaces", func() {
	It("should add vectors", func() {
		vector1 := ComplexVectorSpaces.Vector{1 + 2i, -3 + 4i}
		vector2 := ComplexVectorSpaces.Vector{5 + 6i, 7 - 8i}
		sum := vector1.Add(vector2)

		Expect(sum).To(Equal(ComplexVectorSpaces.Vector{6 + 8i, 4 - 4i}))
	})

	It("should invert vectors", func() {
		vector := ComplexVectorSpaces.Vector{1 + 2i, -3 + 4i}
		inverse := vector.Inverse()

		Expect(inverse).To(Equal(ComplexVectorSpaces.Vector{-1 - 2i, 3 - 4i}))
		Expect(vector.Add(inverse)).To(Equal(ComplexVectorSpaces.Vector{0 + 0i, 0 + 0i}))
	})

	It("should multiply vectors by scalars", func() {
		vector := ComplexVectorSpaces.Vector{2 + 3i, 4 - 2i}

		productFromReal := vector.ScalarMultiply(2)
		Expect(productFromReal).To(Equal(ComplexVectorSpaces.Vector{4 + 6i, 8 - 4i}))

		productFromComplex := vector.ScalarMultiply(1 - 1i)
		Expect(productFromComplex).To(Equal(ComplexVectorSpaces.Vector{5 + 1i, 2 - 6i}))
	})

	It("should add matrices", func() {
		matrix1 := ComplexVectorSpaces.Matrix{
			{1 + 2i, 3 + 4i},
			{5 + 6i, 7 + 8i},
		}
		matrix2 := ComplexVectorSpaces.Matrix{
			{9 + 10i, 11 + 12i},
			{13 + 14i, 15 + 16i},
		}
		sum := matrix1.Add(matrix2)

		Expect(sum).To(Equal(ComplexVectorSpaces.Matrix{
			{10 + 12i, 14 + 16i},
			{18 + 20i, 22 + 24i},
		}))
	})

	It("should invert matrices", func() {
		matrix := ComplexVectorSpaces.Matrix{
			{1 + 2i, 3 + 4i},
			{5 + 6i, 7 + 8i},
		}
		inverse := matrix.Inverse()

		Expect(inverse).To(Equal(ComplexVectorSpaces.Matrix{
			{-1 - 2i, -3 - 4i},
			{-5 - 6i, -7 - 8i},
		}))
		Expect(matrix.Add(inverse)).To(Equal(ComplexVectorSpaces.Matrix{
			{0 + 0i, 0 + 0i},
			{0 + 0i, 0 + 0i},
		}))
	})

	It("should multiply matrices by scalars", func() {
		matrix := ComplexVectorSpaces.Matrix{
			{1 + 2i, 3 + 4i},
			{5 + 6i, 7 + 8i},
		}

		productFromReal := matrix.ScalarMultiply(2)
		Expect(productFromReal).To(Equal(ComplexVectorSpaces.Matrix{
			{2 + 4i, 6 + 8i},
			{10 + 12i, 14 + 16i},
		}))

		productFromComplex := matrix.ScalarMultiply(1 - 1i)
		Expect(productFromComplex).To(Equal(ComplexVectorSpaces.Matrix{
			{3 + 1i, 7 + 1i},
			{11 + 1i, 15 + 1i},
		}))
	})
})
