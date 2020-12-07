package database

import (
    "database/sql"
    "strings"
    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
    db, err := sql.Open("postgres", getDsn())
    if err != nil {
        return nil, err
    }

    return db, nil
}

func getDsn() string {
    var s strings.Builder

    s.WriteString("postgres://")
    s.WriteString("postgres:pass@localhost/goaddress")

    return s.String()
}
