package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/covenroven/goaddress/internal/api"
)

func Init() (chi.Router, error) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/addresses", func(r chi.Router) {
	    r.Get("/", api.IndexAddresses);
	    r.Post("/", api.StoreAddress);
	    r.Post("/batch", api.BatchStoreAddresses);
	    r.Get("/{addressID}", api.ShowAddress);
	    r.Put("/{addressID}", api.UpdateAddress);
	    r.Delete("/{addressID}", api.DeleteAddress);
	})

	return r, nil
}
