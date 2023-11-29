package main

import (
	"fmt"
	"strconv"
)

func variable() {
	fmt.Println("tes")

	var angka int = 5
	fmt.Println(angka)

	var str string = "Hicolleagues"
	fmt.Println(str, "keren")

	var flt float64 = 3.14
	fmt.Println(flt)

	var boolean bool = true
	fmt.Println(boolean)

	//banking dev
	//membuat variabel tanpa harus mendeskripsikan tipe data
	angka2 := 10
	fmt.Println(angka2)

	//mengubah int menjadi string
	tempString := strconv.Itoa(angka2)
	fmt.Println(tempString)

	//mengubah string menjadi angka
	tempAngka, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tempAngka)
	}

}
