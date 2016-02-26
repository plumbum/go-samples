package main

import (
	"github.com/cbroglie/mustache"
	"log"
	"fmt"
	"time"
	"github.com/Pallinder/go-randomdata"
	"github.com/ccirello/goherokuname"
)

type Person struct {
	Id int
	Login string
	FirstName string
	LastName string
	Email string
	Age int
}

type Page struct {
	Persons []Person
	Title string
}

func (p Page) Date() string {
	return time.Now().Format("2 Jan 06 15:04")
}

var template = `<h1>{{{Title}}}</h1>
<ul>
{{#Persons}}
	<li id="user{{Id}}" title="{{Login}}">{{FirstName}} {{LastName}} &lt;{{Email}}&gt;; Age: {{Age}}.</li>
{{/Persons}}
</ul>
<p>at {{Date}}</p>`

var layout = `<html>
<head>
	<title>{{Title}}</title>
</head>
<body>
{{{content}}}
<p>Footer</p>
</body>
</html>`

func main() {

	var page Page
	page.Title = "<b>Page title</b>"

	for i := range [5]struct{}{} {
		pers := Person{i,
			goherokuname.Haikunate(),
			randomdata.FirstName(randomdata.Male),
			randomdata.LastName(),
			randomdata.Email(),
			randomdata.Number(13, 77)}
		page.Persons = append(page.Persons, pers)
	}

	html, err := mustache.RenderInLayout(template, layout, page)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(html)

}
