package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jritsema/gotoolbox"
	"github.com/jritsema/gotoolbox/web"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS

	//go:embed css/output.css
	css embed.FS

	//parsed templates
	html *template.Template
)

func main() {

	//exit process immediately upon sigterm
	handleSigTerms()

	//parse templates
	var err error
	html, err = web.TemplateParseFSRecursive(templateFS, ".html", true, nil)
	if err != nil {
		panic(err)
	}

	//add routes
	router := http.NewServeMux()
	router.Handle("GET /css/output.css", http.FileServer(http.FS(css)))

	//add
	router.Handle("GET /company/add", web.Action(addCompany))
	router.Handle("POST /company", web.Action(saveNewCompany))
	router.Handle("GET /company", web.Action(cancelSaveNewCompany))

	//edit
	router.Handle("GET /company/edit/{id}", web.Action(editCompany))
	router.Handle("PUT /company/{id}", web.Action(saveExistingCompany))
	router.Handle("GET /company/{id}", web.Action(cancelSaveExistingCompany))

	//delete
	router.Handle("DELETE /company/{id}", web.Action(deleteCompany))

	//home
	router.Handle("GET /", web.Action(index))
	router.Handle("GET /index.html", web.Action(index))

	//logging/tracing
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := tracing(nextRequestID)(logging(logger)(router))

	port := gotoolbox.GetEnvWithDefault("PORT", "8080")
	logger.Println("listening on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, middleware); err != nil {
		logger.Println("http.ListenAndServe():", err)
		os.Exit(1)
	}
}

func handleSigTerms() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
