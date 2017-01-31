package convert

import (
	"bufio"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"strconv"
	"strings"
)

var NumString = [...]string{"０", "１", "２", "３", "４", "５", "６", "７", "８", "９"}

// 文字列の数字を日本語に変換
func JapaneseNumString(s string) (num string, err error) {
	for i := 0; i < len(s); i++ {
		idx, _ := strconv.Atoi(s[i : i+1])
		if err != nil {
			return "", nil
		}
		num += NumString[idx]
	}
	return
}

// convert code shiftjis to utf8
func ShiftJisToUtf8(str string) (utf8 string) {
	reader := strings.NewReader(str)
	tReader := transform.NewReader(reader, japanese.ShiftJIS.NewDecoder())
	scanner := bufio.NewScanner(tReader)
	for scanner.Scan() {
		utf8 = scanner.Text()
		break
	}
	if err := scanner.Err(); err != nil {
		return ""
	}

	return
}
