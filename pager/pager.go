package pager

import (
	"strconv"
)

type Pager struct {
	PerCount      int    // content count per page
	ItemNum       int    // display item num
	FormatPage    Format // default func(page string) string { return page }
	FormatCurrent Format // default is same FormatPage
	FormatBefore  Format
	FormatNext    Format
	FormatFirst   Format
	FormatLast    Format
	FormatWrap    FormatWrap
	Url           Format
}

// format page data
type Format func(page string) string
type FormatWrap func(pagerStr string) string

func NewInstance(perCount int, itemNum int) *Pager {
	if perCount < 1 {
		perCount = 1
	}
	if itemNum < 1 {
		itemNum = 1
	}
	return &Pager{
		PerCount:   perCount,
		ItemNum:    itemNum,
		FormatPage: func(page string) string { return page },
	}
}

// set format simply for bootstrap
func (self *Pager) SetFormatSimple(baseUrl string) *Pager {
	self.FormatPage = func(page string) string {
		return `<li><a href="` + baseUrl + page + `">` + page + `</a></li>`
	}
	self.FormatCurrent = func(page string) string {
		return `<li class="active"><a href="` + baseUrl + page + `">` + page + `</a></li>`
	}
	self.FormatBefore = func(page string) string {
		return `<li><a href="` + baseUrl + page + `">&lt;</a></li>`
	}
	self.FormatNext = func(page string) string {
		return `<li><a href="` + baseUrl + page + `">&gt;</a></li>`
	}
	self.FormatFirst = func(page string) string {
		return `<li><a href="` + baseUrl + page + `">&laquo;</a></li>`
	}
	self.FormatLast = func(page string) string {
		return `<li><a href="` + baseUrl + page + `">&raquo;</a></li>`
	}
	self.FormatWrap = func(pagerStr string) string { return `<ul class="pagination">` + pagerStr + `</ul>` }

	return self
}

func (self *Pager) Create(page int, allCount int) string {
	if allCount <= 0 {
		return ""
	}
	allPage, minPage, maxPage := self.getPageInfos(page, allCount)

	// 文字列にして取得
	str := ""
	if minPage > 2 && page <= allPage {
		if self.FormatFirst != nil {
			str += self.FormatFirst("1")
		}
	}

	if minPage > 1 && page <= allPage {
		if self.FormatBefore != nil {
			str += self.FormatBefore(strconv.Itoa(page - 1))
		}
	}

	for i := minPage; i <= maxPage; i++ {
		if i == page {
			if self.FormatCurrent == nil {
				str += self.FormatPage(strconv.Itoa(i))
			} else {
				str += self.FormatCurrent(strconv.Itoa(i))
			}
		} else {
			str += self.FormatPage(strconv.Itoa(i))
		}
	}

	if page > 0 && maxPage < allPage {
		if self.FormatNext != nil {
			str += self.FormatNext(strconv.Itoa(page + 1))
		}
	}

	if page > 0 && maxPage < allPage-1 {
		if self.FormatLast != nil {
			str += self.FormatLast(strconv.Itoa(allPage))
		}
	}

	if self.FormatWrap != nil {
		str = self.FormatWrap(str)
	}

	return str
}

// get page infos
func (self *Pager) getPageInfos(page int, allCount int) (allPage int, minPage int, maxPage int) {
	allPage = ((allCount - 1) / self.PerCount) + 1

	// ページの補正
	if page < 1 {
		page = 1
	} else if page > allPage {
		page = allPage
	}

	// 開始終了ページを取得
	minPage = page - ((self.ItemNum - 1) / 2)
	if minPage < 1 {
		minPage = 1
	}
	maxPage = minPage + self.ItemNum - 1
	if maxPage > allPage {
		maxPage = allPage
	}
	if maxPage-minPage < self.ItemNum-1 {
		minPage = maxPage - (self.ItemNum - 1)
		if minPage < 1 {
			minPage = 1
		}
	}

	return
}
