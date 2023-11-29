package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")

	var angka int = 5
	fmt.Println(angka)

	var str string = "Hello World guyss"
	fmt.Println(str, "keren")

	var flt float64 = 3.04
	fmt.Println(flt)

	var boolean bool = true
	fmt.Println(boolean)

	angka2 := 10
	fmt.Println(angka2)

	tempString := strconv.Itoa(angka2)
	fmt.Println(tempString)

	tempAngka, err := strconv.Atoi(tempString)
	if err != nil {
		fmt.Println("Terjadi kesalahan dalam konversi:", err)
	} else {
		fmt.Println(tempAngka)
	}
}
