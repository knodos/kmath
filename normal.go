package kmath

import (
    "math"
)

var k float64 = math.Sqrt(math.Pi * 2.0)

// NormalArea calculates the area below the normal curve.
//
// den = 2^n*n!*(2*n+1)
//
// x - x(3,5,7...) / den
//
func NormalArea_(x float64) float64 {
    
    a := x
    // xx = x*x
    n := 1.0
    f := 1.0
    
    sign := true
    
    for {
        
        f = float64(Factorial(int64(n)))*math.Pow(2,n)*(2*n+1)
        n +=2
        println(n,f,a)
        
        if sign {
            a -= math.Pow(x,n) / f
            sign = false
        } else {
            a += math.Pow(x,n) / f
            sign = true
        }
         
        if n>10 {
            break
        }
    }
    
    println("a",a)
    println("/",math.Exp(-x*x/2.0) / k)
    
    return a * math.Exp(-x*x/2.0) / k
}

func NormalArea(x float64) float64 {

    area := 0.0
    inc  := 0.0001

    for i:=0.0;i<100.0;i+=inc {
        area += Normal(i) * inc      
    }

    return area
}

// Normal returns the y value at the given x of the standard unit normal 
// distribution (mean = 0 and standard deviation = 1).
func Normal(x float64) float64 {
    return math.Exp( -x*x/2.0 ) / k
}