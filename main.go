package main

import "fmt"

func main() {
  lambda.Start(handleRequest)
}

func handleRequest(event Event) {
  fmt.Println("sample")
}
