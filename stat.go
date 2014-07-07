package kmath

import (
    "sort"
    "errors"
    "math"
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

// Cover
func Cover(v []float64, lev float64) (float64,float64) {
    var i int
    N := len(v)  

    Ad := 0.0
    Au := 0.0
    
    for i=0;i<N;i++ {
        if v[i]>=lev {
            Au += (v[i]-lev)
            Ad += lev
        } else {
            Ad += v[i]
        }
    }
    
    n := float64(N)
    return Ad/(lev*n), Au/(lev*n)
}

// Pearson returns the Pearson correlation coeficient of two arrays
func Pearson (a, b [] float64) (float64, error) {

    n := len(a)
    
    if n!=len(b) {
        return 0,errors.New("lengths differ")
    }

    ma := 0.0
    mb := 0.0
    
    for i:=0;i<n;i++ {
        ma += a[i]
        mb += b[i]
    }
    
    ma /= float64(n)
    mb /= float64(n)
    
    nu := 0.0
    d1 := 0.0
    d2 := 0.0
    
    var da, db float64
    
    for i:=0; i<n; i++ {
        da = a[i] - ma
        db = b[i] - mb
        nu += da*db
        d1 += da*da
        d2 += db*db
    }
    
    return nu / ( math.Sqrt(d1) * math.Sqrt(d2) ), nil
}