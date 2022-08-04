package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// https://stackoverflow.com/questions/64693710/parse-json-file-in-golang

type IoTservices_slice struct {
	IoTservices []IoTservices `json:"j_IoTservices"`
}

type IoTservices struct {
	Name    string `json:"Name"`
	Ip      string `json:"Ip"`
	Service string `json:"Service"`
	Power   string `json:"Power"`

	// Request Request `json:"request"`
}

// type Request struct {
// 	Method string `json:"method"`
// }

func Parser_iotservice() IoTservices_slice {

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

	data := IoTservices_slice{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	// fmt.Printf("%+v\n", data)

	return data
}
