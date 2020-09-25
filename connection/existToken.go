package connection

import (
	"encoding/json"
	"hostify/handlers"
	"hostify/io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// ExistToken verify that the hostify.key file exists
func ExistToken() bool {

	hostifyFilePath := handlers.TokenPath()

	if _, errorFile := os.Stat(hostifyFilePath); errorFile != nil {
		return false
	}

	return true

}

// ValidateToken token request
func ValidateToken() {
	client := &http.Client{
		// * 1 minute time out
		Timeout: time.Duration(60 * time.Second),
	}

	request, errorReq := http.NewRequest("GET", "https://api-hostify-service.herokuapp.com/api/validate", nil)

	if errorReq != nil {
		io.ErrorMessage("creating server request\n" + io.Trace)
		log.Fatal(errorReq)
	}

	request.Header.Add("x-access-token", handlers.GetToken())
	resp, errorGet := client.Do(request)

	if errorGet != nil {
		io.ErrorMessage("making request for validate user token\n" + io.Trace)
		log.Fatal(errorGet)
	}

	body, errorBody := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if errorBody != nil {
		io.ErrorMessage("geting the server response\n" + io.Trace)
		log.Fatal(errorBody)
	}

	var bodyJSON map[string]interface{}

	json.Unmarshal(body, &bodyJSON)

	if bodyJSON["error"] == true {
		io.ErrorMessage("rejected request")
		log.Fatal(bodyJSON["message"])
	} else {
		io.SuccessMessage("the token was verified successfully")
	}
}
