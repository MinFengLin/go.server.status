package main

import (
	"log"
	"net/http"
	"fmt"
	"net"
	"time"

	iotservice "service"

	"github.com/gin-gonic/gin"
)

var (
	iotservice_data iotservice.IoTservices_slice
	homeservice_data iotservice.Homeservices_slice
)

func check_status(ii int) {
	withtimeout := net.Dialer{Timeout: 500*time.Millisecond}
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

func update_data() {
	// iotservice_data data type =  iotservice.IoTservices_slice
	iotservice_data = iotservice.Parser_iotservice()
	// DEBUG_USE:fmt.Printf("%+v\n", iotservice_data)
	homeservice_data = iotservice.Parser_homeservices()

	// var homeservice_data_swap

	for ii, _:= range homeservice_data.Homeservices {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 */
		go check_status(ii)
	}
}

func ticker_update(time_set int) {
	ticker := time.NewTicker(time.Duration(time_set) * time.Second)

	go func() {
		for {
			select {
		 		case <- ticker.C:
					update_data()
			}
		}
	}()
}

func main() {
	server := gin.Default()
	server.LoadHTMLFiles("index.html")
	server.Static("/Img", "./Img_dir")
	server.Static("/css", "./css_dir")
	// update at first
	update_data()
	ticker_update(10) // time(sec)

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IoTservices":  iotservice_data.IoTservices,
			"Homeservices": homeservice_data.Homeservices,
		})
	})

	log.Fatal(server.Run(":80"))
}
