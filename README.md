# Quantum Computing Exercises
Collection of Quantum computing programming exercises.
- Quantum Computing for Computer Scientists (Noson S. Yanofsky, Mirco A. Mannucci)

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
