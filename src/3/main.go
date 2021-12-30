package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i += 1 {
		fmt.Printf("%d ", rand5())
	}
	fmt.Printf("\n")
	for i := 0; i < 100; i += 1 {
		fmt.Printf("%d ", rand13())
	}
	fmt.Printf("\n")
}

//已知
func _rand13() int32 {
	return rand.Int31()%13 + 1
}

func rand5() int32 {
	x := int32(math.MaxInt32)
	for x > 10 {
		x = _rand13()
	}
	return x%5 + 1
}
func rand13() int32 {
	x := int32(math.MaxInt32)
	for x > 26 {
		x = (rand5()-1)*6 + rand5()
	}
	return x%13 + 1
}
