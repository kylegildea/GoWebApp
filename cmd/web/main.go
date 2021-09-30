package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kylegildea/go-course/pkg/config"
	"github.com/kylegildea/go-course/pkg/handlers"
	"github.com/kylegildea/go-course/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	fmt.Printf("%v\n", http.ListenAndServe(portNumber, nil))
}
