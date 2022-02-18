package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const geturl = "http://localhost:8000/get"

func getRequest(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//var responsestring strings.Builder
	content := string(dataBytes)
	fmt.Println(content)

	return nil
}

func main() {
	fmt.Println("Welcome to WebRequest Verbs")
	err := getRequest(geturl)
	if err != nil {
		panic(err)
	}

}
