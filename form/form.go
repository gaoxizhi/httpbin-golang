// Package form
// Title   form.go
// Author  gaox·Eric
// Date    2021/6/10 21:31)
// Description  form表单内容

package form

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// RequestBodyHandler 获取请求体
// 使用net/http包创建一个函数处理器，函数的入参必须是http.ResponseWriter和http.Request指针类型
func RequestBodyHandler(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	//创建一个字节切片
	body := make([]byte, length)
	//读取请求体
	r.Body.Read(body)
	fmt.Fprintln(w, "请求体中的内容是:", string(body))
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))
	//解析表单
	r.ParseForm()
	//获取请求参数
	fmt.Fprintln(w, "请求参数为:", r.Form)
}

// RequestFormHandler 获取请求表单
func RequestFormHandler(w http.ResponseWriter, r *http.Request) {

	//解析表单
	r.ParseForm()
	//获取请求参数
	fmt.Fprintln(w, "请求参数为:", r.Form)
}

// RequestPostFormHandler 获取请求表单
func RequestPostFormHandler(w http.ResponseWriter, r *http.Request) {

	//解析表单
	r.ParseForm()
	form := r.PostForm
	for s := range form {
		fmt.Println(s, "=", form.Get(s))
	}
	//获取请求参数
	fmt.Fprintln(w, "请求参数为:", r.PostForm)
}

// FileFormHandler 获取文件表单
func FileFormHandler(w http.ResponseWriter, r *http.Request) {
	//解析表单
	r.ParseMultipartForm(1024)
	//打印表单数据
	fmt.Fprintln(w, r.MultipartForm)
	//解析表单
	r.ParseForm()
	form := r.PostForm
	for s := range form {
		fmt.Println(s, "=", form.Get(s))
	}
	//获取请求参数
	fmt.Fprintln(w, "请求参数为:", r.PostForm)
}

// FileFormHandler1 获取文件表单
func FileFormHandler1(w http.ResponseWriter, r *http.Request) {
	//解析表单
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["file"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}

// FileFormHandler2 获取文件表单
func FileFormHandler2(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}
