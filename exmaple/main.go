package main

import (
	"encoding/json"
	groq "github.com/learnLi/groq_client"
	"log"
	"net/http"
	"time"
)

var (
	tokens map[string]*groq.Token
)

func init() {
	tokens = make(map[string]*groq.Token)
}

func main() {

	sessionToken := "<you-session-token>"
	organization := "<you-organization>"
	account := groq.NewAccount(sessionToken, organization)
	flog := log.Default()

	http.HandleFunc("/v1/chat/completions", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodOptions {
			// Set headers for CORS
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.Header().Set("content-type", "application/json")
			w.Write([]byte(`{"message": "pong"}`))
			return
		}
		var api_request groq.APIRequest

		err := json.NewDecoder(r.Body).Decode(&api_request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		api_request.Model = "mixtral-8x7b-32768"
		client := groq.NewBasicClient()
		Token := tokens[account.Organization]
		if Token == nil || Token.ExpiresIn.Before(time.Now()) {
			authenticateResponse, err := groq.GetSessionToken(client, account.SessionToken, "")
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			Token = groq.NewToken(authenticateResponse.Data.SessionJwt)
			tokens[account.Organization] = Token
		}
		response, err := groq.ChatCompletions(client, api_request, Token.AccessToken, account.Organization, "")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		flog.Println(response.StatusCode, r.RemoteAddr, r.RequestURI)
		defer response.Body.Close()
		groq.NewReadWriter(w, response).StreamHandler()
	})
	_ = http.ListenAndServe(":8080", nil)
	flog.Println("Listening on :8080")
}
