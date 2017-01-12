package sync

import (
	"sync"
)

func NewMutexer() *Mutexer {
	return &Mutexer{
		mu: sync.Mutex{},
	}
}

// mutex処理構造体
type Mutexer struct {
	mu sync.Mutex
}

// mutex.Lockをしながら関数をコールする
func (self *Mutexer) MutexFunc(f func()) {
	self.mu.Lock()
	defer self.mu.Unlock()
	f()
}

// mutexMap用データ
type MutexMapDatas map[string]interface{}

// Mutexerを利用したmap[string]interface{}構造体
type MutexMap struct {
	*Mutexer
	Datas MutexMapDatas
}

// 追加
func (self *MutexMap) Add(key string, data interface{}) {
	self.MutexFunc(func() {
		self.Datas[key] = data
	})
}

// 取得
func (self *MutexMap) Get(key string) (data interface{}, ok bool) {
	self.MutexFunc(func() {
		data, ok = self.Datas[key]
	})
	return
}

// 削除
func (self *MutexMap) Delete(key string) {
	self.MutexFunc(func() {
		_, ok := self.Datas[key]
		if ok {
			delete(self.Datas, key)
		}
	})
	return
}

// クリア
func (self *MutexMap) ClearAll() {
	self.MutexFunc(func() {
		self.Datas = MutexMapDatas{}
	})
	return
}

// キーの取得
func (self *MutexMap) Keys() (keys []string) {
	self.MutexFunc(func() {
		keys := make([]string, 0, len(self.Datas))
		for key, _ := range self.Datas {
			keys = append(keys, key)
		}
	})
	return keys
}

// データをコピーして取得
func (self *MutexMap) CopyDatas() (datas MutexMapDatas) {
	self.MutexFunc(func() {
		datas = make(map[string]interface{}, len(self.Datas))
		for key, data := range self.Datas {
			datas[key] = data
		}
	})
	return
}

// MutexMapの新インスタンス作成
func NewMutexMap() *MutexMap {
	return &MutexMap{
		Mutexer: NewMutexer(),
		Datas:   MutexMapDatas{},
	}
}
