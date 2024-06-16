package helper

type response struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"code"`
	Data       interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) response {
	response := response{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}

	return response
}
