package kmath

import (
    "math/big"
)

func FactorialBig (n int64) *big.Int {
   if n <= 0 {
      return big.NewInt(1)
   }
   if (n==0) {
      return big.NewInt(1)
   }
   bigN := big.NewInt(n)
   return bigN.Mul(bigN, FactorialBig(n-1))
}

func Factorial (n int64) int64 {
   if n <= 0 {
      return 1
   }
   return n * Factorial(n-1)
}