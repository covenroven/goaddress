package model

import (
    "database/sql"
)

type Address struct {
    Id int `json: "id"`
    Street string `json: "street"`
    City string `json: "city"`
    Province string `json: "province"`
    PostalCode string `db: "postal_code", json: "postal_code"`
    Country string `json: "country"`
}

type User struct {
    Id int `json: "id"`
    Name sql.NullString `json: "name"`
    Email string `json: "email"`
}
