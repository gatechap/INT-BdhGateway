package apimodel

type BadRequestResponse struct {
	Uuid  string `json:"uuid" binding:"required"`
	Error string `json:"error" binding:"required"`
}

type HttpErrorResponse struct {
	Timestamp string `json:"@timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}
