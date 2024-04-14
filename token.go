package groq

import "time"

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   time.Time `json:"expires_in"`
}

func NewToken(at string) *Token {
	return &Token{at, time.Now().Add(3 * time.Minute)}
}
