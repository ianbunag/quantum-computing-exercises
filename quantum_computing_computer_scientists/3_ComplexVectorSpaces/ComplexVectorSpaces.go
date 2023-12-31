package ComplexVectorSpaces

import "fmt"

type Scalar = complex128

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
	inversed := make(Vector, len(vector))

	for index, scalar := range vector {
		inversed[index] = scalar * -1
	}

	return inversed
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

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func (vector Vector) Conjugate() Vector {
	conjugated := make(Vector, len(vector))

	for index, scalar := range vector {
		conjugated[index] = complex(real(scalar), -imag(scalar))
	}

	return conjugated
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
	inversed := make(Matrix, len(matrix))

	for index := range inversed {
		inversed[index] = matrix[index].Inverse()
	}

	return inversed
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
	matrixColumns, matrixRows := len(matrix), len(matrix[0])
	multiplierColumns, multiplierRows := len(multiplier), len(multiplier[0])

	if matrixColumns != multiplierRows {
		panic(fmt.Sprintf("The number of columns in the first matrix (%d) must equal the number of rows in the second matrix (%d)", len(matrix), len(multiplier[0])))
	}

	product := make(Matrix, multiplierColumns)

	for columnIndex := 0; columnIndex < multiplierColumns; columnIndex += 1 {
		product[columnIndex] = make(Vector, matrixRows)

		for rowIndex := 0; rowIndex < matrixRows; rowIndex += 1 {
			for correspondingIndex := 0; correspondingIndex < matrixColumns; correspondingIndex += 1 {
				product[columnIndex][rowIndex] += matrix[correspondingIndex][rowIndex] * multiplier[columnIndex][correspondingIndex]
			}
		}
	}

	return product
}

// Average time complexity: O(n^2)
// Worst time complexity:   O(n^2)
// Space complexity:        O(n^2)
func (matrix Matrix) Conjugate() Matrix {
	conjugated := make(Matrix, len(matrix))

	for columnIndex := range conjugated {
		conjugated[columnIndex] = matrix[columnIndex].Conjugate()
	}

	return conjugated
}

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n^2)
func (matrix Matrix) Transpose() Matrix {
	transposed := make(Matrix, len(matrix))

	for columnIndex := range transposed {
		transposed[columnIndex] = make(Vector, len(matrix[columnIndex]))
		copy(transposed[columnIndex], matrix[columnIndex])
	}

	for columnIndex := range transposed {
		for rowIndex := len(transposed[columnIndex]) - 1; rowIndex > columnIndex; rowIndex-- {
			transposed[rowIndex][columnIndex], transposed[columnIndex][rowIndex] = transposed[columnIndex][rowIndex], transposed[rowIndex][columnIndex]
		}
	}

	return transposed
}

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n^2)
func (matrix Matrix) Adjoint() Matrix {
	adjointed := make(Matrix, len(matrix))

	for columnIndex := range adjointed {
		adjointed[columnIndex] = matrix[columnIndex].Conjugate()
	}

	for columnIndex := range adjointed {
		for rowIndex := len(adjointed[columnIndex]) - 1; rowIndex > columnIndex; rowIndex-- {
			adjointed[rowIndex][columnIndex], adjointed[columnIndex][rowIndex] = adjointed[columnIndex][rowIndex], adjointed[rowIndex][columnIndex]
		}
	}

	return adjointed
}

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
// Programming Drill 2.2.3 Write a function that accepts a vector and a matrix and outputs the vector resulting from the “action.”
func (matrix Matrix) Act(vector Vector) Vector {
	return matrix.Multiply(Matrix{vector})[0]
}
