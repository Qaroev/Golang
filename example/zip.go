package main

import (
	"net/http"
	"bytes"
	"archive/zip"
	"io/ioutil"
	"log"
)

func downloadZip(w http.ResponseWriter, r *http.Request) {

	// create an empty buffer and allocate some memory to it
	buf := new(bytes.Buffer)

	// create an archive with zip.NewWriter and use the empty buffer
	// so the zip write to it
	zipArchive := zip.NewWriter(buf)

	// read the file to be zipped, getting it's content in \[]byte
	file, err := ioutil.ReadFile("1.go")
	if err != nil {
		log.Printf("Error reading file from server: \n%v", err)
	}

	// create a file with whatever name you please in the zip archive
	archiveContent, err := zipArchive.Create("1.go")
	if err != nil {
		log.Printf("Error creating file in zip archive: \n%v", err)
	}

	// add the file's contents using the file variable gotten from the
	// reading of the file on line 9
	_, error := archiveContent.Write(file)
	if error != nil {
		log.Printf("Error writing data to file in zip archive: \n%v", err)
	}

	// close the zip archive
	errors := zipArchive.Close()
	if errors != nil {
		log.Printf("Error closing zip archive: \n%v", err)
	}

	// create an appropriate content-type, the last thing you want is
	// to somehow have an application/json type for a zip file
	w.Header().Add("content-type", "application/zip")
	// write the content of the buf variable created earlier which now has all
	// the content inside it
	w.Write(buf.Bytes())
}
func main()  {
	http.HandleFunc("/download", downloadZip)
	http.ListenAndServe(":8000", nil)
}

