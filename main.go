package main

import (
	"fmt"
	"time"
)

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	var user []User

	db.Find(&user)//db deki verileri getirme işlemi yapıyor
	db.First(&user)// db deki ilk eşleşen veriyi gettirir
	for index, usercik := range user {
		fmt.Println(index, usercik)
	}
	fmt.Println(user)
}
