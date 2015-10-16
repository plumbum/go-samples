package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/context"
	auth "github.com/plumbum/go-http-auth"
	"github.com/k0kubun/pp"
	"golang.org/x/crypto/bcrypt"
"log"
)

var Users = make(map[string]string);

func getHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func init() {
	log.Print("Generate password hashes")
	Users["user"], _ = getHash("hello")
	Users["john"], _ = getHash("hello")
	Users["bill"], _ = getHash("hello")
	Users["max"], _ = getHash("hello")
	pp.Println (Users)
}

func Secret(user, realm string) string {
	if hash, ok := Users[user]; ok {
		return hash;
	}
	return ""
}

type ContextHandler interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type ContextHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

func (f ContextHandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	f(ctx, w, r)
}

func handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pp.Print(ctx)
	authInfo := auth.FromContext(ctx)
	authInfo.UpdateHeaders(w.Header())
	if authInfo == nil || !authInfo.Authenticated {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", authInfo.Username)
}

func authenticatedHandler(a auth.AuthenticatorInterface, h ContextHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := a.NewContext(context.Background(), r)
		h.ServeHTTP(ctx, w, r)
	})
}

func main() {
	authenticator := auth.NewBasicAuthenticator("example.com", Secret)
	http.Handle("/", authenticatedHandler(authenticator, ContextHandlerFunc(handle)))
	log.Println("Start listen at :8080")
	http.ListenAndServe(":8080", nil)
}

