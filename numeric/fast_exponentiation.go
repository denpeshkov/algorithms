package numeric

import "math/big"

// FastExp computes a^n using exponentiation by squaring (fast exponentiation).
// a and n are non-negative integers.
func FastExp(a, n uint64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	x := FastExp(a, n/2)
	x.Mul(x, x) // x^2

	if n%2 == 0 {
		return x
	} else {
		return x.Mul(new(big.Int).SetUint64(a), x)
	}
}
