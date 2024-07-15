package groq

type Account struct {
	SessionToken string `json:"session_token"`
	Organization string `json:"organization"`
	IsAPIKey     bool   `json:"is_api_key"`
}

func NewAccount(sessionToken string, organization string) *Account {
	return &Account{SessionToken: sessionToken, Organization: organization, IsAPIKey: false}
}

func NewAccountWithAPIKey(sessionToken string, organization string, isAPIKey bool) *Account {
	return &Account{SessionToken: sessionToken, Organization: organization, IsAPIKey: isAPIKey}
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
