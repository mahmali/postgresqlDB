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

	db.Find(&user) //db deki ilk veriyi getirme işlemi yapıyor
	for index, usercik := range user {
		fmt.Println(index, usercik)
	}
	fmt.Println(user)
}
