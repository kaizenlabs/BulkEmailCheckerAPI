package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var emailValidate string = "userEmail@example.com"     // <---Put in the email address to validate
	var apiKey string = "asdfasdf78gf7asdfDJSDfasdfjasdfk" // <---Put in your API v4 Key from panel.bulkemailchecker.com/rest-api-v4/

	// Use http.get to make the request

	res, err := http.PostForm("http://api-v4.bulkemailchecker.com", url.Values{"key": {apiKey}, "email": {emailValidate}})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	jsonFromHTTP, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonFromHTTP))

}
