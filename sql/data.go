package sql

import (
	_sql "database/sql"
	"strconv"
)

type RowStrData map[string]_sql.NullString
type RowData map[string]string

// Rowsから値を取得
func RowsToDatas(rows *_sql.Rows) (datas []RowStrData, err error) {
	// Scan対象をしぼるためカラム情報の取得
	columns, _ := rows.Columns()
	count := len(columns)
	ptrs := make([]interface{}, count)

	for rows.Next() {
		vals := make([]_sql.NullString, count)
		for i, _ := range columns {
			ptrs[i] = &vals[i]
		}
		err := rows.Scan(ptrs...)
		if err != nil {
			return nil, err
		}

		// マップに登録しなおす
		data := RowStrData{}
		for i, name := range columns {
			data[name] = vals[i]
		}
		datas = append(datas, data)
	}
	return
}

// Rowsから値を取得
func RowsToMap(rows *_sql.Rows) (datas map[string]RowData, err error) {
	// Scan対象をしぼるためカラム情報の取得
	columns, _ := rows.Columns()
	count := len(columns)
	ptrs := make([]interface{}, count)

	datas = map[string]RowData{}
	for rows.Next() {
		vals := make([]interface{}, count)
		for i, _ := range columns {
			ptrs[i] = &vals[i]
		}
		err := rows.Scan(ptrs...)
		if err != nil {
			return nil, err
		}

		// マップに登録しなおす
		data := RowData{}
		id := ""
		for i, name := range columns {
			data[name] = ConvertValToString(vals[i])

			// TODO:id以外のキーに対応
			if name == "id" {
				id = data[name]
			}
		}
		datas[id] = data
	}
	return
}

// Rowsから値を取得
func RowsToString(rows *_sql.Rows) (datas []RowData, err error) {
	// Scan対象をしぼるためカラム情報の取得
	columns, _ := rows.Columns()
	count := len(columns)
	ptrs := make([]interface{}, count)

	datas = []RowData{}
	for rows.Next() {
		vals := make([]interface{}, count)
		for i, _ := range columns {
			ptrs[i] = &vals[i]
		}
		err := rows.Scan(ptrs...)
		if err != nil {
			return nil, err
		}

		// マップに登録しなおす
		data := RowData{}
		for i, name := range columns {
			data[name] = ConvertValToString(vals[i])
		}
		datas = append(datas, data)
	}
	return
}

// 数値で取得
func (self RowData) Int(key string) int {
	s, ok := self[key]
	if !ok {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// interfaceを文字列に変換
func ConvertValToString(val interface{}) string {
	v, ok := val.([]byte)
	if ok {
		str := string(v)
		return str
	}

	i, ok := val.(int64)
	if ok {
		return strconv.FormatInt(i, 10)
	}

	// 上記以外ではnilのはず
	return ""
}

// テンポラリデータとして利用する用のデータ
type TmpRowDatas struct {
	Datas []RowData
}

// データをクリア
func (self *TmpRowDatas) Clear() {
	self.Datas = nil
}

// データの存在チェック
func (self *TmpRowDatas) IsExists() bool {
	return self.Datas != nil
}
