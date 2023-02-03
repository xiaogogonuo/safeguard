package global

import (
	"github.com/xiaogogonuo/safeguard/internal/constant"
	"html/template"
)

var GTemplate *template.Template

func init() {
	t, err := template.ParseFiles(constant.TemplatePath)
	if err != nil {
		panic(err)
	}
	GTemplate = t
}
