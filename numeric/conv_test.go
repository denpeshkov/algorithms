// Based on https://github.com/golang/go/blob/master/src/strconv/itoa_test.go
// Based on https://github.com/golang/go/blob/master/src/strconv/atoi_test.go

package numeric

import (
	"errors"
	"fmt"
	"testing"
)

type parseIntTest struct {
	in   string
	base int
	out  int64
	err  error
}

var parseIntTests = []parseIntTest{
	{"", 10, 0, ErrSyntax},
	{"0", 10, 0, nil},
	{"-0", 10, 0, nil},
	{"+0", 10, 0, nil},
	{"1", 10, 1, nil},
	{"-1", 10, -1, nil},
	{"+1", 10, 1, nil},
	{"12345", 10, 12345, nil},
	{"-12345", 10, -12345, nil},
	{"012345", 10, 12345, nil},
	{"-012345", 10, -12345, nil},
	{"98765432100", 10, 98765432100, nil},
	{"-98765432100", 10, -98765432100, nil},
	{"9223372036854775807", 10, 1<<63 - 1, nil},
	{"-9223372036854775807", 10, -(1<<63 - 1), nil},
	{"9223372036854775808", 10, 1<<63 - 1, ErrRange},
	{"-9223372036854775808", 10, -1 << 63, nil},
	{"9223372036854775809", 10, 1<<63 - 1, ErrRange},
	{"-9223372036854775809", 10, -1 << 63, ErrRange},
	{"-1_2_3_4_5", 10, 0, ErrSyntax}, // base=10 so no underscores allowed
	{"-_12345", 10, 0, ErrSyntax},
	{"_12345", 10, 0, ErrSyntax},
	{"1__2345", 10, 0, ErrSyntax},
	{"12345_", 10, 0, ErrSyntax},
	{"123%45", 10, 0, ErrSyntax},
	{"012345", 10, 12345, nil},
	{"000000000012345", 10, 12345, nil},

	// other bases
	{"g", 17, 16, nil},
	{"10", 25, 25, nil},
	{"holycow", 35, (((((17*35+24)*35+21)*35+34)*35+12)*35+24)*35 + 32, nil},
	{"holycow", 36, (((((17*36+24)*36+21)*36+34)*36+12)*36+24)*36 + 32, nil},

	// base 2
	{"0", 2, 0, nil},
	{"-1", 2, -1, nil},
	{"1010", 2, 10, nil},
	{"1000000000000000", 2, 1 << 15, nil},
	{"111111111111111111111111111111111111111111111111111111111111111", 2, 1<<63 - 1, nil},
	{"1000000000000000000000000000000000000000000000000000000000000000", 2, 1<<63 - 1, ErrRange},
	{"-1000000000000000000000000000000000000000000000000000000000000000", 2, -1 << 63, nil},
	{"-1000000000000000000000000000000000000000000000000000000000000001", 2, -1 << 63, ErrRange},

	// base 8
	{"-10", 8, -8, nil},
	{"57635436545", 8, 057635436545, nil},
	{"100000000", 8, 1 << 24, nil},

	// base 16
	{"10", 16, 16, nil},
	{"-123456789abcdef", 16, -0x123456789abcdef, nil},
	{"7fffffffffffffff", 16, 1<<63 - 1, nil},
}

func TestParseInt(t *testing.T) {
	for i := range parseIntTests {
		test := &parseIntTests[i]
		out, err := ParseInt(test.in, test.base)
		if test.out != out || !errors.Is(test.err, err) {
			t.Errorf("ParseInt(%q, %v, 64) = %v, %v; want %v, %v",
				test.in, test.base, out, err, test.out, test.err)
		}
	}
}

func BenchmarkParseInt(b *testing.B) {
	b.Run("Pos", func(b *testing.B) {
		benchmarkParseInt(b, 1)
	})
	b.Run("Neg", func(b *testing.B) {
		benchmarkParseInt(b, -1)
	})
}

type benchCase struct {
	name string
	num  int64
}

func benchmarkParseInt(b *testing.B, neg int) {
	cases := []benchCase{
		{"7bit", 1<<7 - 1},
		{"26bit", 1<<26 - 1},
		{"31bit", 1<<31 - 1},
		{"56bit", 1<<56 - 1},
		{"63bit", 1<<63 - 1},
	}
	for _, cs := range cases {
		b.Run(cs.name, func(b *testing.B) {
			s := fmt.Sprintf("%d", cs.num*int64(neg))
			for i := 0; i < b.N; i++ {
				out, _ := ParseInt(s, 10)
				BenchSink += int(out)
			}
		})
	}
}

type itob64Test struct {
	in   int64
	base int
	out  string
}

var itob64tests = []itob64Test{
	{0, 10, "0"},
	{1, 10, "1"},
	{-1, 10, "-1"},
	{12345678, 10, "12345678"},
	{-987654321, 10, "-987654321"},
	{1<<31 - 1, 10, "2147483647"},
	{-1<<31 + 1, 10, "-2147483647"},
	{1 << 31, 10, "2147483648"},
	{-1 << 31, 10, "-2147483648"},
	{1<<31 + 1, 10, "2147483649"},
	{-1<<31 - 1, 10, "-2147483649"},
	{1<<32 - 1, 10, "4294967295"},
	{-1<<32 + 1, 10, "-4294967295"},
	{1 << 32, 10, "4294967296"},
	{-1 << 32, 10, "-4294967296"},
	{1<<32 + 1, 10, "4294967297"},
	{-1<<32 - 1, 10, "-4294967297"},
	{1 << 50, 10, "1125899906842624"},
	{1<<63 - 1, 10, "9223372036854775807"},
	{-1<<63 + 1, 10, "-9223372036854775807"},
	{-1 << 63, 10, "-9223372036854775808"},

	{0, 2, "0"},
	{10, 2, "1010"},
	{-1, 2, "-1"},
	{1 << 15, 2, "1000000000000000"},

	{-8, 8, "-10"},
	{057635436545, 8, "57635436545"},
	{1 << 24, 8, "100000000"},

	{16, 16, "10"},
	{-0x123456789abcdef, 16, "-123456789abcdef"},
	{1<<63 - 1, 16, "7fffffffffffffff"},
	{1<<63 - 1, 2, "111111111111111111111111111111111111111111111111111111111111111"},
	{-1 << 63, 2, "-1000000000000000000000000000000000000000000000000000000000000000"},

	{16, 17, "g"},
	{25, 25, "10"},
	{(((((17*35+24)*35+21)*35+34)*35+12)*35+24)*35 + 32, 35, "holycow"},
	{(((((17*36+24)*36+21)*36+34)*36+12)*36+24)*36 + 32, 36, "holycow"},
}

func TestFormatInt(t *testing.T) {
	for _, test := range itob64tests {
		s := FormatInt(test.in, test.base)
		if s != test.out {
			t.Errorf("FormatInt(%v, %v) = %v; want %v", test.in, test.base, s, test.out)
		}
	}
}

var BenchSink int // make sure compiler cannot optimize away benchmarks

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range itob64tests {
			s := FormatInt(test.in, test.base)
			BenchSink += len(s)
		}
	}
}
