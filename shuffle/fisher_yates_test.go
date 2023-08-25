// Based on: https://cs.opensource.google/go/go/+/refs/tags/go1.21.0:src/math/rand/rand_test.go
package shuffle

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type statsResults struct {
	mean        float64
	stddev      float64
	closeEnough float64
	maxError    float64
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func nearEqual(a, b, closeEnough, maxError float64) bool {
	absDiff := math.Abs(a - b)
	if absDiff < closeEnough { // Necessary when one value is zero and one value is close to zero.
		return true
	}
	return absDiff/max(math.Abs(a), math.Abs(b)) < maxError
}

// checkSimilarDistribution returns success if the mean and stddev of the
// two statsResults are similar.
func (this *statsResults) checkSimilarDistribution(expected *statsResults) error {
	if !nearEqual(this.mean, expected.mean, expected.closeEnough, expected.maxError) {
		s := fmt.Sprintf("mean %v != %v (allowed error %v, %v)", this.mean, expected.mean, expected.closeEnough, expected.maxError)
		fmt.Println(s)
		return errors.New(s)
	}
	if !nearEqual(this.stddev, expected.stddev, expected.closeEnough, expected.maxError) {
		s := fmt.Sprintf("stddev %v != %v (allowed error %v, %v)", this.stddev, expected.stddev, expected.closeEnough, expected.maxError)
		fmt.Println(s)
		return errors.New(s)
	}
	return nil
}

func getStatsResults(samples []float64) *statsResults {
	res := new(statsResults)
	var sum, squaresum float64
	for _, s := range samples {
		sum += s
		squaresum += s * s
	}
	res.mean = sum / float64(len(samples))
	res.stddev = math.Sqrt(squaresum/float64(len(samples)) - res.mean*res.mean)
	return res
}

func checkSampleDistribution(t *testing.T, samples []float64, expected *statsResults) {
	t.Helper()
	actual := getStatsResults(samples)
	err := actual.checkSimilarDistribution(expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

// encodePerm converts from a permuted slice of length n, such as Perm generates, to an int in [0, n!).
// See https://en.wikipedia.org/wiki/Lehmer_code.
// encodePerm modifies the input slice.
func encodePerm(s []int) int {
	// Convert to Lehmer code.
	for i, x := range s {
		r := s[i+1:]
		for j, y := range r {
			if y > x {
				r[j]--
			}
		}
	}
	// Convert to int in [0, n!).
	m := 0
	fact := 1
	for i := len(s) - 1; i >= 0; i-- {
		m += s[i] * fact
		fact *= len(s) - i
	}
	return m
}

func TestShuffle(t *testing.T) {
	rand.Seed(1)
	top := 6

	if testing.Short() {
		top = 3
	}

	for n := 3; n <= top; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			// Calculate n!.
			nfact := 1
			for i := 2; i <= n; i++ {
				nfact *= i
			}

			p := make([]int, n) // re-usable slice for Shuffle generator

			shuffle := func() int {
				// Generate permutation using Shuffle.
				for i := range p {
					p[i] = i
				}

				FisherYates(p)

				return encodePerm(p)
			}

			// Gather chi-squared values and check that they follow
			// the expected normal distribution given n!-1 degrees of freedom.
			// See https://en.wikipedia.org/wiki/Pearson%27s_chi-squared_test and
			// https://www.johndcook.com/Beautiful_Testing_ch10.pdf.
			nsamples := 10 * nfact
			if nsamples < 200 {
				nsamples = 200
			}
			samples := make([]float64, nsamples)
			for i := range samples {
				// Generate some uniformly distributed values and count their occurrences.
				const iters = 1000
				counts := make([]int, nfact)
				for i := 0; i < iters; i++ {
					counts[shuffle()]++
				}
				// Calculate chi-squared and add to samples.
				want := iters / float64(nfact)
				var χ2 float64
				for _, have := range counts {
					err := float64(have) - want
					χ2 += err * err
				}
				χ2 /= want
				samples[i] = χ2
			}

			// Check that our samples approximate the appropriate normal distribution.
			dof := float64(nfact - 1)
			expected := &statsResults{mean: dof, stddev: math.Sqrt(2 * dof)}
			errorScale := max(1.0, expected.stddev)
			expected.closeEnough = 0.10 * errorScale
			expected.maxError = 0.08
			checkSampleDistribution(t, samples, expected)
		})
	}
}
