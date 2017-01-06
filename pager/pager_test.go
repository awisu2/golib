package pager

import (
	"testing"
)

func TestPageCount(t *testing.T) {
	// perCount, itemNum
	p := (&Pager{PerCount: 3, ItemNum: 4}).Init()

	// page, allCount
	actual := p.Create(1, 10)
	expected := "1234"
	if actual != expected {
		t.Errorf(`perCountでallCountが割り切れないときの返却値が間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(2, 9)
	expected = "123"
	if actual != expected {
		t.Errorf(`perCountでallCountが割り切れるときの返却値が間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(3, 22)
	expected = "2345"
	if actual != expected {
		t.Errorf(`pageが中間の時間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(7, 21)
	expected = "4567"
	if actual != expected {
		t.Errorf(`pageが一番最後の場合に間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(6, 21)
	expected = "4567"
	if actual != expected {
		t.Errorf(`pageがを含む場合に間違っています。 "%v" want "%v"`, actual, expected)
	}
}

//	PerCount      int    // content count per page
//	ItemNum       int    // display item num
//	FormatPage    string // page format. defalut "%v"
//	FormatCurrent string // current page format. same formatPage if it's blank
//	FormatBefore  string // before page format. not display if it's blank
//	FormatNext    string // next page format. not display if it's blank
//	FormatFirst   string // first page format. not display if it's blank
//	FormatLast    string // last page format. not display if it's blank
//	UrlFormat       string // set format args `baseurl, page`, if BaseUrl set any value.
func TestPagerWithFormat(t *testing.T) {
	// perCount, itemNum
	p := (&Pager{PerCount: 2, ItemNum: 5}).Init().SetFormatSimple("/page")

	// page, allCount
	actual := p.Create(1, 100)
	expected := `<div class="pager">1<a hfer="/page/2">2</a><a hfer="/page/3">3</a><a hfer="/page/4">4</a><a hfer="/page/5">5</a><a hfer="/page/2">&gt;</a><a hfer="/page/50">&raquo;</a></div>`
	if actual != expected {
		t.Errorf(`ページが先頭のとき返却値が間違っています "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(2, 100)
	expected = `<div class="pager"><a hfer="/page/1">1</a>2<a hfer="/page/3">3</a><a hfer="/page/4">4</a><a hfer="/page/5">5</a><a hfer="/page/3">&gt;</a><a hfer="/page/50">&raquo;</a></div>`
	if actual != expected {
		t.Errorf(`ページが２番目のとき返却値が間違っています "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(3, 100)
	expected = `<div class="pager"><a hfer="/page/1">1</a><a hfer="/page/2">2</a>3<a hfer="/page/4">4</a><a hfer="/page/5">5</a><a hfer="/page/4">&gt;</a><a hfer="/page/50">&raquo;</a></div>`
	if actual != expected {
		t.Errorf(`ページが３番目のとき返却値が間違っています "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(4, 100)
	expected = `<div class="pager"><a hfer="/page/3">&lt;</a><a hfer="/page/2">2</a><a hfer="/page/3">3</a>4<a hfer="/page/5">5</a><a hfer="/page/6">6</a><a hfer="/page/5">&gt;</a><a hfer="/page/50">&raquo;</a></div>`
	if actual != expected {
		t.Errorf(`ページが4番目のとき返却値が間違っています "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(5, 100)
	expected = `<div class="pager"><a hfer="/page/1">&laquo;</a><a hfer="/page/4">&lt;</a><a hfer="/page/3">3</a><a hfer="/page/4">4</a>5<a hfer="/page/6">6</a><a hfer="/page/7">7</a><a hfer="/page/6">&gt;</a><a hfer="/page/50">&raquo;</a></div>`
	if actual != expected {
		t.Errorf(`ページが5番目のとき返却値が間違っています "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(25, 100)
	expected = `<div class="pager"><a hfer="/page/1">&laquo;</a><a hfer="/page/24">&lt;</a><a hfer="/page/23">23</a><a hfer="/page/24">24</a>25<a hfer="/page/26">26</a><a hfer="/page/27">27</a><a hfer="/page/26">&gt;</a><a hfer="/page/50">&raquo;</a></div>`
	if actual != expected {
		t.Errorf(`ページが中間のとき返却値が間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(46, 100)
	expected = `<div class="pager"><a hfer="/page/1">&laquo;</a><a hfer="/page/45">&lt;</a><a hfer="/page/44">44</a><a hfer="/page/45">45</a>46<a hfer="/page/47">47</a><a hfer="/page/48">48</a><a hfer="/page/47">&gt;</a><a hfer="/page/50">&raquo;</a></div>`
	if actual != expected {
		t.Errorf(`ページが最後からよっつ前の時返却値が間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(47, 100)
	expected = `<div class="pager"><a hfer="/page/1">&laquo;</a><a hfer="/page/46">&lt;</a><a hfer="/page/45">45</a><a hfer="/page/46">46</a>47<a hfer="/page/48">48</a><a hfer="/page/49">49</a><a hfer="/page/48">&gt;</a></div>`
	if actual != expected {
		t.Errorf(`ページが最後からみっつ前の時返却値が間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(48, 100)
	expected = `<div class="pager"><a hfer="/page/1">&laquo;</a><a hfer="/page/47">&lt;</a><a hfer="/page/46">46</a><a hfer="/page/47">47</a>48<a hfer="/page/49">49</a><a hfer="/page/50">50</a></div>`
	if actual != expected {
		t.Errorf(`ページが最後からふたつ前の時返却値が間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(49, 100)
	expected = `<div class="pager"><a hfer="/page/1">&laquo;</a><a hfer="/page/48">&lt;</a><a hfer="/page/46">46</a><a hfer="/page/47">47</a><a hfer="/page/48">48</a>49<a hfer="/page/50">50</a></div>`
	if actual != expected {
		t.Errorf(`ページが最後からひとつ前の時返却値が間違っています。 "%v" want "%v"`, actual, expected)
	}

	actual = p.Create(50, 100)
	expected = `<div class="pager"><a hfer="/page/1">&laquo;</a><a hfer="/page/49">&lt;</a><a hfer="/page/46">46</a><a hfer="/page/47">47</a><a hfer="/page/48">48</a><a hfer="/page/49">49</a>50</div>`
	if actual != expected {
		t.Errorf(`ページが最後の時値が間違っています。 "%v" want "%v"`, actual, expected)
	}
}
