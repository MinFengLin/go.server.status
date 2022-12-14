package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func Percent_to_color_charge(charge_v string) string {
	int_p, _ := strconv.Atoi(charge_v)

	switch {
	case int_p >= 75:
		return "success"
	case int_p >= 60:
		return "info"
	case int_p >= 45:
		return "primary"
	case int_p >= 30:
		return "warning"
	default:
		return "danger"
	}
}

func Parser_upsinfo() UpsInfo_slice {

	filename := "./configs/go_ups_data.json"
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
