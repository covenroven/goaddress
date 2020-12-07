package model

type Model interface {}

type Response struct {
    Status int
    Message string
    Data []Model
}
