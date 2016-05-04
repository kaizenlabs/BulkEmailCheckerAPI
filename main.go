package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//BulkEmailAPIResponse creates a struct in the format of the JSON response from bulkemailchecker.com
type BulkEmailAPIResponse struct {
	Status               string  `json:"status"`
	Event                string  `json:"event"`
	Details              string  `json:"details"`
	Email                string  `json:"email"`
	EmailSuggested       string  `json:"emailSuggested"`
	Mailbox              string  `json:"mailbox"`
	Domain               string  `json:"domain"`
	MxIP                 string  `json:"mxIp"`
	MxLocation           string  `json:"mxLocation"`
	IsComplainer         bool    `json:"isComplainer"`
	IsDisposable         bool    `json:"isDisposable"`
	IsFreeService        bool    `json:"isFreeService"`
	IsOffensive          bool    `json:"isOffensive"`
	IsRoleAccount        bool    `json:"isRoleAccount"`
	ValidationsRemaining int     `json:"validationsRemaining"`
	HourlyQuotaRemaining int     `json:"hourlyQuotaRemaining"`
	ExecutionTime        float64 `json:"execution"`
}

func getResponse(body []byte) (*BulkEmailAPIResponse, error) {
	var res = new(BulkEmailAPIResponse)
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res, err
}

func main() {
	var emailValidate string = "example@example.com"  // <---Put in the email address to validate
	var apiKey string = "xyzGYdh3234BdksadfDajad4542" // <---Put in your API v4 Key from panel.bulkemailchecker.com/rest-api-v4/

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

	results, err := getResponse([]byte(jsonFromHTTP))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		"Email: "+results.Email+"\n",
		"Status: "+results.Status+"\n",
		"Event: "+results.Event+"\n",
		"Details: "+results.Details+"\n",
		"Email Suggested: "+results.EmailSuggested+"\n",
		"Domain: "+results.Domain+"\n",
		"Mailbox: "+results.Mailbox+"\n",
		"Mailbox IP: "+results.MxIP+"\n",
		"Mailbox Location: "+results.MxLocation+"\n",
		"Is Complainer? "+strconv.FormatBool(results.IsComplainer)+"\n",
		"Is Disposable? "+strconv.FormatBool(results.IsDisposable)+"\n",
		"Is Free Service? "+strconv.FormatBool(results.IsFreeService)+"\n",
		"Is Offensive? "+strconv.FormatBool(results.IsOffensive)+"\n",
		"Is Role Account? "+strconv.FormatBool(results.IsRoleAccount)+"\n",
		"-------------------------------"+"\n",
		"Validations Remaining: "+strconv.Itoa(results.ValidationsRemaining)+"\n",
		"Hourly Quota Remaining: "+strconv.Itoa(results.HourlyQuotaRemaining)+"\n",
		"Execution Time: "+strconv.FormatFloat(results.ExecutionTime, 'f', -1, 32)+"\n",
	)

}
