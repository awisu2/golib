package http

import (
	"github.com/awisu2/golib/db"
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
	DBs     map[string]*db.DB
	Any     map[string]interface{}
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

func (self *Session) DBOpen(host string, database string) (_db *db.DB, err error) {
	key := host + "/" + database
	if self.DBs != nil {
		_db, ok := self.DBs[key]
		if ok {
			return _db, nil
		}
	}

	_db, err = db.Open(host, database)
	if err != nil {
		return nil, err
	}

	if self.DBs == nil {
		self.DBs = map[string]*db.DB{}
	}
	self.DBs[key] = _db

	return
}

func (self *Session) Clear() {
	self.DBClose()
}

func (self *Session) DBClose() {
	if self.DBs != nil {
		for _, v := range self.DBs {
			v.Close()
		}
	}
}

func (self *Session) SetAny(k string, v interface{}) {
	if self.Any == nil {
		self.Any = map[string]interface{}{}
	}
	self.Any[k] = v
}

func (self *Session) GetAny(k string) (v interface{}) {
	if self.Any == nil {
		return
	}
	v, _ = self.Any[k]
	return
}
