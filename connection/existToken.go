package connection

import (
	"encoding/json"
	"fmt"
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
	// * check if the token exists
	if _, errorFile := os.Stat(hostifyFilePath); errorFile != nil {
		return false
	}

	return true
}

// ValidateToken token request
func ValidateToken() string {
	client := &http.Client{
		// * 1 minute time out
		Timeout: time.Duration(60 * time.Second),
	}

	request, errorReq := http.NewRequest("POST", "https://api-hostify-service.herokuapp.com/api/validate", nil)

	if errorReq != nil {
		io.ErrorMessage("creating server request\n" + io.Trace)
		log.Fatal(errorReq)
	}

	// * send user token
	request.Header.Add("x-access-token", handlers.GetToken())

	// * make request
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
		io.ErrorMessage(
			"rejected request, The token entered is not valid or has expired, start session in hostify and add the token again\n" + io.Trace)
		log.Fatal(bodyJSON["message"])
		return "nil"
	}

	io.SuccessMessage("the token was verified successfully")
	return fmt.Sprintf("%v", bodyJSON["username"])

}
