package mocks

import (
	"github.com/aaronsisler/services.email/models"
)

type MockEmailSender struct {
	CalledWith  models.Email
	ErrToReturn error
	WasCalled   bool
}

func (m *MockEmailSender) SendEmail(email models.Email) error {
	m.CalledWith = email
	m.WasCalled = true
	return m.ErrToReturn
}
