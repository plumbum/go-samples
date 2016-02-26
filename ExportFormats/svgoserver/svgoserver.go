package main

import (
	"github.com/ajstarks/svgo"
	"net/http"
	"time"
	"log"
	"fmt"
)

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/circle", http.HandlerFunc(circle))
	fmt.Println("Listen at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Refresh", "5")
	w.Write([]byte(`<img src="/circle" />`))
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	width := 500
	height := 500
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.Text(width/2, height/4, time.Now().Format("Jan _2 15:04:05 2006"), "text-anchor:middle;font-family:'Droid';font-size:30px;fill:blue")
	canvas.End()
}

