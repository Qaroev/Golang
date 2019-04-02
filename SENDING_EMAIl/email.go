package main

import "github.com/mailgun/mailgun-go"

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun("tutorialedge.net", apiKey)
	m := mg.NewMessage(
		"gulboyqaroev@mail.ru",
		"Hello",
		"Testing some Mailgun!",
		"gulboyqaroev@mail.ru",
	)
	_, id, err := mg.Send(nil,m)
	return id, err
}


func main(){
	SendSimpleMessage("gulboyqaroev@mail.ru", "key-12345671234567")
}