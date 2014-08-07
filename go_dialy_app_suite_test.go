//package go_dialy_app_test
package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dialy", func() {
  Context("Show Top", func() {
    It("returns a 200 Status Code", func() {
      Request("GET", "/", top)
      Expect(recorder.Code).To(Equal(200))
      //Expect(recorder.Body).To(ContainSubstring("Dialya"))
    })
  })
})

