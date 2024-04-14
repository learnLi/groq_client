package groq

type Account struct {
	SessionToken string `json:"session_token"`
	Organization string `json:"organization"`
}

func NewAccount(sessionToken string, organization string) *Account {
	return &Account{SessionToken: sessionToken, Organization: organization}
}

type Profile struct {
	User struct {
		Orgs struct {
			Object string `json:"object"`
			Data   []struct {
				Object             string `json:"object"`
				Id                 string `json:"id"`
				Created            int64  `json:"created"`
				Name               string `json:"name"`
				Description        string `json:"description"`
				Personal           bool   `json:"personal"`
				Priority           int    `json:"priority"`
				VerificationStatus string `json:"verification_status"`
				Settings           struct {
				} `json:"settings"`
				Role string `json:"role"`
			} `json:"data"`
		} `json:"orgs"`
	} `json:"user"`
}
