package fib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fibonacci Hidden Tests", func() {

	Context("Fibonacci", func() {

		It("calculate(14) should equal 377", func() {
			Expect(calculate(14)).Should(Equal(377))
		})

		It("calculate(15) should equal 610", func() {
			Expect(calculate(15)).Should(Equal(610))
		})

		It("calculate(18) should equal 2584", func() {
			Expect(calculate(18)).Should(Equal(2584))
		})

		It("calculate(19) should equal 4181", func() {
			Expect(calculate(19)).Should(Equal(4181))
		})
	})
})
