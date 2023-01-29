package static

import (
	"github.com/jordan-wright/email"
	"github.com/xiaogogonuo/safeguard/internal/config"
	"github.com/xiaogogonuo/safeguard/pkg/mail"
	v2 "gopkg.in/yaml.v2"
	"io/ioutil"
)

var Mail *mail.Mail

func init() {
	body, _ := ioutil.ReadFile("config/mail.yml")
	var c config.Mail
	if err := v2.Unmarshal(body, &c); err != nil {
		panic(err)
	}
	e := email.NewEmail()
	e.From = c.From
	e.To = c.To
	e.Cc = c.Cc
	Mail = &mail.Mail{E: e}
}
