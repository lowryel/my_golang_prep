package main

import (
	"fmt"

    "net/url"
    "net/http"
    "log"
    "os"
)


func SendSMS(to string) string{
	// Set up the message
	apiKey := os.Getenv("VONAGE_API_KEY")
	apiSecret := os.Getenv("VONAGE_API_SECRET")

	fmt.Println("Starting VONAGE API......")
    text, from := "Hello world", "E64 format" //E64 format phone number (+233544971528)
	msgData := url.Values{}
	msgData.Set("api_key", apiKey)
	msgData.Set("api_secret", apiSecret)
	msgData.Set("to", to) // Replace with the recipient's phone number
	msgData.Set("from", from)            // Replace with your Vonage virtual number
	msgData.Set("text", text)   // Message text

	// Send the SMS using the Vonage API
	resp, err := http.PostForm("https://rest.nexmo.com/sms/json", msgData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Status:", resp.Status)
    fmt.Println(text)
    return text
}

func main(){
	var to string
	fmt.Println("Please enter phone: ")
	fmt.Scanln(&to)
	SendSMS(to)
}