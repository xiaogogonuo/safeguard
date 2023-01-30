package main

import (
	"github.com/xiaogogonuo/safeguard/internal/global"
	"github.com/xiaogogonuo/safeguard/internal/model"
)

func main() {
	TestMail()
}

func TestMail() {
	body, _ := model.Notification{
		From: "max",
		To:   "hinton",
	}.Write(global.GTemplate)
	_ = global.GMail.SendHtml("Rust", body, "LICENSE", "README.md")
}
