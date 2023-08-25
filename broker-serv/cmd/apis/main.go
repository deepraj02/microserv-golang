package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker on port %s\n", webPort)

	//define the http server
	serv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//Start the server
	err := serv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
