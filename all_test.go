package kmath

import ( 
    "testing"
    //"math"
    "fmt"
)

func BenchmarkIsPrimeZ (b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPrime(0)
    } 
}

func BenchmarkIsPrimeEven (b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPrime(1000)
    } 
}

func BenchmarkIsPrime2 (b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPrime(7919)
    } 
}

func BenchmarkFirstPrimes (b *testing.B) {

    var n, tot uint64
    
    for i := 0; i < b.N; i++ {
        n = 1
        tot = 0
        for {
            n++
            if IsPrime(n) {
                tot++
                
                if tot>1000000 { 
                    break
                }
            }
        }
    } 
}

func TestPrimes (t *testing.T) {

    var i uint64
    var b bool
    for i=0; i<8000; i++ {
        b = IsPrime(i)
        if b {
            println(i)
        }
    } 

}

func TestSqrt (t *testing.T) {

    var i uint64

    for i=4; i<10; i++ {
        println(i,Sqrt(i))
    } 

}

func TestCover (t *testing.T) {

    v := make([]float64,3)
    v[0] = 50
    v[1] = 100
    v[2] = 100
    
    ad,au := Cover(v,50)
    
    fmt.Println("cover",ad,au)
}

func TestCorr (t *testing.T) {

    x := make([]float64,3)
    x[0] = 100
    x[1] = 50
    x[2] = 200
    
    y := make([]float64,3)
    y[0] = 50
    y[1] = 1000
    y[2] = -100
    
    c,_ := Corr(x,y)
    fmt.Println("corr",c)
}




