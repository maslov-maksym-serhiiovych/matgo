package operations

// VectorMultiply O(n)
func VectorMultiply(a []float64, b []float64) float64 {
	sum := 0.0

	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}

	return sum
}
