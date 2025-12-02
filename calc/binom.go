package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	b := 1
	c := 1

	for i := 2; i <= n; i++ {
		b *= i
	}

	for i := 2; i <= k; i++ {
		c *= i
	}

	return c / b
}
