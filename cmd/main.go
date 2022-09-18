package main

import (
	"fmt"
	"net/http"

	"github.com/tkdailey11/gator/pkg/functions"
	"github.com/tkdailey11/gator/pkg/server"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("")
}

func main() {
  var az functions.AzureFunctionApp
  az.HttpFunctions = append(az.HttpFunctions, helloHandler)
  message, err := server.Hello("World")
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Println(message)
  }
}