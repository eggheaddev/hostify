package connection

import (
	"encoding/json"
	"fmt"
	"hostify/handlers"
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
		Timeout: time.Duration(5 * time.Second),
	}

	request, errorReq := http.NewRequest("GET", "https://api-hostify-service.herokuapp.com/api/validate", nil)

	if errorReq != nil {
		log.Fatal(errorReq)
	}

	request.Header.Add("x-access-token", handlers.GetToken())
	resp, errorGet := client.Do(request)

	if errorGet != nil {
		log.Fatal(errorGet)
	}

	body, errorBody := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if errorBody != nil {
		log.Fatal(errorBody)
	}

	var bodyJSON map[string]interface{}

	json.Unmarshal(body, &bodyJSON)

	fmt.Println("body => ", string(body))
	fmt.Println("json => ", bodyJSON)
}
