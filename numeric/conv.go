package numeric

import (
	"errors"
	"math"
)

var ErrRange = errors.New("value out of range")
var ErrSyntax = errors.New("invalid syntax")

// ParseInt interprets a string s in the given base, for 2 <= base <= 36 and returns the corresponding value.
// The string may begin with a leading sign: "+" or "-".
func ParseInt(s string, base int) (int64, error) {
	if len(s) == 0 {
		return 0, ErrSyntax
	}

	var ui uint64
	isNeg := false

	if s[0] == '-' {
		isNeg = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	var d byte
	// conversion optimized and no allocations
	for _, c := range []byte(s) {
		switch {
		case c >= '0' && c <= '9':
			d = c - '0'
		case lower(c) >= 'a' && lower(c) <= 'z':
			d = lower(c) - 'a' + 10
		default:
			return 0, ErrSyntax
		}
		ui = ui*uint64(base) + uint64(d)
	}
	if !isNeg {
		if ui > math.MaxInt64 {
			return math.MaxInt64, ErrRange
		}
		return int64(ui), nil
	}
	if ui > -math.MinInt64 {
		return math.MinInt64, ErrRange
	}
	return -int64(ui), nil
}

func lower(c byte) byte {
	return c | ('a' - 'A')
}

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

// FormatInt returns the string representation of i in the given base, for 2 <= base <= 36.
// The result uses the lower-case letters 'a' to 'z' for digit values >= 10.
func FormatInt(i int64, base int) string {
	if base < 2 || base > 36 {
		panic("FormatInt: illegal base")
	}

	// uint64 because -MaxInt64 == MaxInt64 in two's complement
	ui := uint64(i)
	neg := false
	if i < 0 {
		neg = true
		ui = -ui
	}

	var b [64 + 1]byte // +1 for sign of 64 bit value in base 2
	j := len(b) - 1
	for {
		b[j] = digits[ui%uint64(base)]
		ui /= uint64(base)
		if ui == 0 {
			break
		}
		j--
	}
	if neg {
		j--
		b[j] = '-'
	}
	return string(b[j:])
}
