package model

import (
	"bytes"
	"github.com/xiaogogonuo/safeguard/internal/static"
)

// Template 通知邮件的模版
type template struct {
	From    string
	To      string
	Message string
	Date    string
}

func (t *template) Write() (body []byte, err error) {
	buffer := new(bytes.Buffer)
	if err = static.Template.Execute(buffer, t); err != nil {
		return
	}
	body = buffer.Bytes()
	return
}

func New() *template {
	return &template{}
}
