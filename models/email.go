package models

type Header struct {
	Subject string `json:"subject" validate:"required"`
	From    string `json:"from" validate:"required,email"`
	To      string `json:"to" validate:"required,email"`
}

type Email struct {
	Header Header `json:"header" validate:"required"`
	Body   string `json:"body" validate:"required"`
}
