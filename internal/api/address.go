package api

import (
    "net/http"
    "log"
    // "database/sql"
	"github.com/go-chi/chi"
    "github.com/covenroven/goaddress/internal/database"
    "github.com/covenroven/goaddress/internal/model"
)

func IndexAddresses(w http.ResponseWriter, r *http.Request) {
    db, err := database.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var addresses []model.Model
    rows, err := db.Query("SELECT id, street, city, province, postal_code, country FROM addresses")
    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var address model.Address
        rows.Scan(
            &address.Id,
            &address.Street,
            &address.City,
            &address.Province,
            &address.PostalCode,
            &address.Country,
        )

        addresses = append(addresses, address)
    }

    response := model.Response{
        Status: 200,
        Message: "ok",
        Data: addresses,
    }

    responseWithJson(w, response)
}

func StoreAddress(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post user"))
}

func ShowAddress(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()

    userID := chi.URLParam(r, "userID")

    row := db.QueryRow("SELECT * FROM users WHERE id = $1", userID)

    var user model.User
    err := row.Scan(&user.Id, &user.Name, &user.Email)
    if err != nil {
        responseWithJson(w, model.Response{
            Status: 404,
            Message: "Not found",
            Data: []model.Model{},
        })
        return
    }

    responseWithJson(w, model.Response{
        Status: 200,
        Message: "ok",
        Data: []model.Model{user},
    })
}

