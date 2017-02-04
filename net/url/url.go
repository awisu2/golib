package url

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

type Header struct {
	Name  string
	Value string
}

func Post(url string, headers []*Header, values url.Values) (responceBody []byte, err error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}

	for _, header := range headers {
		req.Header.Set(header.Name, header.Value)
	}

	client := &http.Client{}
	responce, err := client.Do(req)
	if err != nil {
		return
	}
	defer responce.Body.Close()

	responceBody, err = ioutil.ReadAll(responce.Body)
	if err != nil {
		return
	}

	return
}
