package main

import (
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
)

func main() {

	for i:=0; i<9; i++ {
		pass := uuid.NewV4()
		log.Printf("Pass: %s; BCrypt hash %s", pass.String(), passHash(pass.String()))
	}

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
