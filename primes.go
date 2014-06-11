package kmath

import "math"

func IsPrime (n uint64) bool {
    
    if n<2 {
        return false
    }
    if n==2 {
        return true
    }
    if n&1 == 0 {
        return false
    }

    max := uint64(math.Sqrt(float64(n)))
    
    var i uint64
    
    for i=2; i<=max; i++ {
        if n%i==0 {
            return false
        }
    }
    return true
}

// Not accurate and not faster than Go's own math.Sqrt
func Sqrt (n uint64) uint64 {

     switch (n) {
     case 0: return 0
     case 1: return 0
     case 2: return 1
     case 3: return 1
     }

     x0 := n/2
     var x uint64
     
     for {
         x = (x0+n/x0)/2
         if x>=x0 {
             break
         }
         x0 = x
         
     }
     
     return x
} 
