package main

import (
	"log"
	"net/http"
	"time"
	"fmt"

	apiservice "service"

	"github.com/gin-gonic/gin"
)

var (
	iotservice_data apiservice.IoTservices_slice
	homeservice_data apiservice.Homeservices_slice
	upsinfo_data apiservice.UpsInfo_slice
)

func update_data() {
	// iotservice_data data type =  apiservice.IoTservices_slice
	iotservice_data = apiservice.Parser_iotservice()
	// DEBUG_USE:fmt.Printf("%+v\n", iotservice_data)
	homeservice_data = apiservice.Parser_homeservices()
	upsinfo_data = apiservice.Parser_upsinfo()
	fmt.Printf("%+v\n", upsinfo_data)

	// var homeservice_data_swap

	for ii, _:= range homeservice_data.Homeservices {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 * -- Check_status (_, time_set, homeservice_data)
		 */
		go apiservice.Check_status(ii, 500, &homeservice_data)
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

func server_run() {
	server := gin.Default()
	server.LoadHTMLFiles("index.html")
	server.Static("/Img", "./Img_dir")
	server.Static("/css", "./css_dir")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IoTservices":  iotservice_data.IoTservices,
			"Homeservices": homeservice_data.Homeservices,
			"Upsinfo": upsinfo_data,
		})
	})

	log.Fatal(server.Run(":80"))
}

func main() {
	// load data at first
	update_data()
	// set ticker to update data
	ticker_update(10) // time(sec)

	server_run()
}
