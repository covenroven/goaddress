package api

import (
    "net/http"
    "log"
    "fmt"
    "database/sql"
    "encoding/json"
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

    querystr := r.URL.Query()
    fmt.Println(querystr.Get("userid"))

    var addresses []model.Model
    var rows *sql.Rows
    if querystr.Get("user_id") != "" {
        rows, err = db.Query(
            "SELECT id, street, city, province, postal_code, country, user_id FROM addresses WHERE user_id = $1",
            querystr.Get("user_id"),
        )
    } else {
        rows, err = db.Query("SELECT id, street, city, province, postal_code, country, user_id FROM addresses")
    }
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
            &address.UserId,
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
    db, _ := database.Connect()
    defer db.Close()

    if r.Body == nil {
        responseWithJson(w, model.Response{
            Status: 422,
            Message: "No parameter provided",
            Data: []model.Model{},
        })
        return;
    }

    var param model.Address
    err := json.NewDecoder(r.Body).Decode(&param)
    if err != nil {
        responseWithJson(w, model.Response{
            Status: 400,
            Message: err.Error(),
            Data: []model.Model{},
        })
        return;
    }

    var address model.Address
    err = db.QueryRow(
        `INSERT INTO addresses(street, city, province, postal_code, country, user_id)
        VALUES($1, $2, $3, $4, $5, $6)
        RETURNING id, street, city, province, postal_code, country, user_id`,
        param.Street,
        param.City,
        param.Province,
        param.PostalCode,
        param.Country,
        param.UserId,
    ).Scan(
        &address.Id,
        &address.Street,
        &address.City,
        &address.Province,
        &address.PostalCode,
        &address.Country,
        &address.UserId,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(address)
    responseWithJson(w, model.Response{
        Status: 201,
        Message: "Created",
        Data: []model.Model{address},
    })
}

type BatchStoreAddressParam struct {
    UserId int
    Addresses []model.Address
}

func BatchStoreAddresses(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()

    if r.Body == nil {
        responseWithJson(w, model.Response{
            Status: 422,
            Message: "No parameter provided",
            Data: []model.Model{},
        })
        return;
    }

    var param BatchStoreAddressParam
    err := json.NewDecoder(r.Body).Decode(&param)
    if err != nil {
        responseWithJson(w, model.Response{
            Status: 400,
            Message: err.Error(),
            Data: []model.Model{},
        })
        return;
    }

    stmt, _ := db.Prepare(
        `INSERT INTO addresses(street, city, province, postal_code, country, user_id)
        VALUES($1, $2, $3, $4, $5, $6)`,
    )
    defer stmt.Close()

    for _, address := range param.Addresses {
        if _, err := stmt.Exec(
            address.Street,
            address.City,
            address.Province,
            address.PostalCode,
            address.Country,
            param.UserId,
        ); err != nil {
            log.Fatal(err)
        }
    }

    responseWithJson(w, model.Response{
        Status: 201,
        Message: "Created",
        Data: []model.Model{},
    })
}

func ShowAddress(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()

    addressID := chi.URLParam(r, "addressID")

    row := db.QueryRow("SELECT id, street, city, province, postal_code, country, user_id FROM addresses WHERE id = $1", addressID)

    var address model.Address
    err := row.Scan(
        &address.Id,
        &address.Street,
        &address.City,
        &address.Province,
        &address.PostalCode,
        &address.Country,
        &address.UserId,
    )
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
        Data: []model.Model{address},
    })
}

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()

    if r.Body == nil {
        responseWithJson(w, model.Response{
            Status: 422,
            Message: "No parameter provided",
            Data: []model.Model{},
        })
    }

    addressID := chi.URLParam(r, "addressID")

    row := db.QueryRow("SELECT id FROM addresses WHERE id = $1", addressID)

    var aid int
    err := row.Scan(&aid)
    if err != nil {
        responseWithJson(w, model.Response{
            Status: 404,
            Message: "Not found",
            Data: []model.Model{},
        })
        return
    }

    var param model.Address
    err = json.NewDecoder(r.Body).Decode(&param)
    if err != nil {
        responseWithJson(w, model.Response{
            Status: 400,
            Message: err.Error(),
            Data: []model.Model{},
        })
    }

    _, err = db.Exec(
        "UPDATE addresses SET street=$1, city=$2, province=$3, postal_code=$4, country=$5, user_id=$6 WHERE id = $7",
        param.Street,
        param.City,
        param.Province,
        param.PostalCode,
        param.Country,
        param.UserId,
        aid,
    )
    if err != nil {
        log.Fatal(err)
    }
    param.Id = aid
    responseWithJson(w, model.Response{
        Status: 200,
        Message: "Updated",
        Data: []model.Model{param},
    })
}

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()

    if r.Body == nil {
        responseWithJson(w, model.Response{
            Status: 422,
            Message: "No parameter provided",
            Data: []model.Model{},
        })
    }

    addressID := chi.URLParam(r, "addressID")

    _, err := db.Exec("DELETE FROM addresses WHERE id = $1", addressID)

    if err != nil {
        log.Fatal(err)
        responseWithJson(w, model.Response{
            Status: 500,
            Message: "Something went wrong",
            Data: []model.Model{},
        })
        return
    }

    responseWithJson(w, model.Response{
        Status: 204,
        Message: "Ok",
    })
}
