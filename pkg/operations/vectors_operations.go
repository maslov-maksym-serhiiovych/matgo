package operations

// VectorMultiply O(n)
func VectorMultiply(a []float64, b []float64) float64 {
	sum := 0.0

	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}

	return sum
}

// VectorDivScalar O(n)
func VectorDivScalar(a []float64, b float64) []float64 {
	n := len(a)
	division := make([]float64, n)

	for i := 0; i < n; i++ {
		division[i] = a[i] / b
	}

	return division
}
