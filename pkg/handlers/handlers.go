package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)



type Handler struct {
	Service Service
}

type Service interface {
	ShowLink(hash string) (string, error)
	SaveShortURL(url string) (string,error)
}
func NewHandler(service Service) *Handler {
	return &Handler{service}
}


func (h *Handler) ShowURL(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]
	url, err := h.Service.ShowLink(hash)
	if err == sql.ErrNoRows || err != nil && err.Error() == "no such hash in cache" {
		jsonError(w, "No such shorturl registered", http.StatusBadRequest)
	} else if err != nil {
		log.Println(fmt.Sprintf("Failed to ShowURL. Error: %v", err))
		jsonError(w, err.Error(), http.StatusInternalServerError)
	}
	resp, err := json.Marshal(map[string]string{
		"longurl":  url,
		"shorturl": hash,
	})
	if err != nil {
		log.Println(fmt.Sprintf("Failed to create response. Error: %v", err))
		jsonError(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(resp)
}

func (h *Handler) SaveURL(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)["url"]
	hash,err := h.Service.SaveShortURL(url)
	if err!=nil{
		jsonError(w,err.Error(),http.StatusInternalServerError)
	}
	resp, err := json.Marshal(map[string]string{
		"longurl":  url,
		"shorturl": hash,
	})

	if err != nil {
		log.Println(fmt.Sprintf("Failed to create response. Error: %v", err))
		jsonError(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(resp)
}

func jsonError(w http.ResponseWriter, msg string, status int) {
	resp, _ := json.Marshal(map[string]string{
		"message": msg,
	})
	w.WriteHeader(status)
	w.Write(resp)
}
