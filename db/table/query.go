package table

import (
	"strconv"
)

//TODO:Check関数にする
//TODO:NullString
//func (self *Info) QueryInsertWithValues(vals []Val) (query string, args []interface{}, err error) {
//	query = self.QueryInsert()
//
//	query_val := ""
//	for _, field := range self.Fields {
//		if field.Name == "id" ||
//			field.AutoTime == AUTO_TIME_DELETE {
//			continue
//		}
//		query_val += ",?"
//	}
//	query_val = "(" + query_val[1:] + ")"
//
//	now := _sql.Now()
//	query_vals := ""
//	for _, val := range vals {
//		for _, field := range self.Fields {
//			if field.Name == "id" || field.AutoTime == AUTO_TIME_DELETE {
//				continue
//			}
//
//			arg := ""
//			if field.AutoTime == AUTO_TIME_CREATE || field.AutoTime == AUTO_TIME_UPDATE {
//				arg = now
//			} else {
//				a, ok := val[field.Name]
//				if !ok {
//					a = ""
//				}
//				arg = a
//			}
//
//			// 基本チェック
//			//TODO:空文字とNULLは違う
//			//			if field.IsNull {}
//
//			// 型チェック
//			switch field.Type {
//			case TYPE_INT, TYPE_MIDDLEINT:
//				if arg != "" {
//					_, err = strconv.Atoi(arg)
//					if err != nil {
//						err = fmt.Errorf("%vが数値ではありません", field.Name)
//						return
//					}
//				}
//			}
//
//			// サイズが大きすぎます
//			if field.Length > 0 {
//				if len(arg) > field.Length {
//					err = fmt.Errorf("%vのサイズが大きすぎます。 size:%v, limit:%v", field.Name, len(arg), field.Length)
//					return
//				}
//			}
//
//			// パラメータチェック
//			if field.Checker != nil {
//				err = field.Checker(arg)
//				if err != nil {
//					return
//				}
//			}
//
//			// 値を追加
//			args = append(args, arg)
//		}
//		query_vals += "," + query_val
//	}
//	query += query_vals[1:]
//
//	return
//}

// テーブル作成クエリの取得
func (self *Info) QueryCreateTable() (query string) {

	for _, field := range self.Fields {
		query += "  `" + field.Name + "`"

		switch field.Type {
		case TYPE_INT:
			query += " INT"
		case TYPE_MIDDLEINT:
			query += " MIDDLEINT"
		case TYPE_STRING:
			query += " VARCHAR"
		case TYPE_TEXT:
			query += " TEXT"
		case TYPE_BOOL:
			query += " bool"
		case TYPE_DATE:
			query += " DATETIME"
		case TYPE_TIME:
			query += " TIMESTAMP"
		}

		if field.Length > 0 {
			query += " (" + strconv.Itoa(field.Length) + ")"
		}

		// null setting
		if field.AutoTime != AUTO_TIME_DELETE {
			if !field.IsNull {
				query += " NOT NULL"
			}
		}

		if field.Name == "id" {
			query += " AUTO_INCREMENT"
		}

		if field.Default != "" {
			query += " DEFAULT " + field.Default
		}

		if field.Comment != "" {
			query += " COMMENT " + "'" + field.Comment + "'"
		}

		query += ",\n"

	}

	query += "  PRIMARY KEY (id),\n"

	// index
	if self.Indexs != nil {
		query += queryIndexs(self.Indexs, false)
	}

	// unique
	if self.Uniqs != nil {
		query += queryIndexs(self.Uniqs, true)
	}

	query = query[:len(query)-2] + "\n"

	query = "CREATE TABLE " + self.TableName + "(\n" + query + ")"

	return
}

// index設定クエリの取得
func queryIndexs(indexes [][]string, isUniq bool) (query string) {
	for _, index := range indexes {
		index_name := ""
		columns := ""
		for _, column := range index {
			index_name += "_" + column
			columns += "," + column
		}

		query += " "
		if isUniq {
			query += " UNIQUE"
		}
		query += " INDEX index_" + index_name[1:] + "(" + columns[1:] + "),\n"
	}
	return
}
