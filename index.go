package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"os"
	"strconv"
	"html/template"

	apiservice "service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func percent_to_color_charge(charge_v string)(string) {
	int_p, _ :=  strconv.Atoi(charge_v)

	switch {
	case int_p >= 75:
		return "success"
	case int_p >= 60 :
		return "info"
	case int_p >= 45:
		return "primary"
	case int_p >= 30:
		return "warning"
	default:
		return "danger"
	}
}

func server_run() {
	server := gin.Default()
	server.SetFuncMap(template.FuncMap{
		"percent_to_color_charge": percent_to_color_charge,
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
			"Qrcodeurl": apiservice.Parser_wifiqrcode(os.Getenv("SSID"), os.Getenv("ENC"), os.Getenv("WIFIPWD")),
			"IoTservices": iotservice_data.IoTservices,
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
