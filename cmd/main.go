package main

import (
	web "github.com/Minfenglin/go.server.status/web"
)

func main() {
	// load data at first
	web.Update_data()
	// set ticker to update data
	web.Ticker_update(10) // time(sec)

	web.Server_run()
}
