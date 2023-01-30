package model

import (
	"bytes"
	"html/template"
)

// Notification 邮件通知
type Notification struct {
	From string
	To   string
}

// Write 将结构体字段值填充到模版对应的位置并返回填充后整个模版文件的字节数据
func (n Notification) Write(t *template.Template) (body []byte, err error) {
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, n); err != nil {
		return
	}
	body = buffer.Bytes()
	return
}
