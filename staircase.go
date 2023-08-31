package main

import "fmt"

//4 0 3 1
//4 1 2 2
//4 2 1 3
//4 3 0

// 4 - 0 + 1 = 3
func main() {
	var n int32 = 4
	var i int32
	for ; i < n; i++ {
		j := int32(0)
		for ; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for ; j < n; j++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
