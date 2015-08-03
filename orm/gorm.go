package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"time"
	"log"
)

type User struct {
    gorm.Model
    Name string
	Password string
}

func main() {
	var err error
	db, err := gorm.Open("sqlite3", "./sqlite3.db")
	chk(err)
	defer db.Close()

	err = db.DB().Ping()
	chk(err)

	// Enable logging
	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)

	if db.HasTable(&User{}) == false {
		fmt.Println("Create User table")
		db.CreateTable(&User{})
	}

	addUser(&db, "John", "pass")
	addUser(&db, "plumbum", "test")
	addUser(&db, "tuxotronic", "mypass")
	addUser(&db, "tux", "mypass")
	addUser(&db, "BigTux", "mypass")

	// Select items
	var users []User // Use slice
	db.Where("name LIKE ?", "%tux%").Find(&users)

	// Okay, display response as JSON
	out, err := json.MarshalIndent(users, "", "  ") // Empty prefix and 2 spaces indent
	chk(err)
	fmt.Println(string(out))

}


func addUser(db *gorm.DB, name string, pass string) {
	user := new(User)
	user.Name = name
	user.Password = passHash(pass)

	ok := db.NewRecord(user)
	fmt.Printf("New user %s %v\n", name, ok)
	db.Create(&user)
}

func passHash(pass string) string {
	startTime := time.Now().UnixNano()
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	chk(err)
	endTime := time.Now().UnixNano()
	log.Println("BCrypt time (ms)", (endTime - startTime)/1000000)
	return string(passHash)
}

func chk(err error) {
	if err != nil {

	}
}
