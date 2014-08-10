package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

  "github.com/go-martini/martini"
  "net/http"
  "net/http/httptest"
	"testing"
)

var (
  response *httptest.ResponseRecorder
)

func TestGoDailyApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoDailyApp Suite")
}

func Request(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get(route, handler)
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}

