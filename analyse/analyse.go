package analyse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// 解析タイプ
const (
	PATTERN_TYPE_FIND = iota
	PATTERN_TYPE_VALUE
)

type Pattern struct {
	Type int // PATTERN_TYPE
	BasePattern
	From     string
	Patterns []BasePattern
}

type BasePattern struct {
	Pattern string
	To      string
}

type Data struct {
	Val    string
	Childs []Datas
}

type Datas map[string]Data

// 文字列からDocumentを取得する
func NewDocumentFromString(str string) (document *goquery.Document, err error) {
	reader := strings.NewReader(str)
	document, err = goquery.NewDocumentFromReader(reader)
	if err != nil {
		return
	}

	return
}

func TestAnalyse(document *goquery.Document, patterns []string) (ss map[string]*goquery.Selection) {
	ps := []*Pattern{
		&Pattern{PATTERN_TYPE_FIND, BasePattern{"body", "body"}, "", nil},
		&Pattern{PATTERN_TYPE_FIND, BasePattern{".work_table_data", "items"}, "body", nil},
		&Pattern{PATTERN_TYPE_FIND, BasePattern{".work_name", "item"}, "", nil},
		&Pattern{PATTERN_TYPE_VALUE, BasePattern{"", "items"}, "items", []BasePattern{
			BasePattern{".work_name", "name"},
		}},
	}
	fmt.Println(ps)

	datas := Datas{}
	fmt.Println(datas)

	ss = map[string]*goquery.Selection{}
	for i, p := range ps {
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
				datass := []Datas{}
				s.Each(func(i int, s *goquery.Selection) {
					for _, p2 := range p.Patterns {
						ds := Datas{}
						// TODO取得方法をAttrとかその他諸々にも対応
						val, _ := s.Find(p2.Pattern).Html()
						ds[p2.To] = Data{Val: val}

						datass = append(datass, ds)
					}
				})
				datas[p.To] = Data{Childs: datass}
			}
		}
	}
	fmt.Println(datas["items"])

	//	selections := document.Find(".work_2col_table")
	//
	//	sss = append(sss, selections)
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
