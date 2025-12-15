//go:build !darwin

package main

func calculateNPUFlops() float64 {
	// Not implemented for non-Darwin systems
	return 0
}
