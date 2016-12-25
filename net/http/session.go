package http

import (
	"net/http"
)

type Queries map[string]string
type Pathes []string

// *http.Requestから取得できる値をまとめた構造体
type Session struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Pathes  Pathes
	Queries Queries
}

// リダイレクト
func (self *Session) Redirect(urlStr string, code int) {
	http.Redirect(self.Writer, self.Request, urlStr, code)
}

// クッキー取得
func (self *Session) Cookie(name string) (*http.Cookie, error) {
	return self.Request.Cookie(name)
}

// 全クッキーの取得
func (self *Session) Cookies() []*http.Cookie {
	return self.Request.Cookies()
}

// クッキー登録
func (self *Session) SetCookie(cookie *http.Cookie) {
	http.SetCookie(self.Writer, cookie)
}

func (self *Session) SetCookies(cookie *http.Cookie) {
	http.SetCookie(self.Writer, cookie)
}
