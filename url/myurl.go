package main

import (
	"net/url"
	"fmt"
)

func main() {

	params := &url.Values{}
	params.Set("first", "1")
	params.Set("second", `"is 2"`)
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

	fmt.Println(myUrl.String())

}
