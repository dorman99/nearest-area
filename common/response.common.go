package commoon

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
}

func FailedResponse(message string, err string, data interface{}) Response {
	sErr := strings.Split(err, "\n")
	ress := Response{
		Message: message,
		Error:   sErr,
		Data:    data,
	}
	return ress
}
