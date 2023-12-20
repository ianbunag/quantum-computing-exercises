package ComplexVectorSpaces

type Scalar complex128

// Column of scalars
type Vector []Scalar

// Columns of scalars
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

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
// Programming Drill 2.1.1 Write three functions that perform the addition, inverse, and scalar multiplication operations for Cn, i.e., write a function that accepts the appropriate input for each of the operations and outputs the vector.
func (vector Vector) Inverse() Vector {
	inverse := make(Vector, len(vector))

	for index, scalar := range vector {
		inverse[index] = scalar * -1
	}

	return inverse
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
// Programming Drill 2.1.1 Write three functions that perform the addition, inverse, and scalar multiplication operations for Cn, i.e., write a function that accepts the appropriate input for each of the operations and outputs the vector.
func (vector Vector) ScalarMultiply(multiplier Scalar) Vector {
	product := make(Vector, len(vector))

	for index, scalar := range vector {
		product[index] = scalar * multiplier
	}

	return product
}

// Average time complexity: O(n^2)
// Worst time complexity:   O(n^2)
// Space complexity:        O(n^2)
// Programming Drill 2.2.1 Convert your functions from the last programming drill so that instead of accepting elements of Cn, they accept elements of Cm×n.
func (matrix Matrix) Add(addend Matrix) Matrix {
	sum := make(Matrix, len(matrix))

	for index := range sum {
		sum[index] = matrix[index].Add(addend[index])
	}

	return sum
}

// Average time complexity: O(n^2)
// Worst time complexity:   O(n^2)
// Space complexity:        O(n^2)
// Programming Drill 2.2.1 Convert your functions from the last programming drill so that instead of accepting elements of Cn, they accept elements of Cm×n.
func (matrix Matrix) Inverse() Matrix {
	sum := make(Matrix, len(matrix))

	for index := range sum {
		sum[index] = matrix[index].Inverse()
	}

	return sum
}

// Average time complexity: O(n^2)
// Worst time complexity:   O(n^2)
// Space complexity:        O(n^2)
// Programming Drill 2.2.1 Convert your functions from the last programming drill so that instead of accepting elements of Cn, they accept elements of Cm×n.
func (matrix Matrix) ScalarMultiply(multiplier Scalar) Matrix {
	product := make(Matrix, len(matrix))

	for index := range product {
		product[index] = matrix[index].ScalarMultiply(multiplier)
	}

	return product
}

// Average time complexity: O(n^3)
// Worst time complexity:   O(n^3)
// Space complexity:        O(n^2)
// Programming Drill 2.2.2 Write a function that accepts two complex matrices of the appropriate size. The function should do matrix multiplication and return the result.
func (matrix Matrix) Multiply(multiplier Matrix) Matrix {
	product := make(Matrix, len(matrix))

	for columnIndex := range product {
		product[columnIndex] = make(Vector, len(matrix[columnIndex]))

		for rowIndex := range product[columnIndex] {
			for correspondingIndex := 0; correspondingIndex < len(matrix); correspondingIndex++ {
				product[columnIndex][rowIndex] += multiplier[columnIndex][correspondingIndex] * matrix[correspondingIndex][rowIndex]
			}
		}
	}

	return product
}

// Average time complexity: O(n^3)
// Worst time complexity:   O(n^3)
// Space complexity:        O(n^2)
func (matrix Matrix) Transpose() Matrix {
	transposed := make(Matrix, len(matrix))

	for columnIndex := range transposed {
		transposed[columnIndex] = make(Vector, len(matrix[columnIndex]))
		copy(transposed[columnIndex], matrix[columnIndex])
	}

	for columnIndex := range transposed {
		for rowIndex := range transposed[columnIndex] {
			if rowIndex > columnIndex {
				transposed[rowIndex][columnIndex], transposed[columnIndex][rowIndex] = transposed[columnIndex][rowIndex], transposed[rowIndex][columnIndex]
			}
		}
	}

	return transposed
}
