package fib

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFibonacci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fibonacci Suite")
}

var _ = Describe("Fibonacci Tests", func() {

	Context("Fibonacci", func() {

		It("calculate(0) should equal 0", func() {
			Expect(calculate(0)).Should(Equal(0))
		})

		It("calculate(5) should equal 5", func() {
			Expect(calculate(5)).Should(Equal(5))
		})

		It("calculate(20) should equal 6765", func() {
			Expect(calculate(20)).Should(Equal(6765))
		})

		It("calculate(25) should equal 75025", func() {
			Expect(calculate(25)).Should(Equal(75025))
		})
	})
})
