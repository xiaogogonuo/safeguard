package mail

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Mail struct {
	From    string
	To      []string
	Cc      []string
	Service struct {
		Host string
		Port int
	}
	Auth struct {
		Username string
		Password string
	}
	E *email.Email
}

func (m *Mail) Send(subject string, html []byte) (err error) {
	m.E.Subject = subject
	m.E.HTML = html
	return m.E.Send("", smtp.PlainAuth("", "", "", ""))
}
