package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dialy", func() {
  Context("top()", func() {
    It("returns a 200 Status Code", func() {
      Request("GET", "/", top)
      Expect(response.Code).To(Equal(200))
      Expect(response.Body).To(ContainSubstring("Daily"))
    })
  })
})

