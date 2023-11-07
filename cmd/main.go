package main

import (
	"enshorter/pkg/handlers"
	"LinkEnshorter/pkg/middleware"
	"enshorter/pkg/repo/cache"
	"enshorter/pkg/repo/database"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"log"
	"net/http"
)


func main() {
	cron := cron.New()
	var handler *handlers.Handler
	if usedb {
		repo := database.NewRepo()
		generator:=generator.NewGenerator(Alphabet)
		service:=service.NewService(repo,generator)
		handler = handlers.NewHandler(service)

	} else {
		repo := cache.NewRepo()
		generator:=generator.NewGenerator(Alphabet)
		service:=service.NewService(repo,generator)
		handler = handlers.NewHandler(service)
	}
	cron.AddFunc("@daily", func() {
		handler.Repo.Clear()
	})
	cron.Start()
	r := mux.NewRouter()

	r.HandleFunc("/api/show/{hash}", handler.ShowURL).Methods(http.MethodGet)
	r.HandleFunc("/api/save/{url}", handler.SaveURL).Methods(http.MethodPost)

	mux := middleware.Panic(r)

	port := ":8080"
	log.Println("Запуск сервера на localhost" + port)
	http.ListenAndServe(port, mux)
}
