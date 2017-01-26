package table

import (
	"strconv"
)

// テーブル作成クエリの取得
func (self *Info) QueryCreateTable() (query string) {

	for _, field := range self.Fields {
		query += "  `" + field.Name + "` " + field.Type.String()

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
