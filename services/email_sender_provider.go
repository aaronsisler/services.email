package services

func NewDefaultEmailSender() (EmailSender, error) {
	// You could make this configurable later (e.g., env var to switch providers)
	return NewSESEmailSender("us-east-1")
}
