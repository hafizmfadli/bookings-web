package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/hafizmfadli/bookings-web/internal/config"
	"github.com/hafizmfadli/bookings-web/internal/handlers"
	"github.com/hafizmfadli/bookings-web/internal/models"
	"github.com/hafizmfadli/bookings-web/internal/render"
)

const portNumber = ":8080"

var session *scs.SessionManager
var app config.AppConfig

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting application at %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// what am ia going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot crete template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	return nil
}
