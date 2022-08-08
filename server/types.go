package server

import (
	"encoding/json"
	"net/http"
)

type requestJson struct {
	Message       string `json:"message"`
	To            string `json:"to"`
	From          string `json:"from"`
	TimeToLifeSec int    `json:"timeToLifeSec"`
}

type responseJson struct {
	Message string `json:"message"`
}

type accessTokenJSson struct {
	AccessToken string `json:"accessToken"`
}

type errorResponse struct {
	status int
	msg    string
}

func (er *errorResponse) Error() string {
	return er.msg
}

func returnErrorResponse(w http.ResponseWriter, er errorResponse) {
	resp := responseJson{Message: er.msg}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.status)
	json.NewEncoder(w).Encode(resp)
}

func returnOkResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
