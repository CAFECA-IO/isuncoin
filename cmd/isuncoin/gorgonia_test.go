package main

import (
	"fmt"

	"gorgonia.org/tensor"
)

func main() {
	a := tensor.New(tensor.Expects(tensor.Float64), tensor.WithShape(2, 2))
	fmt.Println(a)
}
