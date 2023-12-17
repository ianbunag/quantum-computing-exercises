package ComplexFormatter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ComplexFormatter "github.com/yvnbunag/quantum-computing-exercises/quantum_computing_computer_scientists/2_ComplexFormatter"
)

var _ = Describe("2ComplexFormatter", func() {
	It("should format complex numbers", func() {
		intFormatter := ComplexFormatter.New('f', 0)
		Expect(intFormatter.Format(1 + 2i)).To(Equal("(1+2i)"))

		floatFormatter := ComplexFormatter.New('f', 2)
		Expect(floatFormatter.Format(1 + 2i)).To(Equal("(1.00+2.00i)"))
	})
})
