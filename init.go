package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	_ "time"
)

const (
	HOST     = "localhost"
	DATABASE = "db-deneme"
	USER     = "postgres"
	PASSWORD = "1"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

//db baglantısını gerceklestirecegimiz alan
func init() {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	var err error
	db, err = gorm.Open(postgres.Open(vt), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println(err)
	}

	//sqlDB, _ := db.DB()
	//sqlDB.SetMaxOpenConns(10) //max bağlanılacak kişi sayısı
}
