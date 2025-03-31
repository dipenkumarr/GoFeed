package main

import (
	"net/http"

	"github.com/dipenkumarr/GoFeed/internal/auth"
	"github.com/dipenkumarr/GoFeed/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlerwareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, "Error getting API key")
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
