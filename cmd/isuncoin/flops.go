/* Info: (20241218 - Luphia)
 * FLOPS Compute Benchmark Tool
 * ----------------------------
 * Description:
 * This program calculates the total compute power of a machine in terms of FLOPS (Floating Point Operations Per Second).
 * It benchmarks CPU using Gorgonia (Go native) and NPU using Apple's Accelerate framework (CGO) on supported platforms.
 */

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/urfave/cli/v2"
	"gorgonia.org/tensor"
)

var (
	flopsCommand = &cli.Command{
		Name:        "flops",
		Usage:       "Benchmark computing power in FLOPS (CPU, GPU, NPU, TPU, APU)",
		Description: "Benchmarks the machine's floating-point performance across various computing units using external libraries.",
		Action:      runFlops,
	}
)

func runFlops(ctx *cli.Context) error {
	fmt.Println("Running FLOPS benchmark with External Libraries...")

	// 1. CPU Benchmark (Gorgonia Matrix Multiplication)
	fmt.Print("Benchmarking CPU (Gorgonia/Go)... ")
	cpuFlops := calculateCPUFlops()
	fmt.Printf("%.2f GFLOPS\n", cpuFlops/1e9)

	// 2. GPU Benchmark (Not Implemented)
	fmt.Print("Benchmarking GPU... ")
	gpuFlops := calculateGPUFlops()
	fmt.Printf("%.2f GFLOPS\n", gpuFlops/1e9)

	// 3. NPU Benchmark (Accelerate/BLAS on macOS, 0 elsewhere)
	fmt.Print("Benchmarking NPU (Apple Accelerate)... ")
	npuFlops := calculateNPUFlops()
	fmt.Printf("%.2f GFLOPS\n", npuFlops/1e9)

	// 4. TPU Benchmark (Not Implemented)
	fmt.Print("Benchmarking TPU... ")
	tpuFlops := calculateTPUFlops()
	fmt.Printf("%.2f GFLOPS\n", tpuFlops/1e9)

	// 5. APU Benchmark (Not Implemented)
	fmt.Print("Benchmarking APU... ")
	apuFlops := calculateAPUFlops()
	fmt.Printf("%.2f GFLOPS\n", apuFlops/1e9)

	// Total
	totalFlops := cpuFlops + gpuFlops + npuFlops + tpuFlops + apuFlops
	fmt.Println("------------------------------------------------")
	fmt.Printf("Total Compute Power: %.2f GFLOPS\n", totalFlops/1e9)
	fmt.Printf("Total Compute Power: %.2f TFLOPS\n", totalFlops/1e12)

	return nil
}

func calculateCPUFlops() float64 {
	// Use Gorgonia for Matrix Multiplication on CPU
	// Matrix size: N x N
	N := 1024
	backing := make([]float64, N*N)
	for i := range backing {
		backing[i] = rand.Float64()
	}

	A := tensor.New(tensor.WithShape(N, N), tensor.WithBacking(backing))
	B := tensor.New(tensor.WithShape(N, N), tensor.WithBacking(backing))

	start := time.Now()
	// Perform MatMul: C = A * B
	_, err := tensor.MatMul(A, B)
	if err != nil {
		fmt.Printf("Error: %v ", err)
		return 0
	}
	elapsed := time.Since(start).Seconds()

	// FLOPs for Matrix Multiplication of N x N is 2 * N^3 (approximately)
	ops := 2.0 * float64(N) * float64(N) * float64(N)
	return ops / elapsed
}

func calculateGPUFlops() float64 {
	// Not implemented: Requires External Libraries (e.g. CUDA/OpenCL)
	return 0
}

func calculateTPUFlops() float64 {
	// Not implemented: Requires External Libraries
	return 0
}

func calculateAPUFlops() float64 {
	// Not implemented: Requires External Libraries
	return 0
}
