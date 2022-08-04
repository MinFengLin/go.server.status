package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// https://stackoverflow.com/questions/64693710/parse-json-file-in-golang

type Homeservices_slice struct {
	Homeservices []Homeservices `json:"j_Homeservices"`
}

type Homeservices struct {
	Ip        string `json:"Ip"`
	Service   string `json:"Service"`
	Port      string `json:"Port"`
	Other_cfg string `json:"Other_cfg"`
}

func Parser_homeservices() Homeservices_slice {

	filename := "./local_api/service_data.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
	}

	data := Homeservices_slice{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	// fmt.Printf("%+v\n", data)

	return data
}
