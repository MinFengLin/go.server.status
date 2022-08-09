package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// https://stackoverflow.com/questions/64693710/parse-json-file-in-golang
// http://puremonkey2010.blogspot.com/2020/11/parsing-json-files-with-golang.html read

// Start -  json struct
type UpsInfo_slice struct {
	Battery struct {
		Charge  string `json:"charge"`
		Runtime string `json:"runtime"`
		Type    string `json:"type"`
		Voltage string `json:"voltage"`
	} `json:"battery"`
	Input   struct {
		Voltage string `json:"voltage"`
	} `json:"input"`
	Upsfullinfo Upsfullinfo_s `json:"ups"`
}

// UPS full info
// https://github.com/bemasher/JSONGen
type Upsfullinfo_s struct {
	Beeper struct {
		Status string `json:"status"`
	} `json:"beeper"`
	Load   string       `json:"load"`
	Mfr    string       `json:"mfr"`
	Model  string       `json:"model"`
	Serial string       `json:"serial"`
	Status string       `json:"status"`
	Test   struct {
		Result string `json:"result"`
	} `json:"test"`
}
// End -  json struct

func Parser_upsinfo() UpsInfo_slice {

	filename := "./local_api/json/go_ups_data.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
	}

	data := UpsInfo_slice{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	// fmt.Printf("%+v\n", data)

	return data
}
