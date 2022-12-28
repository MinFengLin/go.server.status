package pkg

/*
 * QR Code
 */

type Wifi_info_s struct {
	SHOW    string
	URL     string
	SSID    string
	WIFIPWD string
}

/*
 * IoTservices
 * https://stackoverflow.com/questions/64693710/parse-json-file-in-golang
 */

type IoTservices_slice struct {
	IoTservices []IoTservices `json:"j_IoTservices"`
}

type IoTservices struct {
	Ip        string `json:"Ip"`
	Service   string `json:"Service"`
	Other_cfg string `json:"Other_cfg"`
	Status    string `json:"POWER"`
}

/*
 * Homeservices
 * https://stackoverflow.com/questions/64693710/parse-json-file-in-golang
 */

type Homeservices_slice struct {
	Homeservices []Homeservices `json:"j_Homeservices"`
}

type Homeservices struct {
	Ip        string `json:"Ip"`
	Service   string `json:"Service"`
	Port      string `json:"Port"`
	Other_cfg string `json:"Other_cfg"`
	Status    string `json:"Status"`
}

/*
 * UPS
 * https://github.com/bemasher/JSONGen
 * https://stackoverflow.com/questions/64693710/parse-json-file-in-golang
 * http://puremonkey2010.blogspot.com/2020/11/parsing-json-files-with-golang.html read
 */

type UpsInfo_slice struct {
	Battery struct {
		Charge  string `json:"charge"`
		Runtime string `json:"runtime"`
		Type    string `json:"type"`
		Voltage string `json:"voltage"`
	} `json:"battery"`
	Input struct {
		Voltage string `json:"voltage"`
	} `json:"input"`
	Upsfullinfo Upsfullinfo_s `json:"ups"`
}

type Upsfullinfo_s struct {
	Beeper struct {
		Status string `json:"status"`
	} `json:"beeper"`
	Load   string `json:"load"`
	Mfr    string `json:"mfr"`
	Model  string `json:"model"`
	Serial string `json:"serial"`
	Status string `json:"status"`
	Test   struct {
		Result string `json:"result"`
	} `json:"test"`
}

/*
 * OS
 */

type Disk_slice struct {
	Disk_data []struct {
		Name         string `json:"Name"`
		Path         string `json:"Path"`
		Total_Space  uint64 `json:"Total_Space"`
		Free_Space   uint64 `json:"Free_Space"`
		Free_Space_P uint64 `json:"Free_Space_Precent"`
	} `json:"j_Disk"`
}

type Ram_s struct {
	Total_mem string
	Free_mem  string
	Mem_p     uint64
}
