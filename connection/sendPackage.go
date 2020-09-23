package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SendPackage test
func SendPackage() {
	jsonMap, errorJSONMAP := json.Marshal(map[string]string{
		"ServiceName": "hostify cli1",
		"description": "hostify cli service1",
		"ip":          "127.0.0.1:3000/cli1",
	})

	if errorJSONMAP != nil {
		log.Fatal(errorJSONMAP)
	}

	response, errorResp := http.Post("http://127.0.0.1:3000/connect",
		"application/json", bytes.NewBuffer(jsonMap))

	if errorResp != nil {
		log.Fatal(errorResp)
	}

	defer response.Body.Close()

	body, errorBody := ioutil.ReadAll(response.Body)

	if errorBody != nil {
		log.Fatal(errorBody)
	}

	fmt.Println("body str => ", string(body))

	var bodyJSON map[string]interface{}

	json.Unmarshal([]byte(body), &bodyJSON)

	fmt.Println("json => ", bodyJSON)
}
