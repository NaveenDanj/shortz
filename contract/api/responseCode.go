package auth

import (
	"encoding/json"
	"net/http"
)

type ResponseCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var ResponseCodeString = map[int]ResponseCode{
	http.StatusOK:                  {Code: "OK", Message: "OK"},
	http.StatusUnauthorized:        {Code: "UNAUTHORIZED", Message: "UNAUTHORIZED"},
	http.StatusForbidden:           {Code: "FORBIDDEN", Message: "FORBIDDEN"},
	http.StatusNotFound:            {Code: "NOT_FOUND", Message: "NOT_FOUND"},
	http.StatusInternalServerError: {Code: "INTERNAL_SERVER_ERROR", Message: "INTERNAL_SERVER_ERROR"},
	http.StatusBadRequest:          {Code: "BAD_REQUEST", Message: "BAD_REQUEST"},
	http.StatusConflict:            {Code: "CONFLICT", Message: "CONFLICT"},
}

func ResponseWriter(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"code":         ResponseCodeString[statusCode].Code,
		"message_code": ResponseCodeString[statusCode].Message,
		"message":      message,
	})
}

func ResponseWithDataWriter(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":         ResponseCodeString[statusCode].Code,
		"message_code": ResponseCodeString[statusCode].Message,
		"data":         data,
	})
}
