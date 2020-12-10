package database

import (
    "database/sql"
    "strings"
    "os"
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
    s.WriteString(os.Getenv("DB_USERNAME"))
    s.WriteString(":")
    s.WriteString(os.Getenv("DB_PASSWORD"))
    s.WriteString("@")
    s.WriteString(os.Getenv("DB_HOST"))
    s.WriteString(":")
    s.WriteString(os.Getenv("DB_PORT"))
    s.WriteString("/")
    s.WriteString(os.Getenv("DB_DATABASE"))

    return s.String()
}
