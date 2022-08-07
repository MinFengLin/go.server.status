package main

import (
	"log"
	"net/http"
	"time"

	apiservice "service"

	"github.com/gin-gonic/gin"
)

var (
	iotservice_data apiservice.IoTservices_slice
	homeservice_data apiservice.Homeservices_slice
)

func update_data() {
	// iotservice_data data type =  apiservice.IoTservices_slice
	iotservice_data = apiservice.Parser_iotservice()
	// DEBUG_USE:fmt.Printf("%+v\n", iotservice_data)
	homeservice_data = apiservice.Parser_homeservices()

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
