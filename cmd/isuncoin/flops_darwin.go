//go:build darwin

package main

/*
#cgo LDFLAGS: -framework Accelerate
#include <Accelerate/Accelerate.h>
*/
import "C"

import (
	"math/rand"
	"time"
)

func calculateNPUFlops() float64 {
	// Use macOS Accelerate Framework (BLAS) via CGO
	// This often leverages the AMX (Apple Matrix Coprocessor) on M-series chips
	N := C.int(2048) // Larger matrix for Accelerate
	alpha := C.double(1.0)
	beta := C.double(0.0)

	// Allocate memory in C-compatible way (or just pass slices, CGO handles pinning usually, but large arrays are cleaner managed)
	// For simplicity, we create slices and pass ptrs.
	size := int(N * N)
	a := make([]float64, size)
	b := make([]float64, size)
	c := make([]float64, size)

	// Initialize
	for i := 0; i < size; i++ {
		a[i] = rand.Float64()
		b[i] = rand.Float64()
	}

	start := time.Now()

	// cblas_dgemm(Order, TransA, TransB, M, N, K, alpha, A, lda, B, ldb, beta, C, ldc)
	// C = alpha * A * B + beta * C
	// RowMajor = 101, NoTrans = 111
	C.cblas_dgemm(101, 111, 111, N, N, N, alpha,
		(*C.double)(&a[0]), N,
		(*C.double)(&b[0]), N,
		beta,
		(*C.double)(&c[0]), N)

	elapsed := time.Since(start).Seconds()

	ops := 2.0 * float64(N) * float64(N) * float64(N)
	return ops / elapsed
}
