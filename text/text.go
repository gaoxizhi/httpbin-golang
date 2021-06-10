// Package text
// Title   text.go
// Author  gaox·Eric
// Date    2021/6/10 21:41)
// Description  文本相应

package text

import (
	"encoding/json"
	"net/http"
)

// TextRequestHandler 文本相应
func TextRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("你的请求我已经收到"))
}

// HtmlRequestHandler 网页文本相应
func HtmlRequestHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<html>
	<head> 
		<title>测试响应内容为网页</title> <meta charset="utf-8"/>
	</head> 
	<body>
	我是以网页的形式响应过来的! 
	</body>
</html>
`
	w.Write([]byte(html))
}

type User struct {
	ID       int
	Username string
	Password string
}

// JsonRequestHandler 网页文本相应
func JsonRequestHandler(w http.ResponseWriter, r *http.Request) {
	//设置响应头中内容的类型
	w.Header().Set("Content-Type", "application/json")
	user := User{
		ID: 1, Username: "admin", Password: "123456",
	}
	//将 user 转换为 json 格式
	json, _ := json.Marshal(user)
	w.Write(json)
}
