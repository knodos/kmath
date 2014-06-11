package kmath

import (
    "sort"
)

// Gini returns the Gini coefficient. It is a measure of inequality in a 
// frequency distribution. 0 indicates perfect equality, 1 maximum inequality.
//
// See https://en.wikipedia.org/wiki/Gini_coefficient
func Gini(in []float64) float64 {

    var i int
    N := len(in)  
    
	// Get the total of the series
	c := 0.0
	for i=0;i<N;i++ {
		c += in[i]
	}

	// Create a vector Y with what part of the total is each input value. 
	y := make([]float64, N)

	for i=0;i<N;i++ {
		y[i] = in[i] / c
	}

    // sort Y values in ascending order
	sort.Float64s(y)

	// Normalize the X
	x := 1.0 / float64(N)

	// Calculate areas below the diagonal
	gini := 0.0
	acc := 0.0
	pacc := 0.0
	
	for i := 0; i < N; i++ {
		acc += y[i]
		pacc += x
		gini += (pacc - acc) * x
	}

    // Normalize to 0...1 (area of triangle is 0.5 max)
	return 2 * gini
}