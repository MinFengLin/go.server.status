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
	Battery     Battery_s `json:"battery"`
	Input       Input_s   `json:"input"`
	Upsfullinfo Ups_s     `json:"ups"`
}
// Battery
type Battery_s struct {
	Charge  string `json:"charge"`
	Runtime string `json:"runtime"`
	Type    string `json:"type"`
	Voltage string `json:"voltage"`
}
// Input
type Input_s struct {
	Voltage string `json:"voltage"`
}
// UPS full info
type Ups_s struct {
	Beeper UPS_Beeper_s `json:"beeper"`
	Load   string       `json:"load"`
	Mfr    string       `json:"mfr"`
	Model  string       `json:"model"`
	Serial string       `json:"serial"`
	Status string       `json:"status"`
	Test   UPS_Test_r   `json:"test"`
}

type UPS_Beeper_s struct {
	Status string `json:"status"`
}

type UPS_Test_r struct {
	Result string `json:"result"`
}
// End -  json struct

// TODO: (down) not check if it correct

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
