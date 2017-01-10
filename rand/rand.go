package rand

import (
	"encoding/base32"
	"math/rand"
	"time"
)

// パッケージが初回に読み込まれたときに実行される
func init() {
	rand.Seed(time.Now().UnixNano())
}

// ランダム文字列の取得
func String(length int) (s string) {
	b := make([]byte, length)
	rand.Read(b)
	s = base32.StdEncoding.EncodeToString(b)[:length]

	return
}
