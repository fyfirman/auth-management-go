package mail_server

import (
	"os"

	"github.com/resend/resend-go/v2"
)

type SendEmailRequest struct {
	From    string
	To      []string
	Subject string
	Html    string
}

type MailInterface interface {
	Send(request *SendEmailRequest)
}

type Mail struct {
}

func New() *Mail {
	return &Mail{}
}

func (m *Mail) Send(request *SendEmailRequest) (bool, error) {
	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)

	_, err := client.Emails.Send(&resend.SendEmailRequest{
		From:    request.From,
		To:      request.To,
		Subject: request.Subject,
		Html:    request.Html,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
