package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL.Path, r.Method)
		log.Println("---")
		defer func() {
			if err := recover(); err != nil {
				log.Println("recovered", err)
				resp, _ := json.Marshal(map[string]string{
					"message": "Panic",
				})
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
