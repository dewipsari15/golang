package main

import "fmt"

// type User struct {
// 	Nama string
// 	Umur int
// }

func object() {

	// var orang User
	orang2 := User{
		Nama: "Sany",
		Umur: 20,
	}

	fmt.Println(orang2.Nama, orang2.Umur)
}
