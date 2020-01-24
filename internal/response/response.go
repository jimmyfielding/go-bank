package response

type ErrorResponse struct {
	Message string `json:"message"`
}

type OKResponse struct {
	Message string `json:"message"`
}
