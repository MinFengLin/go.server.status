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

	// iotservice_data data type =  iotservice.IoTservices_slice
	iotservice_data := iotservice.Parser_iotservice()
	// DEBUG_USE:fmt.Printf("%+v\n", iotservice_data)

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IoTservices": iotservice_data.IoTservices,
		})
	})

	log.Fatal(server.Run(":8222"))
}
