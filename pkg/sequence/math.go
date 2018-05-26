package sequence

import (
	"fmt"
	"math/big"
)

// Tuple represents a tuple.
type Tuple struct {
	a, b *big.Int
}

// Algorithm is an interface describe the methods required by algorithms.
type Algorithm interface {
	F(*big.Int) *Tuple
}

// FastDoublingMethod implements the Algorithm interface. It uses the
// fast doubling method algorithm:
//
// 	F(2n) = F(n)[2*F(n+1) – F(n)]
// 	F(2n + 1) = F(n)2 + F(n+1)2
//
// where, F(0) = 0, F(1) = 1
//
type FastDoublingMethod struct{}

// Fibonacci calculates the fibonacci number for the given number.
func Fibonacci(n int64, algo Algorithm) (k *big.Int, err error) {
	if n < 0 {
		err = fmt.Errorf("n must be greater than zero")
		return nil, err
	}

	tuple := algo.F(big.NewInt(n))
	k = tuple.a

	return k, nil
}

// F implements the Algorithm interface.
func (f *FastDoublingMethod) F(n *big.Int) *Tuple {
	if n.Cmp(big.NewInt(0)) == 0 {
		return &Tuple{big.NewInt(0), big.NewInt(1)}
	}

	var quotient big.Int
	quotient.Div(n, big.NewInt(2))
	tuple := f.F(&quotient)

	fOf2n := func(t *Tuple) *big.Int {
		// t.a * (t.b*2 - t.a) <=> z * (x*2 - y)
		var x, y, z big.Int
		x.Mul(t.b, big.NewInt(2))
		y.Sub(&x, t.a)
		z.Mul(t.a, &y)
		return &z

	}

	fOf2nPlus1 := func(t *Tuple) *big.Int {
		// t.a*t.a + t.b*t.b <=> x^2 + y^2
		var x, y, z big.Int
		x.Exp(t.a, big.NewInt(2), nil)
		y.Exp(t.b, big.NewInt(2), nil)
		z.Add(&x, &y)
		return &z
	}

	a := fOf2n(tuple)
	b := fOf2nPlus1(tuple)

	var mod big.Int
	mod.Mod(n, big.NewInt(2))
	if mod.Cmp(big.NewInt(0)) == 0 {
		return &Tuple{a, b}
	}

	var sum big.Int
	return &Tuple{b, sum.Add(a, b)}
}
