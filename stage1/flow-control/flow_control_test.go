package flow_control

import (
	"fmt"
	"math"
	"testing"
)

func TestIfStatement(t *testing.T) {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	// math.Pow returns x^y.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
