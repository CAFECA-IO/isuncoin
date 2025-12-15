package main

import (
	"fmt"
	"testing"

	"gorgonia.org/tensor"
)

func TestGorgoniaMisc(t *testing.T) {
	a := tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(2, 2))
	fmt.Println(a)
}
