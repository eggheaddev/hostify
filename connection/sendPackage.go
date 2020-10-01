package connection

import (
	"bytes"
	"encoding/json"
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

func uploadFile(root string, path string, user string, done bool) []byte {

	file, errorOpen := os.Open(root)

	if errorOpen != nil {
		colors.ErrorMessage(
			"an error occurred opening the file\n" + colors.Trace)
		log.Fatal(errorOpen)
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, errorWriter := writer.CreateFormFile("file", filepath.Base(file.Name()))

	if errorWriter != nil {
		colors.ErrorMessage(
			"There was an error creating the file submission form\n" + colors.Trace)
		log.Fatal(errorWriter)
	}

	io.Copy(part, file)

	writer.Close()
	request, errorReq := http.NewRequest(
		"POST", "https://storage-hostify-service.herokuapp.com/api/upload", body)

	if errorReq != nil {
		colors.ErrorMessage(
			"an error occurred creating the request\n" + colors.Trace)
		log.Fatal(errorReq)
	}

	hostifyFields := handlers.ReadJSON()

	var finish string = "null"

	if done {
		finish = "finish"
	}

	request.Header.Add("package-description", fmt.Sprintf("%v", hostifyFields["description"]))
	request.Header.Add("package-repository", fmt.Sprintf("%v", hostifyFields["repository"]))
	request.Header.Add("package-version", fmt.Sprintf("%v", hostifyFields["version"]))
	request.Header.Add("package-name", fmt.Sprintf("%v", hostifyFields["name"]))
	request.Header.Add("entry-file", fmt.Sprintf("%v", hostifyFields["entry"]))
	request.Header.Add("file-path", strings.ReplaceAll(path, "\\", "/"))
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("upload-done", finish)
	request.Header.Add("owner-name", user)

	client := &http.Client{
		// * 1 minute time out
		Timeout: time.Duration(60 * time.Second),
	}
	response, errorClient := client.Do(request)

	if errorClient != nil {
		colors.ErrorMessage(
			"making the request to send the files\n" + colors.Trace)
		log.Fatal(errorClient)
	}
	defer response.Body.Close()

	content, erroRead := ioutil.ReadAll(response.Body)

	var responseJSON map[string]interface{}

	json.Unmarshal(content, &responseJSON)

	if responseJSON["error"] == true {
		colors.ErrorMessage(fmt.Sprintf("%v", responseJSON["message"]) + colors.Trace)
		log.Fatal(responseJSON["message"])
	}

	if erroRead != nil {
		colors.ErrorMessage(
			"an error occurred reading the response from the server\n" + colors.Trace)
		log.Fatal(erroRead)
	}

	return content
}

// SendFiles test
func SendFiles(root string, user string, done bool) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			time.Sleep(900 * time.Millisecond)
			fmt.Printf(
				"%v |- uploading:%v %v %v %v \n", colors.Green,
				colors.Reset, colors.Yellow, path, colors.Reset)

			defer uploadFile(filepath.Join(handlers.Cwd(), path), path, user, done)
		}
		return nil
	})
}
