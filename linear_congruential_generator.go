// https://en.wikipedia.org/wiki/Linear_congruential_generator
package main

import (
	"math"
)

var a int64 = 1664525
var c int64 = 1013904223
var m int64 = int64(math.Pow(2, 32.0))
var seed int64 = 1

func nextRand() int64 {
	seed = (a*seed + c) % m
	return seed
}

func nextRandFloat() float64 {
	return float64(nextRand()) / float64(m)
}

// func main() {
// 	fmt.Println("a", a)
// 	fmt.Println("c", c)
// 	fmt.Println("m", m)
// 	fmt.Println("seed", seed)

// }
