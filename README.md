# Quantum Computing Exercises
Collection of Quantum computing programming exercises.
From the book Quantum Computing for Computer Scientists (Noson S. Yanofsky, Mirco A. Mannucci).

## Install test executable
```sh
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
```

## Bootstrap exercise
```sh
./scaffold.sh quantum_computing_computer_scientists complex_number
```

## Run tests
```sh
# Run all tests
ginkgo ./...

# Run specific tests matching description
ginkgo --focus "should add complex number" ./...

# Run specific tests matching filename
ginkgo --focus-file quantum_computing_computer_scientists/complex_number/ComplexNumber_test.go ./...
```

## Time and space complexities

### Quick guide
| Complexity | Name        | Description                                         | Example                 |
|------------|-------------|-----------------------------------------------------|-------------------------|
| O(1)       | Constant    | Metric is the same regardless of input              | Indexed access          |
| O(log n)   | Logarithmic | Metric is lower with each additional step           | Binary trees            |
| O(n)       | Linear      | Metric depends linearly on the input                | For loop                |
| O(n^2)     | Quadratic   | Metric grows quadratically for each additional step | Nested for loop         |
| O(2^n)     | Exponential | Metric grows exponentially for each additional step | Fibonacci calculation   |
| O(n!)      | Factorial   | Metric grows factorially for each additional step   | Permutation calculation |

### Notes
- Constant complexities may be disregarded, for example:
  - O(2) is considered as O(1)
  - O(2n) is considered as O(n)
  - O(nk) is considered as O(n)
- Space complexity should include everything created in the lifecycle of the
  algorithm:
  - Auxiliary variables (space used while the algorithm is being executed)
  - Output

### Cheatsheet
[Cheatsheet](https://www.bigocheatsheet.com/)
