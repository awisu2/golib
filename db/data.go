package db

import (
	_sql "database/sql"
	"strconv"
)

type RowData map[string]NullString

// Rowsから値を取得
func RowsToDatas(rows *_sql.Rows) (datas []RowData, err error) {
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
		data := RowData{}
		for i, name := range columns {
			data[name] = NullString(vals[i])
		}
		datas = append(datas, data)
	}
	return
}

// Rowsから値を取得
func RowsToMap(rows *_sql.Rows) (datas map[string]map[string]string, err error) {
	// Scan対象をしぼるためカラム情報の取得
	columns, _ := rows.Columns()
	count := len(columns)
	ptrs := make([]interface{}, count)

	datas = map[string]map[string]string{}
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
		data := map[string]string{}
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
func RowsToString(rows *_sql.Rows) (datas []map[string]string, err error) {
	// Scan対象をしぼるためカラム情報の取得
	columns, _ := rows.Columns()
	count := len(columns)
	ptrs := make([]interface{}, count)

	datas = []map[string]string{}
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
		data := map[string]string{}
		for i, name := range columns {
			data[name] = ConvertValToString(vals[i])
		}
		datas = append(datas, data)
	}
	return
}

func (self RowData) Get(key string) NullString {
	v, ok := self[key]
	if !ok {
		return NullString{}
	}
	return v
}

func (self RowData) Text(key string) string {
	v, ok := self[key]
	if !ok {
		return ""
	}
	return v.Text()
}

func (self RowData) Int(key string, def int) int {
	v, ok := self[key]
	if !ok {
		return def
	}

	i, err := strconv.Atoi(v.Text())
	if err != nil {
		return def
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
