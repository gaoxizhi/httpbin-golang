// Package cookies
// Title   cookies.go
// Author  gaox·Eric
// Date    2021/6/10 22:19)
// Description  cookie测试

package cookies

import (
	"fmt"
	"net/http"
)

// CookieHandler cookie测试
func CookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{Name: "user1", Value: "admin", HttpOnly: true}
	cookie2 := http.Cookie{Name: "user2", Value: "superAdmin", HttpOnly: true, MaxAge: 60}
	//将 Cookie 发送给浏览器,即添加第一个 Cookie
	w.Header().Set("Set-Cookie", cookie1.String())
	//再添加一个 Cookie
	w.Header().Add("Set-Cookie", cookie2.String())
	//获取请求头中的 Cookie
	cookies := r.Header["Cookie"]
	fmt.Fprintln(w, cookies)
}
