//FIlename:cmd/api/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// global variable to hold the application
// version number
const version = "1.0.0"

type config = struct {
	port int
	env  string// development, staging, production etc.
}

// setup dependency injection
type application struct {
	config config //variable name "config" of type config
	logger *log.Logger
}

func main() {
	var cfg config
	//read in the flags that are needed to populate our config
	//get the argumeents for the user for the server configuration
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()
	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//Creat an instance of our application struct
	//create an object if type application
	app := &application{
		config: cfg,
		logger: logger, //for terminal

	}
	//create our new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//create our HTTP server

	srv := &http.Server {
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout:30 * time.Second,
	}
	//start our server
	logger.Printf("Starting %s Server on port %d", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err) //if error then print out the error
}
