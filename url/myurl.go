package main

import (
	"net/url"
	"fmt"
)

func main() {

	var myUrl url.URL
	myUrl.Scheme = "https"
	myUrl.User = url.UserPassword("user", "password")
	myUrl.Host = "second.host.tld"
	myUrl.Path = "/uri/path/part"
	myUrl.Fragment = "fragment"

	query := &url.Values{}
	query.Add("array", "1")
	query.Add("array", "2")
	query.Add("array", "3")
	query.Set("first", "1")
	query.Set("second", `"is 2"`)
	query.Set("third", "Last number")
	myUrl.RawQuery = query.Encode()

	fmt.Println(myUrl.String())

}
