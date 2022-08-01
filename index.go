package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/64693710/parse-json-file-in-golang
type IoTservices_Data struct {
	IoTservices []IoTservices `json:"j_IoTservices"`
}

type IoTservices struct {
	Name    string `json:"Name"`
	Ip      string `json:"Ip"`
	Service string `json:"Service"`
	Power   string `json:"Power"`

	// Request Request `json:"request"`
}

// type Request struct {
// 	Method string `json:"method"`
// }

func main() {
	server := gin.Default()
	server.LoadHTMLFiles("index.html")

	filename := "iotservice.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}

	data := IoTservices_Data{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return
	}

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IoTservices": data.IoTservices,
		})
	})

	log.Fatal(server.Run(":8222"))
}
