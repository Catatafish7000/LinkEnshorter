package main

import (
	"enshorter/pkg/handlers"
	"enshorter/pkg/middleware"
	"enshorter/pkg/repo/cache"
	"enshorter/pkg/repo/database"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

const usedb = true

func main() {
	cron := cron.New()
	var handler *handlers.Handler
	if usedb {
		repo := database.NewRepo()
		handler = handlers.NewHandler(repo)

	} else {
		repo := cache.NewRepo()
		handler = handlers.NewHandler(repo)
	}
	cron.AddFunc("@every 1s", func() {
		handler.Repo.Clear()
	})
	cron.Start()
	r := mux.NewRouter()

	r.HandleFunc("/api/{hash}", handler.ShowURL).Methods(http.MethodGet)
	r.HandleFunc("/api/{url}", handler.SaveURL).Methods(http.MethodPost)

	mux := middleware.Panic(r)

	port := ":8080"
	log.Println("Запуск сервера на localhost" + port)
	http.ListenAndServe(port, mux)
}