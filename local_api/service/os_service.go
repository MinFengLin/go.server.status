package service

import (
	"os/exec"
	"os"
	"io/ioutil"
	"regexp"
	"fmt"
	"encoding/json"
	"syscall"
	"math"
)

type Disk_slice struct {
	Disk_data []struct {
		Name         string `json:"Name"`
		Path         string `json:"Path"`
		Total_Space  uint64 `json:"Total_Space"`
		Free_Space   uint64 `json:"Free_Space"`
		Free_Space_P uint64 `json:"Free_Space_Precent"`
	} `json:"j_Disk"`
}

func Parser_load_average() [3]string {
	uptime_cmd := exec.Command("uptime")
	uptime_cmd_r, _ := uptime_cmd.Output()

	var load_average_r [3]string
	// uptime_cmd_s, _ := strconv.ParseFloat(string(uptime_cmd_r), 10)
	fmt.Println("============================================\n")
	fmt.Printf("%s\n", uptime_cmd_r)
	reg, _ := regexp.Compile(`average:? ([0-9\.]+),[\s]+([0-9\.]+),[\s]+([0-9\.]+)`)
	for ii, avgs := range reg.FindStringSubmatch(string(uptime_cmd_r)) {
		if (ii != 0) { load_average_r[ii-1] = avgs }
	}
	fmt.Println("============================================\n")

	for i := 0; i < 3; i++ {
		fmt.Println(load_average_r[i])
	}

	return load_average_r
}

func Parser_uptime_users() [2]string {
	uptime_d := `uptime -p | sed 's/up //g'`
	users_d := "uptime | grep -ohe '[0-9.*] user[s,]'| sed 's/,//g'"

	uptime_r, _ := exec.Command("bash","-c",uptime_d).Output()
	users_r, _ := exec.Command("bash","-c",users_d).Output()

	var uptime_users [2]string

	uptime_users[0] = string(uptime_r)
	uptime_users[1] =string(users_r)

	return uptime_users
}

func Percent_to_byte_disk(disk_d uint64)(string) {
	fmt.Println("===============aaaabbb======================")
	fmt.Printf("%d\n", disk_d)
	fmt.Println("===============aaaabbb======================")

	byte_arr := [9]string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	disk_d_byte := math.Floor(math.Log(float64(disk_d)) / math.Log(1024))
	fmt.Printf("%f\n", disk_d_byte)


	disk_d_print := float64(disk_d) / math.Pow(1024, math.Floor(disk_d_byte))
	fmt.Printf("%f\n", math.Pow(1024, math.Floor(disk_d_byte)))
	fmt.Printf("%f\n", disk_d_print)

	// total_print := total_byte / math.pow(1024, math.floor(byte_arr))
	// aa123 := free_print + "<small>" + byte_arr[free_byte] + "</small>" + "/" + total_print + "<small>" + byte_arr[total_byte] + "</small>"

	// fmt.Println("===============aaaaaaa======================")
	// fmt.Printf("%s\n", aa123)
	// fmt.Println("===============aaaaaaa======================")

	disk_d_print_p := fmt.Sprintf("%.2f", disk_d_print) + byte_arr[int(disk_d_byte-1)]

	return disk_d_print_p
	// return free_print + "<small>" + byte_arr[free_byte] + "</small>" + "/" + total_print + "<small>" + byte_arr[total_byte] + "</small>"
}

func Percent_to_color_disk(disk_v uint64)(string) {
	int_p :=  disk_v

	switch {
	case int_p >= 75:
		return "danger"
	case int_p >= 60 :
		return "warning"
	case int_p >= 45:
		return "primary"
	case int_p >= 30:
		return "info"
	default:
		return "success"
	}
}

func Parser_disk_space(ii int, disk *Disk_slice) {
	if (disk.Disk_data[ii].Name == "local") {
		disk.Disk_data[ii].Path, _ = os.Getwd()
		fmt.Printf("%s -> path ->\n", disk.Disk_data[ii].Path)
	}

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(disk.Disk_data[ii].Path, &fs)
	if err != nil {
		fmt.Printf("failed to parser disk space: %v", err)
	}
	disk.Disk_data[ii].Free_Space = fs.Bfree * uint64(fs.Bsize)
	disk.Disk_data[ii].Total_Space = fs.Blocks * uint64(fs.Bsize)
	disk.Disk_data[ii].Free_Space_P = 100 - uint64(math.Round(float64(disk.Disk_data[ii].Free_Space)/float64(disk.Disk_data[ii].Total_Space*100) * 10000))
	fmt.Printf("%s -> Free_Space -> %d\n", disk.Disk_data[ii].Name, disk.Disk_data[ii].Free_Space)
	fmt.Printf("%s -> Total_Space -> %d\n", disk.Disk_data[ii].Name, disk.Disk_data[ii].Total_Space)
	fmt.Printf("%s -> Free_Space_P -> %d\n", disk.Disk_data[ii].Name, disk.Disk_data[ii].Free_Space_P)

}

func Parser_disk_info() Disk_slice {
	filename := "./local_api/json/disk_data.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
	}

	data := Disk_slice{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	fmt.Printf("%+v\n", data)

	return data
}