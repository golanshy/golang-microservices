package repositories

type Oauth0ErrorResponse struct {
	StatusCode       int    `json:"status_code"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
