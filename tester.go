package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func mains() {

	url := "http://localhost:8080/videos"
	method := "GET"

	payload := strings.NewReader(`{
    "title": "Test You tube",
    "description": "1234",
    "url": "https://youtube.com"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
