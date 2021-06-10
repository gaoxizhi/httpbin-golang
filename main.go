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
)

func main() {
	http.HandleFunc("/requestBody", form.RequestBodyHandler)
	http.HandleFunc("/requestForm", form.RequestFormHandler)
	http.HandleFunc("/requestPostForm", form.RequestPostFormHandler)
	http.HandleFunc("/fileForm", form.FileFormHandler)
	http.HandleFunc("/fileForm1", form.FileFormHandler1)
	http.HandleFunc("/fileForm2", form.FileFormHandler2)

	http.HandleFunc("/text", text.TextRequestHandler)
	http.HandleFunc("/html", text.HtmlRequestHandler)
	http.HandleFunc("/json", text.JsonRequestHandler)

	http.HandleFunc("/redirect", redirect.RedirectHandler)

	http.HandleFunc("/template", htmlTemplate.TemplateHandler)

	http.HandleFunc("/cookie", cookies.CookieHandler)

	http.ListenAndServe(":8080", nil)
}
