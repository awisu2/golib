package analyse

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// 解析タイプ
const (
	PATTERN_TYPE_FIND = iota
	PATTERN_TYPE_VALUE
)

const (
	VALUE_PATTERN_TYPE_HTML = iota
	VALUE_PATTERN_TYPE_TEXT
	VALUE_PATTERN_TYPE_ATTR
)

type Pattern struct {
	Type        int // PATTERN_TYPE
	Pattern     string
	To          string
	From        string
	ValPatterns []ValPattern
}

type ValPattern struct {
	Pattern string
	To      string
	ValType int
	Attr    string
}

type val map[string]string

// 複数または単数のデータを持たせるためこういう構造体にする
type Data struct {
	Vals []val
}

func (self *Data) Val() val {
	if self.Vals == nil {
		self.Vals = []val{}
	}
	return self.Vals[0]
}

type Datas map[string]*Data

// 文字列からDocumentを取得する
func NewDocumentFromString(str string) (document *goquery.Document, err error) {
	reader := strings.NewReader(str)
	document, err = goquery.NewDocumentFromReader(reader)
	if err != nil {
		return
	}

	return
}

func TestAnalyse(document *goquery.Document, patterns []Pattern) (ss map[string]*goquery.Selection, datas Datas) {
	ss = map[string]*goquery.Selection{}
	datas = Datas{}
	for i, p := range patterns {
		switch p.Type {
		case PATTERN_TYPE_FIND:
			{
				if i == 0 || p.From == "" {
					ss[p.To] = document.Find(p.Pattern)
				} else {
					ss[p.To] = ss[p.From].Find(p.Pattern)
				}
			}
		case PATTERN_TYPE_VALUE:
			{
				if p.From == "" {
					continue
				}
				s, ok := ss[p.From]
				if !ok {
					continue
				}

				// TODO:複数じゃないパターンもあるよ
				data := &Data{}
				s.Each(func(i int, s *goquery.Selection) {
					d := map[string]string{}
					for _, p2 := range p.ValPatterns {
						v := ""
						switch p2.ValType {
						case VALUE_PATTERN_TYPE_HTML:
							v, _ = s.Find(p2.Pattern).Html()
						case VALUE_PATTERN_TYPE_TEXT:
							v = s.Find(p2.Pattern).Text()
						case VALUE_PATTERN_TYPE_ATTR:
							// TODO:Attr名が欲しい
							v, _ = s.Find(p2.Pattern).Attr(p2.Attr)
						}

						d[p2.To] = v
					}
					data.Vals = append(data.Vals, d)
				})
				datas[p.To] = data
			}
		}
	}

	return
}

// パターンを取得して解析
func AnalyseByPattern(document *goquery.Document, patterns []string) {
	selections := document.Find("body")

	selections.Each(func(i int, s *goquery.Selection) {
		v, err := s.Html()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(i, " : ", v)
	})

}

func Array(selection *goquery.Selection) (arr []*goquery.Selection) {
	arr = []*goquery.Selection{}

	selection.Each(func(i int, s *goquery.Selection) {
		arr = append(arr, s)
	})
	return
}

// jsonをanalyse用のpatternに変換する
func JsonToPatterns(s string) (patterns []Pattern, err error) {
	dec := json.NewDecoder(strings.NewReader(s))

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		return
	}

	// while the array contains values

	for dec.More() {
		var pattern Pattern
		err = dec.Decode(&pattern)
		if err != nil {
			return
		}
		patterns = append(patterns, pattern)
	}
	return
}
