package analyse

import (
	"testing"
)

func TestNewDocumentFromString(t *testing.T) {
	actual, err := NewDocumentFromString("<html><body><div>1</div><div>2</div></body></html>")
	if actual == nil {
		t.Errorf("got %v want not nil", actual)
	}

	if err != nil {
		t.Errorf("got err %v", err)
	}
}
