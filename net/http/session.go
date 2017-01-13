package http

import (
	"fmt"
	"github.com/awisu2/golib/db"
	"net/http"
	"strings"
)

// *http.Requestから取得できる値をまとめた構造体
type Session struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	Pathes    []string
	Queries   map[string]string
	Querieses map[string][]string
	DBs       map[string]*db.DB
	Any       map[string]interface{}
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

func (self *Session) GetArrayByQueries(key string) []string {
	array, ok := self.Querieses[key]
	if ok {
		return array
	}

	v, ok := self.Queries[key]
	if ok {
		return []string{v}
	}
	return []string{}
}

// デバイスタイプ型
type DeviceType int

func (self DeviceType) Int() int {
	return int(self)
}

// デバイスタイプ
const (
	DEVICE_TYPE_PC DeviceType = iota + 1
	DEVICE_TYPE_MOBILE
	DEVICE_TYPE_TABLET
)

// デバイスタイプの取得
func (self *Session) GetDeviceType() DeviceType {
	return GetDeviceType(self.Request)
}

// デバイスタイプの取得
func GetDeviceType(r *http.Request) DeviceType {
	ua := strings.ToLower(r.UserAgent())
	deviceType := DEVICE_TYPE_PC
	if strings.Index(ua, "iphone") >= 0 ||
		strings.Index(ua, "ipod") >= 0 ||
		(strings.Index(ua, "android") >= 0 && strings.Index(ua, "mobile") >= 0) ||
		(strings.Index(ua, "windows") >= 0 && strings.Index(ua, "phone") >= 0) ||
		(strings.Index(ua, "firefox") >= 0 && strings.Index(ua, "mobile") >= 0) ||
		strings.Index(ua, "blackberry") >= 0 ||
		strings.Index(ua, "bb") >= 0 {
		deviceType = DEVICE_TYPE_MOBILE
	} else if strings.Index(ua, "ipad") >= 0 ||
		(strings.Index(ua, "windows") >= 0 && strings.Index(ua, "touch") >= 0 && strings.Index(ua, "tablet pc") >= 0) ||
		(strings.Index(ua, "android") >= 0 && strings.Index(ua, "mobile") >= 0) ||
		(strings.Index(ua, "firefox") >= 0 && strings.Index(ua, "tablet") >= 0) ||
		(strings.Index(ua, "kindle") >= 0 && strings.Index(ua, "silk") >= 0) ||
		strings.Index(ua, "playbook") >= 0 {
		deviceType = DEVICE_TYPE_TABLET
	}
	fmt.Println(ua)
	return deviceType
}
