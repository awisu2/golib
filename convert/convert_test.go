package convert

import (
	"testing"
)

func TestConvert(t *testing.T) {
	seed := "123456"
	expected := "１２３４５６"

	actual, err := JapaneseNumString(seed)
	if err != nil {
		t.Errorf("%v", err)
	}
	if expected != actual {
		t.Errorf(`not right convert "%v" want "%v"`, actual, expected)
	}
}
