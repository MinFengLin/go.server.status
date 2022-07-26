package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// // Book ...
// type Book struct {
// 	Title  string
// 	Author string
// }

// IoTservices
type IoTservices struct {
	// 開頭一定要大寫
	Ip      string
	Service string
	Power   string
}

func main() {
	server := gin.Default()
	server.LoadHTMLFiles("index.html")

	services := make([]IoTservices, 0)

	services = append(services, IoTservices{
		Ip:      "192.168.1.1",
		Service: "Author 1",
		Power:   "OFF",
	})

	services = append(services, IoTservices{
		Ip:      "192.168.1.2",
		Service: "Author 1",
		Power:   "ON",
	})

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IoTservices": services,
		})
	})
	// books := make([]Book, 0)
	// books = append(books, Book{
	// 	Title:  "Title 1",
	// 	Author: "Author 1",
	// })
	// books = append(books, Book{
	// 	Title:  "Title 2",
	// 	Author: "Author 2",
	// })

	// server.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"books": books,
	// 	})
	// })

	log.Fatal(server.Run(":8222"))

}
