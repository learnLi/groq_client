package groq

type Models struct {
	Object string `json:"object"`
	Data   []struct {
		Id            string `json:"id"`
		Object        string `json:"object"`
		Created       int    `json:"created"`
		OwnedBy       string `json:"owned_by"`
		Active        bool   `json:"active"`
		ContextWindow int    `json:"context_window"`
	} `json:"data"`
}
