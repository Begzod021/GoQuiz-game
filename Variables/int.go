package main

import "fmt"


func main(){
	//-128 <= int8 < 127=
	var one int8 = -123

	// -32768 <= int16 <= 32767
	var two int16 = 32755

	// -2147483648 <= int32 <= 2147483647
	var three int32 = -214756666


	// -9223372036854775808 <= int64 <= 9223372036854775807
	var four int64 = 9223372036854775807

	fmt.Println(one, two, three, four)

}