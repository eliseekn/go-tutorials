package handlers

import (
	"encoding/json"
	"http-server/types"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	res := &types.Response{
		Status:  "success",
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
