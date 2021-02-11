package responses

// Pagination :nodoc:
type Pagination struct {
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	TotalData   int         `json:"total_data"`
	Data        interface{} `json:"data,omitempty"`
}
