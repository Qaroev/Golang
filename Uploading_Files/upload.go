package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("File Upload EndPoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("Uploading_Files/tg_image_1463071270.jpeg")

	if err != nil {
		fmt.Println("Error Retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %v", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("Uploaded File: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-image", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "SuccessFully Upload file")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}
