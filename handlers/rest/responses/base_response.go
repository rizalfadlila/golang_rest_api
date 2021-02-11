package responses

// Response is base response, should be extended into more specific response
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}
