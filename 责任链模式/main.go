package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	Handle(content string)
	next(handler Handler, content string)
}

// AdHandler 广告过滤
type AdHandler struct {
	handler Handler
}

func (ad *AdHandler) Handle(content string) {
	fmt.Println("执行广告过滤...")
	newContent := strings.Replace(content, "广告", "**", 1)
	fmt.Println(newContent)
}

func (ad *AdHandler) next(handler Handler, content string) {
	if ad.handler != nil {
		ad.handler.Handle(content)
	}
}

// YellowHandler 涉黄过滤
type YellowHandler struct {
	handler Handler
}

func (yellow *YellowHandler) Handle(content string) {
	fmt.Println("执行涉黄过滤...")
	newContent := strings.Replace(content, "涉黄", "**", 1)
	fmt.Println(newContent)
	yellow.next(yellow.handler, newContent)
}

func (yellow *YellowHandler) next(handler Handler, content string) {
	if yellow.handler != nil {
		yellow.handler.Handle(content)
	}
}

type SensitiveHandler struct {
	handler Handler
}

func (sensitive *SensitiveHandler) Handle(content string) {
	fmt.Println("执行敏感词过滤...")
	newContent := strings.Replace(content, "敏感词", "**", 1)
	fmt.Println(newContent)
	sensitive.next(sensitive.handler, newContent)
}

func (sensitive *SensitiveHandler) next(handler Handler, content string) {
	if sensitive.handler != nil {
		sensitive.handler.Handle(content)
	}
}

func main() {
	adHandler := &AdHandler{}
	yellowHandler := &YellowHandler{}
	sensitiveHandler := &SensitiveHandler{}
	// 将责任链串起来
	adHandler.handler = yellowHandler
	yellowHandler.handler = sensitiveHandler

	adHandler.Handle("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容")
}
