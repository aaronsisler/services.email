package services

func NewDefaultEmailSender() (EmailSender, error) {
	return NewSESEmailSender("us-east-1")
}
