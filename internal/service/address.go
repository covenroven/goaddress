package service

import (
    "net/http"
    "strconv"
)

var AddressUrl = "http://localhost:3100"

func GetAddressOfUser (userID int) (http.Response, error) {
    res, err := http.Get(AddressUrl + "/addresses?user_id=" + strconv.Itoa(userID))
    if err != nil {
        return nil, err
    }

    return res, nil
}
