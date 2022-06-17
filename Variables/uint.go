package main

import "fmt"

// uint

func main() {

	// 0<= uint8 / byte <= 255
	var one uint8 = 255

	// 0 <= uint16 <= 65535
	var two uint16 = 65535

	// 0 <= uint32 <= 4294967295
	var three uint32 = 4294967295

	// 0 <= uint64 <= 18446744073709551615
	var four uint64 = 18446744073709551615


	fmt.Println(one, two, three, four)
}

