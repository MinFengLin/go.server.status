package main

import (
	"log"
	"net/http"

	iotservice "service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLFiles("index.html")
	server.Static("/Img_dir", "./Img")

	// iotservice_data data type =  iotservice.IoTservices_slice
	iotservice_data := iotservice.Parser_iotservice()
	// DEBUG_USE:fmt.Printf("%+v\n", iotservice_data)
	homeservice_data := iotservice.Parser_homeservices()

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IoTservices":  iotservice_data.IoTservices,
			"Homeservices": homeservice_data.Homeservices,
		})
	})

	log.Fatal(server.Run(":80"))
}
