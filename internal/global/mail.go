package global

import (
	"github.com/jordan-wright/email"
	"github.com/xiaogogonuo/safeguard/internal/config"
	"github.com/xiaogogonuo/safeguard/internal/constant"
	"github.com/xiaogogonuo/safeguard/pkg/mail"
	v2 "gopkg.in/yaml.v2"
	"io/ioutil"
)

var GMail *mail.Mail

func init() {
	body, err := ioutil.ReadFile(constant.MailConfigPath)
	if err != nil {
		panic(err)
	}
	var c config.Mail
	if err = v2.Unmarshal(body, &c); err != nil {
		panic(err)
	}
	GMail = &mail.Mail{
		From:     c.From,
		To:       c.To,
		Cc:       c.Cc,
		Host:     c.Host,
		Port:     c.Port,
		Username: c.Username,
		Password: c.Password,
		E:        email.NewEmail(),
	}
}
