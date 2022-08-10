package service

import (
	"os/exec"
	"regexp"
	"fmt"
)

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
	uptime_d := "uptime | awk {'print $3'}"
	users_d := "uptime | awk {'print $4 \" \" $5'} | sed 's/,//g'"

	uptime_r, _ := exec.Command("bash","-c",uptime_d).Output()
	users_r, _ := exec.Command("bash","-c",users_d).Output()

	var uptime_users [2]string

	uptime_users[0] = string(uptime_r)
	uptime_users[1] =string(users_r)

	return uptime_users
}
