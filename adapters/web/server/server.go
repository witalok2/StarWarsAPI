package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/witalok2/starwarsplanet/adapters/web/handler"
	"github.com/witalok2/starwarsplanet/application"
)

type Webserver struct {
	Service application.PlanetServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {
	router := mux.NewRouter()
	middleware := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakePlanetHandlers(router, middleware, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
