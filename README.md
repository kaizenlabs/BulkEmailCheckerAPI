# Bulk Email Checker API for GoLang

This code is made to interact with Bulk Email Checker's RESTful API (v4) with GoLang. To get an API you will need to create an account [here] and then purchase validation credits to interact with the API.

**BulkEmailChecker API Summary:**

>This RESTful API uses 1 validation per successful attempt and can be integrated into nearly any website or application using the programming language of your choice. Your application needs to make a simple HTTP GET or POST request to the RESTful API with 2 requirements, your api key and the email address to be validated. The api will respond with a JSON or XML response for the result of the verification you just requested. Please see "Main Status Responses" for a list of all possible responses you may receive when querying the RESTful API. 

### Response Types

**Successful Response:**

>Status="passed"- A passed response is an address that has passed all tests.  
>--Event="mailbox_exists"- The address provided passed all tests.

**Failed Response:**

>Status="failed"- A failed response is an address that has failed 1 or more tests.  
--Event="mailbox_does_not_exist"- The address provided does not exist.  
--Event="mailbox_full"- The mailbox is full.  
--Event="invalid_syntax"- The address provided is not in a valid format.  
--Event="domain_does_not_exist"- The address provided does not have a valid dns.  
--Event="mxserver_does_not_exist"- The address provided does not have a mx server.  
--Event="blacklisted_external"- The address provided was found in your (Mailchimp, Postmark, Sendgrid, etc.) blacklist.  
--Event="blacklisted_internal"- The address provided was found in your internal blacklist containing previously failed addresses.  
--Event="blacklisted_domain"- The domain provided was found in your domain blacklist.  

**Unknown Response:**

>Status="unknown"- A unknown response is an address that can not be accurately tested.  
--Event="is_catchall"- Is a Catchall mx server config.  
--Event="is_greylisting"- Greylisting is active on this server.  
--Event="inconclusive"- Transient error, please try again later.  

**Additional Responses:**

>"email"- The address that was validated.  
"emailSuggested"- The address contained a typo and we have suggested the most likely repaired version.  
"mailbox"- The mailbox that was validated.  
"domain"- The domain that was validated.  
"mxIp"- The ip address of the mail server that was validated.  
"mxLocation"- The location of the mail server that was validated.  
"isComplainer"- The address is a frequent complainer.  
"isDisposable"- The address is a disposable email address.   
"isFreeService"- The address is a free email account.  
"isOffensive"- The address contains bad words.  
"isRoleAccount"- The address is a role account.  
"validationsRemaining"- The amount of validations remaining in your account balance.   
"hourlyQuotaRemaining"- The amount of api requests remaining in your account for this hourly segment.    
"execution"- The amount of time taken to validate this address in ms.  

### JSON Response Example

>{  
"status": "passed",  
"event": "mailbox_exists",
"details": "The address provided passed all tests.",  
"email": "userEmail@example.com",  
"emailSuggested": "",  
"mailbox": "userEmail",  
"domain": "example.com",  
"mxIp": "54.124.145.16",  
"mxLocation": "US",  
"isComplainer": false,  
"isDisposable": false,  
"isFreeService": true,  
"isOffensive": false,  
"isRoleAccount": false,  
"validationsRemaining": 2120,  
"hourlyQuotaRemaining": 1499,  
"execution": 0.097  
}



### Version
BulkEmailChecker API v4

### Tech

This repository uses the following technologies.

* [GoLang] 
* [JSON] 

### GoLang Packages Used
Main executable file uses the following packages from the GoLang standard library:
* [encoding/json] - Encodes & Decodes JSON Objects 
* [fmt] - Implements formatted I/O with functions analogous to C's printf and scanf.
* [io/ioutil] - Implements some I/O utility functions.
* [log] - Implements a simple logging package. 
* [net/http] - Provides HTTP client and server implementations.
* [net/url] - Parses URLs and implements query escaping.
* [strconv] - Implements conversions to and from string representations of basic data types.

### Installation

You need to install the Go Lang on your machine. For details on how to do this, check out
[GoLang].org


### Development

Written by John Anthony Radosta  
Advanced Sports Technology, Inc.




[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)


   [GoLang]: <https://golang.org>
   [git-repo-url]: <https://github.com/JohnAntonusMaximus/BulkEmailCheckerAPI>
   [AdvancedSportsTechnology]: <http://advancedsportstech.com>
   [LinkedIn]: <http://www.linkedin.com/in/johnradosta>
   [fmt]: <https://golang.org/pkg/fmt/>
   [Github]: <https://github.com/JohnAntonusMaximus/>
   [encoding/json]: <https://golang.org/pkg/encoding/json/>
   [io/ioutil]: <https://golang.org/pkg/io/ioutil/>
   [log]: <https://golang.org/pkg/log/>
   [net/http]: <https://golang.org/pkg/net/http/>
   [net/url]: <https://golang.org/pkg/net/url/>
   [strconv]: <https://golang.org/pkg/strconv/>
   [GithubReadme]: <https://github.com/JohnAntonusMaximus/BulkEmailCheckerAPI/blob/master/README.md>
   [JSON]:  <http://www.json.org/>
   [here]: <https://panel.bulkemailchecker.com/create-account/>


