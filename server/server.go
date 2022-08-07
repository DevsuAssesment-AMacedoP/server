package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type DevOpsHandler struct{}

func (devopsh DevOpsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		devOpsPost(w, r)

	default:
		http.Error(w, "ERROR", http.StatusNotFound)
	}
}

func devOpsPost(w http.ResponseWriter, r *http.Request) {
	var req requestJson

	err := decodeJsonBody(w, r, &req)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	resp := responseJson{Message: fmt.Sprintf("Hello %s your message will be sent", req.To)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
