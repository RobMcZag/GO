package newmath

import (
	"fmt"
	"testing"
)

func TestSqrt (t *testing.T) {
	const in, out = 4, 2
	if x:=Sqrt(in); x != out {
		t.Errorf("Wrong result! Expected %v - Sqrt(%v) is not %v.", out, in, x)
	}
	//t.Logf("Correct result! Sqrt(%v) == %v.", in, out)
}
