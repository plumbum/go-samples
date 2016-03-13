package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"time"
	"log"
	"os"
)

type User struct {
	gorm.Model
	Name     string `gorm:"unique_index"`
	Password string
	Email    *string `gorm:"unique_index"`
	Phone    *string `gorm:"unique_index"`
	Birthday *time.Time
	ActivateKey *string `gorm:"size:32;index"`
	ActivateExpire *time.Time
}

func main() {
	var err error

	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := gorm.Open("sqlite3", "./sqlite3.db")
	chk(err)
	defer db.Close()
	log.Print("Connection ok")

	// Использование прямого соединение с DB
	err = db.DB().Ping()
	chk(err)

	// Enable logging
	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Таблицы будут иметь имена в единственном числе. Т.е. без суффикса 's'
	// db.SingularTable(true)

	/*
	// Создаём таблицу, если её ещё нет
	if db.HasTable(&User{}) == false {
		fmt.Println("Create User table")
		db.CreateTable(&User{})
	}
	*/

	log.Print("Automigrate start")
	db.AutoMigrate(&User{})

	log.Print("Add new users")
	addUser(db, "John", "pass")
	addUser(db, "plumbum", "test")
	addUser(db, "tuxotronic", "mypass")
	addUser(db, "tux", "mypass")
	addUser(db, "BigTux", "mypass")

	// Select items
	var users []User // Use slice
	db.Where("name LIKE ?", "%tux%").Find(&users)

	// Okay, display response as JSON
	out, err := json.MarshalIndent(users, "", "  ") // Empty prefix and 2 spaces indent
	chk(err)
	fmt.Println(string(out))

}

func addUser(db *gorm.DB, name string, pass string) error {
	user := new(User)
	user.Name = name
	user.Password = passHash(pass)

	// ok := db.NewRecord(user)
	// log.Printf("New user %s %v\n", name, ok)
	if err := db.Create(&user).Error; err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func passHash(pass string) string {
	startTime := time.Now()
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	chk(err) // Я не представляю, какие ошибки могуть быть в ходе генерации хэша, поэтому здесь паника
	log.Print("BCrypt done in time ", time.Since(startTime))
	return string(passHash)
}

func chk(err error) {
	if err != nil {
		log.Panic(err)
	}
}
