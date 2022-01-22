package api

import (
	"encoding/json"
	"net/http"
	myError "notifications/internal/error"
	"strconv"
)

type HttpResponse struct {
	Error   error       `json:"Error"`
	Message string      `json:"Message"`
	Result  interface{} `json:"Result"`
}

func NewHttpResponse(w http.ResponseWriter, error *myError.MyError, message string, result interface{}) {
	newResponse := HttpResponse{}
	if error != nil {
		newResponse.Error = error.Cause
		newResponse.Message = strconv.Itoa(error.StatusCode) + ": " + error.Message
	} else {
		newResponse.Message = message
	}
	newResponse.Result = result
	jsonNewResp, err := json.Marshal(newResponse)
	if err != nil {
		temp := myError.JsonMarshalError
		temp1, _ := json.Marshal(temp)
		w.WriteHeader(temp.StatusCode)
		w.Write(temp1)
		return
	}
	w.Write(jsonNewResp)
}
