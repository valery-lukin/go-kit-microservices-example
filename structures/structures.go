package structures

type GreetingRequest struct {
	Name string `json:"name"`
}

type GreetingResponse struct {
	Payload string `json:"payload"`
	Err     string `json:"err,omitempty"`
}

type IntRequest struct {
}

type IntResponse struct {
	Payload int `json:"payload"`
}
