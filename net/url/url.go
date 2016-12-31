package url

import (
	"io/ioutil"
	"net/http"
)

// サイトから読み取り、それをテキストで表示する
func Get2String(url string) (str string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	byteArray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	response.Body.Close()
	str = string(byteArray)

	return
}
