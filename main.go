package main

import (
	"fmt"

	gen "github.com/anilkusc/golang-random-string-generator/generator"
)

func main() {
	fmt.Println(gen.Generate(10))

}
