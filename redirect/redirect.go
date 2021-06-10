// Package redirect
// Title   redirect.go
// Author  gaox·Eric
// Date    2021/6/10 21:48)
// Description  文件描述

package redirect

import "net/http"

// RedirectHandler 网页重定向

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	//以下操作必须要在 WriteHeader 之前进行
	w.Header().Set("Location", "https:www.baidu.com")
	w.WriteHeader(302)
}
