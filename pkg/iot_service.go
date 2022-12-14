package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Check_iotservice_realtime_status(ii int, time_set int, iotservice_data *IoTservices_slice) {
	withtimeout := http.Client{Timeout: time.Duration(time_set) * time.Millisecond}
	url := "http://" + iotservice_data.IoTservices[ii].Ip + "/" + iotservice_data.IoTservices[ii].Other_cfg

	resp, err := withtimeout.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		if err = json.NewDecoder(resp.Body).Decode(&iotservice_data.IoTservices[ii]); err != nil {
			fmt.Println(err)
		}
	}
	// defer resp.Body.Close()

	// if err = json.NewDecoder(resp.Body).Decode(&iotservice_data.IoTservices[ii]); err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println("====================================")
	fmt.Println(iotservice_data.IoTservices[ii].Status)
	fmt.Println("====================================")
}

func Parser_iotservice() IoTservices_slice {

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

	data := IoTservices_slice{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	fmt.Printf("%+v\n", data)
	// Check_iotservice_realtime_status(&data)

	return data
}
