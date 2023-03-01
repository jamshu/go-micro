package main

import (
	"fmt"
	"log"
	"net/http"
)

var webPort = "80" 

type Config struct {}
func main() {
	app := Config{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	fmt.Printf("running the broker service on the port %s\n", webPort)
	err := srv.ListenAndServe()
	
	if err != nil {
		log.Fatal(err)
	}

}