package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func Check_homeservice_status(ii int, time_set int, homeservice_data *Homeservices_slice) {
	withtimeout := net.Dialer{Timeout: time.Duration(time_set) * time.Millisecond}
	conn, err := withtimeout.Dial("tcp", homeservice_data.Homeservices[ii].Ip+":"+homeservice_data.Homeservices[ii].Port)

	if err != nil {
		homeservice_data.Homeservices[ii].Status = "Offline"
		fmt.Print("%s\n", homeservice_data.Homeservices[ii].Status)
	} else {
		homeservice_data.Homeservices[ii].Status = "Online"
		fmt.Print("%s\n", homeservice_data.Homeservices[ii].Status)
		defer conn.Close()
	}
}

func Parser_homeservices() Homeservices_slice {

	filename := "./configs/service_data.json"
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
