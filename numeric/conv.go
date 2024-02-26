package numeric

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

// ParseInt interprets a string s in the given base (2 to 36) and returns the corresponding value.
func ParseInt64(s string, base int) (int64, error) {
	// TODO
	return 0, nil
}

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
