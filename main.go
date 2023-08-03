package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"time"
	"rest"
	"customlogger"
	"config"
	//"filestore"
)
//var LOGGER *log.Logger=nil

var DEFAULT_CONFIG_FILE_PATH="C:\\git\\DSOM-ADE\\fluent-heartbeat-service\\app.properties"
var err error
var CONFIG config.Config

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		//middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/agent", rest.Routes())
	})

	return router
}

func main() {
	log.Printf("Staring app...")
	
	logger:=customlogger.GetInstance()
	
	router := Routes()
	
	
	CONFIG,err:=config.LoadConfig(DEFAULT_CONFIG_FILE_PATH)
	log.Printf("3 %v",CONFIG)
	if err!=nil{
		logger.Fatal("Failed to load config ",err)
	}
	log.Printf(CONFIG.FileStoreName)
	logger.SetPrefix(time.Now().Format("2023-08-02 15:04:05") + " [main.go] ")
	logger.Println(CONFIG.FileStoreName)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		logger.Panicf("Logging err: %s\n", err.Error())
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}
	//log.Printf("listening at port 8080...")
	logger.Println("service listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router)) // Note, the port is usually gotten from the environment.
}