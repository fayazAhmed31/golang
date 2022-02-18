package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name   string `json:"studentName"` //Alias . Also read about omitempty,"-"
	RollNo uint64
	Class  uint64
	place  string //Unable to marshal this field because it is unexported
}

func EncodeJson() {
	students := []student{
		{Name: "Fayaz", RollNo: 1, Class: 16, place: "Achampet"},
		{"Siraz", 2, 18, "Achampet"},
	}

	encodedjson, err := json.MarshalIndent(students, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", encodedjson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{"studentName": "Fayaz","RollNo": 1,"Class": 16}`,
	)

	isValid := json.Valid(jsonDataFromWeb)
	var stud student
	if isValid {
		err := json.Unmarshal(jsonDataFromWeb, &stud)
		if err != nil {
			panic("Unable to Unmarshal the json")
		}
		fmt.Printf("%#v\n", stud)
	} else {
		fmt.Println("Json is NOT Valid")
	}

	//To store json data from web as key value pair
	var webData map[string]interface{}
	err := json.Unmarshal(jsonDataFromWeb, &webData)
	if err != nil {
		panic(err)
	}
	for key, val := range webData {
		fmt.Printf("Key is %s and value is %v and type is %T\n", key, val, val)
	}
	//fmt.Printf("%#v\n", webData)

}

func main() {
	fmt.Println("Welcome to json tutorial")
	EncodeJson()
	DecodeJson()
}
