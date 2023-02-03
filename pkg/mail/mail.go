package mail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Mail struct {
	From     string
	To       []string
	Cc       []string
	Host     string
	Port     int
	Username string
	Password string
	E        *email.Email
}

func (m *Mail) SendHtml(subject string, body []byte, attachments ...string) (err error) {
	m.E.From = m.From
	m.E.To = m.To
	m.E.Cc = m.Cc
	m.E.Subject = subject
	m.E.HTML = body
	for _, attachment := range attachments {
		if _, err = m.E.AttachFile(attachment); err != nil {
			return
		}
	}
	err = m.E.Send(fmt.Sprintf("%s:%d", m.Host, m.Port),
		smtp.PlainAuth("", m.Username, m.Password, m.Host))
	return
}
