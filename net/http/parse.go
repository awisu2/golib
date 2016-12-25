package http

import (
	"net/http"
	"net/url"
	"strings"
)

func Parse(w http.ResponseWriter, r *http.Request) *Session {
	pathes, queries := ConvertQueries(r)
	return &Session{w, r, pathes, queries}
}

// http.Requestを利用しやすく解析
func ConvertQueries(r *http.Request) (pathes Pathes, queries Queries) {

	// pathを分解して取得
	pathes = make([]string, 10) // 取得後にいちいちlenチェックをしたくないので十分な量を用意
	path := r.URL.Path[1:]      // ルート指定子を除く
	for i, v := range strings.Split(path, "/") {
		pathes[i] = v
	}

	// rawqueryを分解して取得
	rawquery := r.URL.RawQuery
	queries = make(map[string]string)
	if rawquery != "" {
		rawquery, _ = url.QueryUnescape(rawquery)
		raws := strings.Split(rawquery, "&")
		for _, v := range raws {
			i := strings.Index(v, "=")
			if i >= 0 && len(v) > i {
				queries[v[:i]] = v[i+1:]
			}
		}
	}

	// formからの情報をマージ
	err := r.ParseForm()
	if err == nil {
		for k, v := range r.Form {
			queries[k] = string(v[0])
			//TODO:select、radioで複数の値が飛んでくる場合は別途
		}
	}

	return
}
