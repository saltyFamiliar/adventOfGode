package aocmath

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
