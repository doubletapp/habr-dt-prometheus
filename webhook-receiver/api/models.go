package api

type DeafultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
