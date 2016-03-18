package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/abbot/go-http-auth"
	"golang.org/x/net/webdav"
)

func GetSecret(user, realm string) string {
	log.Printf("User: %s; Realm: %s\n", user, realm)
	if user == "john" {
		// password is "hello"
		return "$2y$05$DEKfYsAg/ROn8n4TXE3J3OxlnbHr6RyFVd.RX1OMsLCwVOPhpUuKS" // bcrypt
	}
	return ""
}

func main() {
	authenticator := auth.NewBasicAuthenticator("webdav", GetSecret)

	h := new(webdav.Handler)
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	h.FileSystem = webdav.Dir(root)
	h.LockSystem = webdav.NewMemLS()
	h.Prefix = "/dav/"
	h.Logger = func(r *http.Request, e error) {
		log.Printf("%s(%s): %s", r.Method, r.RemoteAddr, r.RequestURI)
		if err != nil {
			log.Println("[ERROR] ", err)
		}
	}

	//then use the Handler.ServeHTTP Method as the http.HandleFunc
	http.HandleFunc("/dav/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		default:
			if username := authenticator.CheckAuth(r); username == "" {
				authenticator.RequireAuth(w, r)
				return
			} else {
			}
		case http.MethodGet, http.MethodOptions, "PROPFIND":
			// No need auth
		}
		h.ServeHTTP(w, r)
		log.Print("...done")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=UTF-8")
		w.Write([]byte(fmt.Sprintf(`Time: %s`, time.Now().String())))
	})

	log.Println("Listen at :5555")
	http.ListenAndServe(":5555", nil)
}
