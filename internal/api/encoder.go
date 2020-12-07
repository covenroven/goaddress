package api

import (
    "net/http"
    "encoding/json"
    "github.com/covenroven/goaddress/internal/model"
)

func responseWithJson(w http.ResponseWriter, r model.Response) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(r.Status)

    return json.NewEncoder(w).Encode(r)
}
