package main

// import (
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// )

// var digits int64 = 10
// var seed int64 = 1234567890

// func nextRand() int64 {
// 	var n string = strconv.FormatInt(int64(seed*seed), 10)
// 	for int64(len(n)) < digits*2 {
// 		n = "0" + n
// 	}
// 	var start int64 = int64(math.Floor(float64(digits / 2)))
// 	var end int64 = start + digits
// 	newSeed, err := strconv.ParseInt(n[start:end], 10, 64)
// 	if err != nil {
// 		// handle error
// 		fmt.Println(err)
// 		os.Exit(2)
// 	}
// 	seed = newSeed
// 	return seed
// }

// func nextRandFloat() float64 {
// 	var maxNumStr string
// 	for i := int64(0); i < digits; i++ {
// 		maxNumStr = "9" + maxNumStr
// 	}
// 	maxNum, err := strconv.ParseInt(maxNumStr, 10, 64)
// 	if err != nil {
// 		// handle error
// 		fmt.Println(err)
// 	}
// 	return float64(nextRand()) / float64(maxNum)
// }
