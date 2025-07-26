package models

// type Signature struct {
// 	Name         string `json:"name"`
// 	EmailAddress string `json:"emailAddress"`
// 	PhoneNumber  string `json:"phoneNumber"`
// }

type Header struct {
	Subject string `json:"subject" validate:"required"`
	From    string `json:"from" validate:"required,email"`
	To      string `json:"to" validate:"required,email"`
}

type Email struct {
	Header Header `json:"header" validate:"required"`
	// Body      string    `json:"body"`
	// Signature Signature `json:"signature"`
}
