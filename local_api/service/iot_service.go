package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"log"
)

// https://stackoverflow.com/questions/64693710/parse-json-file-in-golang

type IoTservices_slice struct {
	IoTservices []IoTservices `json:"j_IoTservices"`
}

type IoTservices struct {
	Ip        string `json:"Ip"`
	Service   string `json:"Service"`
	Other_cfg string `json:"Other_cfg"`
	Status    string `json:"POWER"`
}

func Check_iotservice_realtime_status(iotservice_data *IoTservices_slice) {
	for ii, _:= range iotservice_data.IoTservices {
		url := "http://"+ iotservice_data.IoTservices[ii].Ip + "/" + iotservice_data.IoTservices[ii].Other_cfg
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(&iotservice_data.IoTservices[ii]); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(iotservice_data.IoTservices[ii].Status)
	}
}

func Parser_iotservice() IoTservices_slice {

	filename := "./local_api/json/service_data.json"
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

	
	fmt.Printf("%+v\n", data)
	// Check_iotservice_realtime_status(&data)

	return data
}
