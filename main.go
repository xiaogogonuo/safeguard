package main

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/xiaogogonuo/safeguard/internal/config"
	mt "github.com/xiaogogonuo/safeguard/internal/model/template"
	v2 "gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"net/smtp"
	"time"
)

func main() {
	body, _ := ioutil.ReadFile("config/mail.yml")
	var conf config.Mail
	_ = v2.Unmarshal(body, &conf)
	mail := email.NewEmail()
	mail.From = conf.From
	mail.To = conf.To
	mail.Cc = conf.Cc

	t, _ := template.ParseFiles("static/html/template.html")
	buffer := new(bytes.Buffer)
	_ = t.Execute(buffer, mt.Template{
		From:    "rust最棒",
		To:      "go",
		Message: "rust is the best",
		Date:    time.Now().Format("20060102 15:04:05"),
	})
	mail.HTML = buffer.Bytes()
	_, _ = mail.AttachFile(".gitignore")
	mail.Send(fmt.Sprintf("%s:%d", conf.Service.Host, conf.Service.Port), smtp.PlainAuth("", conf.Auth.Username, conf.Auth.Password, conf.Service.Host))

}

func Send() (err error) {
	mail := email.NewEmail()
	mail.From = "陆佳伟<xiaogogonuo@163.com>"
	mail.To = []string{"xiaogogonuo@163.com", "1670187131@qq.com"}
	mail.Cc = []string{"1670187131@qq.com"}
	mail.Subject = "Test"

	t, _ := template.ParseFiles("template.html")
	body := new(bytes.Buffer)
	_ = t.Execute(body, struct {
		From    string
		To      string
		Message string
		Date    string
	}{
		From:    "go最棒",
		To:      "rust",
		Message: "go is the best",
		Date:    time.Now().Format("20060102 15:04:05"),
	})
	mail.HTML = body.Bytes()
	_, _ = mail.AttachFile(".gitignore")
	return mail.Send("smtp.163.com:25", smtp.PlainAuth("", "xiaogogonuo@163.com", "JDAOREDDCXYMAXXQ", "smtp.163.com"))
}
