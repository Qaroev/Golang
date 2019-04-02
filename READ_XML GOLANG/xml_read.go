package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	xmlfile, err := os.Open("READ_XML GOLANG/users.xml")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(xmlfile)
	byteFile, _ := ioutil.ReadAll(xmlfile)
	fmt.Println(string(byteFile))
	var users Users
	xml.Unmarshal(byteFile, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

}
