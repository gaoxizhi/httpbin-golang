// Package htmlTemplate
// Title   htmlTemplate.go
// Author  gaox·Eric
// Date    2021/6/10 21:56)
// Description  html模版

package htmlTemplate

import (
	"html/template"
	"net/http"
)

// TemplateHandler 模版
func TemplateHandler(w http.ResponseWriter, r *http.Request) { //解析模板文件
	//解析模板文件
	t := template.Must(template.ParseFiles("resources/static/hello.html"))
	//执行模板
	t.Execute(w, "狸猫")
}
