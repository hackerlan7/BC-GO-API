package main

import (
	"fmt"
	"myproject/app"
)

func main() {
	// app.InitDb()
	app.InitLog()
	fmt.Println("server is running")
	app.Router()
}
