// Package httpbin
// Title   main.go
// Author  gaox·Eric
// Date    2021/6/9 23:32)
// Description  golang 网络请求解析

package main

import (
	"httpbin/cookies"
	"httpbin/form"
	"httpbin/htmlTemplate"
	"httpbin/redirect"
	"httpbin/text"
	"net/http"
	"time"
)

func main() {

	//创建多路复用器
	mux := http.NewServeMux()

	// 在main函数中调用http中的HandleFunc函数指定处理指定请求的处理器
	mux.HandleFunc("/requestBody", form.RequestBodyHandler)
	mux.HandleFunc("/requestForm", form.RequestFormHandler)
	mux.HandleFunc("/requestPostForm", form.RequestPostFormHandler)
	mux.HandleFunc("/fileForm", form.FileFormHandler)
	mux.HandleFunc("/fileForm1", form.FileFormHandler1)
	mux.HandleFunc("/fileForm2", form.FileFormHandler2)

	mux.HandleFunc("/text", text.TextRequestHandler)
	mux.HandleFunc("/html", text.HtmlRequestHandler)
	mux.HandleFunc("/json", text.JsonRequestHandler)

	mux.HandleFunc("/redirect", redirect.RedirectHandler)

	mux.HandleFunc("/template", htmlTemplate.TemplateHandler)

	mux.HandleFunc("/cookie", cookies.CookieHandler)

	//创建路由
	//http.ListenAndServe(":8080", mux)
	//创建Server结构，并详细配置里面的字段
	server := http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: 2 * time.Second,
	}

	server.ListenAndServe()

}
