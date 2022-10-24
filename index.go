package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"os"
	"html/template"

	apiservice "service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	iotservice_data apiservice.IoTservices_slice
	homeservice_data apiservice.Homeservices_slice
	upsinfo_data apiservice.UpsInfo_slice
	diskinfo_data apiservice.Disk_slice
)

func update_data() {
	// iotservice_data data type =  apiservice.IoTservices_slice
	iotservice_data = apiservice.Parser_iotservice()
	// DEBUG_USE:fmt.Printf("%+v\n", iotservice_data)
	homeservice_data = apiservice.Parser_homeservices()
	upsinfo_data = apiservice.Parser_upsinfo()
	diskinfo_data = apiservice.Parser_disk_info()
	fmt.Printf("%+v\n", upsinfo_data)

	// var homeservice_data_swap

	for ii, _:= range homeservice_data.Homeservices {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 * -- Check_status (_, time_set, homeservice_data)
		 */
		go apiservice.Check_homeservice_status(ii, 500, &homeservice_data)
	}
	for ii, _:= range iotservice_data.IoTservices {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 * -- Check_status (_, time_set, iotservice_data)
		 */
		go apiservice.Check_iotservice_realtime_status(ii, 1000, &iotservice_data)
	}

	for ii, _:= range diskinfo_data.Disk_data {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 * -- Check_status (_, time_set, iotservice_data)
		 */
		go apiservice.Parser_disk_space(ii, &diskinfo_data)
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
	server.SetFuncMap(template.FuncMap{
		"percent_to_color_charge": apiservice.Percent_to_color_charge,
		"percent_to_color_disk": apiservice.Percent_to_color_disk,
		"percent_to_byte_disk": apiservice.Percent_to_byte_disk,
	})

	server.LoadHTMLFiles("index.html")
	server.Static("/Img", "./Img_dir")
	server.Static("/css", "./css_dir")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// Qrcode
			"Qrcodeurl": apiservice.Parser_wifiqrcode(os.Getenv("SHOW"), os.Getenv("SSID"), os.Getenv("ENC"), os.Getenv("WIFIPWD")),
			// IoTservice
			"IoTservices": iotservice_data.IoTservices,
			// Homeservices
			"Homeservices": homeservice_data.Homeservices,
			// UPS info
			"Upsinfo": upsinfo_data,
			// OS info
			"ram_test": apiservice.Parser_ram(),
			"average": apiservice.Parser_load_average(),
			"uptime_users": apiservice.Parser_uptime_users(),
			"disk_info": diskinfo_data.Disk_data,
			"top_ram": apiservice.Parser_top_ram(),
			"top_cpu": apiservice.Parser_top_cpu(),
		})
	})

	log.Fatal(server.Run(":" + os.Getenv("PORT")))
}

func main() {
	// load data at first
	update_data()
	// set ticker to update data
	ticker_update(10) // time(sec)

	server_run()
}
