package main

import (
	"learn/week3/app"
	"learn/week3/srv"
	"log"
)

func main() {

	app := app.NewApp()
	app.AddServer(
		srv.NewHttpServer("8080"),
		srv.NewRpcServer("8081"),
	)
	if err := app.Run(); err != nil {
		log.Println(err)
	}

}
