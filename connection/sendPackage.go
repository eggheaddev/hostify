package connection

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"hostify/handlers"
	colors "hostify/io"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SendPackage test
// func SendPackage() {
// 	jsonMap, errorJSONMAP := json.Marshal(map[string]string{
// 		"ServiceName": "hostify cli1",
// 		"description": "hostify cli service1",
// 		"ip":          "127.0.0.1:3000/cli1",
// 	})

// 	if errorJSONMAP != nil {
// 		log.Fatal(errorJSONMAP)
// 	}

// 	response, errorResp := http.Post("http://localhost:3000/api/upload",
// 		"application/json", bytes.NewBuffer(jsonMap))

// 	if errorResp != nil {
// 		log.Fatal(errorResp)
// 	}

// 	defer response.Body.Close()

// 	body, errorBody := ioutil.ReadAll(response.Body)

// 	if errorBody != nil {
// 		log.Fatal(errorBody)
// 	}

// 	var bodyJSON map[string]interface{}

// 	json.Unmarshal([]byte(body), &bodyJSON)

// 	if bodyJSON["error"] == true {
// 		log.Fatal(bodyJSON)
// 	} else {
// 		fmt.Println("done")
// 	}

// }

func uploadFile(root string, path string) []byte {

	file, errorOpen := os.Open(root)

	if errorOpen != nil {
		colors.ErrorMessage(
			"an error occurred opening the file" + colors.Trace)
		log.Fatal(errorOpen)
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, errorWriter := writer.CreateFormFile("file", filepath.Base(file.Name()))

	if errorWriter != nil {
		colors.ErrorMessage(
			"There was an error creating the file submission form" + colors.Trace)
		log.Fatal(errorWriter)
	}

	io.Copy(part, file)

	writer.Close()
	request, errorReq := http.NewRequest("POST", "http://localhost:3000/api/upload", body)

	if errorReq != nil {
		colors.ErrorMessage(
			"an error occurred creating the request" + colors.Trace)
		log.Fatal(errorReq)
	}

	hostifyFields := handlers.ReadJSON()

	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("package-name", fmt.Sprintf("%v", hostifyFields["name"]))
	request.Header.Add("file-path", strings.ReplaceAll(path, "\\", "/"))
	request.Header.Add("package-version", fmt.Sprintf("%v", hostifyFields["version"]))
	request.Header.Add("package-repository", fmt.Sprintf("%v", hostifyFields["repository"]))

	client := &http.Client{}
	response, errorClient := client.Do(request)

	if errorClient != nil {
		colors.ErrorMessage(
			"making the request to send the files" + colors.Trace)
		log.Fatal(errorClient)
	}
	defer response.Body.Close()

	content, erroRead := ioutil.ReadAll(response.Body)

	if erroRead != nil {
		colors.ErrorMessage(
			"an error occurred reading the response from the server" + colors.Trace)
		log.Fatal(erroRead)
	}

	return content
}

// SendFiles test
func SendFiles(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			time.Sleep(900 * time.Millisecond)
			fmt.Printf(
				"%v uploading:%v %v %v %v \n", colors.Green,
				colors.Reset, colors.Yellow, path, colors.Reset)

			defer uploadFile(filepath.Join(handlers.Cwd(), path), path)
		}
		return nil
	})
}
