package url

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(url string) (bytes []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func GetHeader(url string) (header http.Header, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return response.Header, nil
}

// サイトから読み取り、それをテキストで表示する
func Get2String(url string) (str string, err error) {
	byteArray, err := Get(url)
	if err != nil {
		return
	}
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
