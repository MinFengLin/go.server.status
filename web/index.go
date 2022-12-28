package web

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	apiservice "github.com/Minfenglin/go.server.status/pkg"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:embed assets/* templates/*
var f embed.FS

var (
	iotservice_data  apiservice.IoTservices_slice
	homeservice_data apiservice.Homeservices_slice
	upsinfo_data     apiservice.UpsInfo_slice
	diskinfo_data    apiservice.Disk_slice
	funcMap          = template.FuncMap{
		"percent_to_color_charge": apiservice.Percent_to_color_charge,
		"percent_to_color_disk":   apiservice.Percent_to_color_disk,
		"percent_to_byte_disk":    apiservice.Percent_to_byte_disk,
	}
)

func Update_data() {
	// iotservice_data data type =  apiservice.IoTservices_slice
	iotservice_data = apiservice.Parser_iotservice()
	// DEBUG_USE:fmt.Printf("%+v\n", iotservice_data)
	homeservice_data = apiservice.Parser_homeservices()
	upsinfo_data = apiservice.Parser_upsinfo()
	diskinfo_data = apiservice.Parser_disk_info()
	fmt.Printf("%+v\n", upsinfo_data)

	// var homeservice_data_swap

	for ii, _ := range homeservice_data.Homeservices {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 * -- Check_status (_, time_set, homeservice_data)
		 */
		go apiservice.Check_homeservice_status(ii, 500, &homeservice_data)
	}
	for ii, _ := range iotservice_data.IoTservices {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 * -- Check_status (_, time_set, iotservice_data)
		 */
		go apiservice.Check_iotservice_realtime_status(ii, 1000, &iotservice_data)
	}

	for ii, _ := range diskinfo_data.Disk_data {
		/* use Goroutine
		 * - if not use it will delay to long for wait service timeout or not
		 * -- Check_status (_, time_set, iotservice_data)
		 */
		go apiservice.Parser_disk_space(ii, &diskinfo_data)
	}
}

func Ticker_update(time_set int) {
	ticker := time.NewTicker(time.Duration(time_set) * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				Update_data()
			}
		}
	}()
}

func header(c *gin.Context) {
	c.HTML(http.StatusOK, "header.tmpl", gin.H{})
}

func qrcode(c *gin.Context) {
	c.HTML(http.StatusOK, "qrcode.tmpl", gin.H{
		// Qrcode
		"Qrcodeurl": apiservice.Parser_wifiqrcode(os.Getenv("SHOW"), os.Getenv("SSID"), os.Getenv("ENC"), os.Getenv("WIFIPWD")),
	})
}

func iot_service(c *gin.Context) {
	c.HTML(http.StatusOK, "iot_service.tmpl", gin.H{
		// IoTservice
		"IoTservices": iotservice_data.IoTservices,
	})
}

func home_service(c *gin.Context) {
	c.HTML(http.StatusOK, "home_service.tmpl", gin.H{
		// Homeservices
		"Homeservices": homeservice_data.Homeservices,
	})
}

func ups_info(c *gin.Context) {
	c.HTML(http.StatusOK, "ups_info.tmpl", gin.H{
		// UPS info
		"Upsinfo": upsinfo_data,
	})
}

func server_info(c *gin.Context) {
	c.HTML(http.StatusOK, "server_info.tmpl", gin.H{
		// OS info
		"ram_test":     apiservice.Parser_ram(),
		"average":      apiservice.Parser_load_average(),
		"uptime_users": apiservice.Parser_uptime_users(),
		"disk_info":    diskinfo_data.Disk_data,
		"top_ram":      apiservice.Parser_top_ram(),
		"top_cpu":      apiservice.Parser_top_cpu(),
	})
}

func Server_run() {
	server := gin.Default()

	templ := template.Must(template.New("").Funcs(funcMap).ParseFS(f, "templates/*.tmpl"))
	server.SetHTMLTemplate(templ)

	// example: ./web/assets/css/org.css
	server.StaticFS("/css", http.FS(f))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.GET("/", header, qrcode, iot_service, home_service, ups_info, server_info)
	log.Fatal(server.Run(":" + os.Getenv("PORT")))
}
