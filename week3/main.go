package main

import (
	"learn/week3/app"
	"learn/week3/srv"
	"log"
)

func main() {

	if err := app.NewApp(srv.NewHttpServer("8080"), srv.NewRpcServer("8081")).Run();err != nil {
		log.Println(err)
	}

}
