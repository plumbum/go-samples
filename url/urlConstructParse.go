package main

import (
	"net/url"
	"log"
	"github.com/kr/pretty"
	"strings"
)

func main() {

	params := &url.Values{}
	params.Set("first", "1")
	params.Set("second", `"is quoted string"`)
	params.Set("third", "Last number")
	params.Set("fourth", "4") // WARNING! All keys sorted by alphabet
	params.Add("array", "1")
	params.Add("array", "2")
	params.Add("array", "3")

	myUrl := url.URL{
		Scheme: "https",
		User: url.UserPassword("user", "password"),
		Host: "second.host.tld",
		Path: "/uri/path/part",
		RawQuery: params.Encode(),
		Fragment: "fragment",
	}

	pretty.Println("Create URI string:", myUrl.String())

	parsedUrl, err := url.Parse(myUrl.String()+"_tail")
	if err != nil {
		log.Fatalln(err)
	}

	pretty.Println(strings.Repeat("=", 80))
	pretty.Println("Parse URI string")
	pretty.Println(parsedUrl)
	pretty.Println(parsedUrl.Query())

}
