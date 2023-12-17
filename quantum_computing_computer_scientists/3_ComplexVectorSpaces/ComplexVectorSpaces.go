package ComplexVectorSpaces

type Scalar complex128

type Vector []Scalar

type Matrix []Vector

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
// Programming Drill 2.1.1 Write three functions that perform the addition, inverse, and scalar multiplication operations for Cn, i.e., write a function that accepts the appropriate input for each of the operations and outputs the vector.
func (vector Vector) Add(addend Vector) Vector {
	sum := make(Vector, len(vector))

	for index, scalar := range vector {
		sum[index] = scalar + addend[index]
	}

	return sum
}
