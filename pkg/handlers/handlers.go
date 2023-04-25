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

const Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const AlphLen = 63

type Handler struct {
	Repo Repo
}

func NewHandler(repo Repo) *Handler {
	return &Handler{Repo: repo}
}

type Repo interface {
	GetURL(hash string) (string, error)
	SaveHashByURL(url string, hash string) error
	Clear()
}

func (h *Handler) ShowURL(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]
	url, err := h.Repo.GetURL(hash)
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
	rand.Seed(time.Now().UTC().UnixNano())
	var hash string
	for {
		hash = CreateHash()
		err := h.Repo.SaveHashByURL(url, hash)
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			continue
		} else {
			log.Println(fmt.Sprintf("Failed to create hash. Error: %v", err))
			jsonError(w, err.Error(), http.StatusInternalServerError)
			break
		}
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

func CreateHash() string {
	hash := make([]byte, 10)
	for i := range hash {
		hash[i] = Alphabet[rand.Intn(AlphLen)]
	}
	return string(hash)
}
func jsonError(w http.ResponseWriter, msg string, status int) {
	resp, _ := json.Marshal(map[string]string{
		"message": msg,
	})
	w.WriteHeader(status)
	w.Write(resp)
}
