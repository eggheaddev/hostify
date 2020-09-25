package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

	response, errorResp := http.Post("http://localhost:3000/upload",
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

// SendFiles test
func SendFiles(url string, filename string, filetype string) []byte {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(part, file)
	writer.WriteField("hostify.json", `{
		"name": "trex",
		"description": "package manager for deno",
		"version": "1.3.0",
		"entry": "cli.ts",
		"repository": "https://github.com/crewdevio/Trex",
		"files": []
	}`)
	writer.Close()
	request, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return content
}
