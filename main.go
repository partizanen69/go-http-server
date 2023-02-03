package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/partizanen69/go-http-server/internal/database"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", testHandler)
	serveMux.HandleFunc("/err", testErrHandler)
	
	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Server has started on address:", addr)
	err := srv.ListenAndServe()
	log.Fatal("Could not start the server because of error:", err)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, database.User{
		Email: "some.email@gmail.com",
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "	application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	
	if payload == nil {
		w.WriteHeader(code)
		return
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		errJsonBytes, _ := json.Marshal(errorBody{
			Error: "Could not marshal resulting payload",
		})
		w.Write(errJsonBytes)
		return
	}

	w.WriteHeader(code)
	w.Write(payloadJson)
}

type errorBody struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	errBody := errorBody{
		Error: err.Error(),
	}
	respondWithJSON(w, code, errBody)
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("this is a test error"))
}