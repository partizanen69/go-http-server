package main

import (
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	email := strings.TrimPrefix(urlPath, "users/")
	if email == "" {
		err := errors.New("Expected user email in the url, but it was not provided")
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	err := apiCfg.dbClient.DeleteUser(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}