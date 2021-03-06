package sql

import (
	"fmt"
	"github.com/awisu2/golib/sql/table"
	"strings"
)

func (self *query) QuerySelect(tableName string) (query string, args []interface{}) {
	args = []interface{}{}

	// join
	query += self.QueryJoin()

	// where
	whereQuery, whereArgs := self.QueryWhere()
	query += whereQuery
	args = append(args, whereArgs...)

	query += self.QueryGroup()
	query += self.QueryOrder()
	query += self.QueryLimit()

	query = "SELECT " + self.column + " FROM " + tableName + query

	return
}

func (self *query) QueryInsert(tableName string) (query string, args []interface{}) {
	// table.Infoの取得
	var info *table.Info
	if self.UseTableInfo && tableInfos != nil {
		info = tableInfos[tableName]
	}

	args = []interface{}{}

	if self.Sets != nil {
		qBefore := ""
		qAfter := ""
		for column, value := range self.Sets {
			qBefore += ", " + column
			qAfter += ", ?"
			args = append(args, value)
		}

		// AutoTime Fields
		if info != nil {
			now := Now()
			for _, field := range info.Fields {
				if field.AutoTime == table.AUTO_TIME_CREATE || field.AutoTime == table.AUTO_TIME_UPDATE {
					qBefore += ", " + field.Name
					qAfter += ", ?"
					args = append(args, now)
				}
			}
		}

		query += " (" + qBefore[2:] + ") VALUES(" + qAfter[2:] + ")"
	}

	query = "INSERT INTO " + tableName + query

	return
}

func (self *query) QueryInserts(tableName string, vals []map[string]interface{}) (query string, args []interface{}) {
	// table.Infoの取得
	var info *table.Info
	if self.UseTableInfo && tableInfos != nil {
		info = tableInfos[tableName]
	}

	if len(vals) <= 0 {
		return
	}

	// 先頭のデータに存在するキーをカラムとして取得
	columns := []string{}
	qColumns := ""
	for k, _ := range vals[0] {
		columns = append(columns, k)
		qColumns += ", " + k
	}

	// AutoTime Columns
	autoTimeColumns := []string{}
	if info != nil {
		for _, field := range info.Fields {
			if field.AutoTime == table.AUTO_TIME_CREATE || field.AutoTime == table.AUTO_TIME_UPDATE {
				qColumns += ", " + field.Name
				autoTimeColumns = append(autoTimeColumns, field.Name)
			}
		}
	}
	qColumns = qColumns[2:]

	args = []interface{}{}
	now := Now()
	for _, fields := range vals {
		q := ""
		for _, column := range columns {
			v, _ := fields[column]
			q += ", ?"
			args = append(args, v)
		}

		// AutoTime Columns
		for range autoTimeColumns {
			q += ", ?"
			args = append(args, now)
		}
		query += ",(" + q[2:] + ")"
	}

	query = "INSERT INTO " + tableName + "(" + qColumns + ") VALUES " + query[1:]

	return
}

func (self *query) QueryUpdate(tableName string) (query string, args []interface{}) {
	args = []interface{}{}

	setQuery, setArgs := self.QuerySet(tableName)
	query += setQuery
	args = append(args, setArgs...)

	// where
	whereQuery, whereArgs := self.QueryWhere()
	query += whereQuery
	args = append(args, whereArgs...)

	query = "UPDATE " + tableName + query

	return
}

// joinクエリ作成
func (self *query) QueryJoin() (query string) {
	if self.Joins == nil {
		return
	}

	for _, join := range self.Joins {
		switch join.Type {
		case JOIN_TYPE_INNER:
			query += " INNER JOIN "
		case JOIN_TYPE_LEFT:
			query += " LEFT OUTER JOIN "
		case JOIN_TYPE_RIGHT:
			query += " RIGHT OUTER JOIN "
		case JOIN_TYPE_FULL:
			query += " FULL OUTER JOIN "
		}
		query += join.Table
		if join.Alias != "" {
			query += " AS " + join.Alias
		}
		if join.On != "" {
			query += " ON " + join.On
		}
	}
	return
}

func (self *query) QueryWhere() (query string, args []interface{}) {
	if self.Wheres != nil {
		q := ""
		for _, v := range self.Wheres {
			if v.Value == nil {
				q += " AND " + v.Name + " IS NULL"
			} else {
				switch v.Operator {
				case WHERE_OPERATOR_IN:
					arr, ok := v.Value.([]interface{})
					if ok && len(arr) > 0 {
						q += " AND " + v.Name + " IN(" + strings.Repeat(",?", len(arr))[1:] + ")"
						args = append(args, arr...)
					} else {
						q += " AND " + v.Name + " IN(?)"
						args = append(args, "NULL")
					}

				default:
					q += " AND " + v.Name + " " + v.Operator.String() + " ?"
					args = append(args, v.Value)
				}
			}
		}
		if q != "" {
			query += " WHERE " + q[5:]
		}
	}
	return
}

func (self *query) QueryOrder() (query string) {
	if self.Orders != nil {
		q := ""
		for _, v := range self.Orders {
			q += ", " + v.Name
			switch v.OrderType {
			case ORDER_TYPE_ASC:
				q += " ASC"
			case ORDER_TYPE_DESC:
				q += " DESC"
			}
		}
		query += " ORDER BY " + q[2:]
	}
	return
}

func (self *query) QueryGroup() (query string) {
	if self.GroupBy != "" {
		query = " GROUP BY " + self.GroupBy
	}
	return
}

func (self *query) QueryLimit() (query string) {
	if self.limit != nil {
		q := ""
		if self.limit.offset != LIMIT_OFFSET_NON {
			q += fmt.Sprintf("%v, %v", self.limit.offset, self.limit.rowcount)
		} else {
			q += fmt.Sprintf("%v", self.limit.rowcount)
		}
		query += " LIMIT " + q
	}
	return
}

func (self *query) QuerySet(tableName string) (query string, args []interface{}) {
	// table.Infoの取得
	var info *table.Info
	if self.UseTableInfo && tableInfos != nil {
		info = tableInfos[tableName]
	}

	args = []interface{}{}

	if self.Sets != nil {
		q := ""
		for column, value := range self.Sets {
			q += ", " + column + "=?"
			args = append(args, value)
		}

		// テーブル情報をもとにdateを更新
		if info != nil {
			now := Now()
			for _, field := range info.Fields {
				if field.AutoTime == table.AUTO_TIME_UPDATE {
					q += ", " + field.Name + "=?"
					args = append(args, now)
				}
			}
		}

		query += " SET " + q[2:]
	}

	return
}

func (self *query) QueryDelete(tableName string) (query string, args []interface{}) {
	// table.Infoの取得
	var info *table.Info
	if self.UseTableInfo && tableInfos != nil {
		info = tableInfos[tableName]
	}

	// infoが見つからない場合はレコード削除に切り替え
	if info == nil {
		return self.QueryForceDelete(tableName)
	}

	isDelete := false
	q := ""
	args = []interface{}{}

	for _, field := range info.Fields {
		if field.AutoTime == table.AUTO_TIME_DELETE {
			isDelete = true
			q += ", " + field.Name + "=?"
			args = append(args, Now())
		}
	}

	// delete_atがない場合レコード削除に切り替え
	if !isDelete {
		return self.QueryForceDelete(tableName)
	}

	query = "UPDATE " + tableName + " SET " + q[2:]

	// where
	whereQuery, whereArgs := self.QueryWhere()
	query += whereQuery
	args = append(args, whereArgs...)

	return
}

func (self *query) QueryUnDelete(tableName string) (query string, args []interface{}) {
	args = []interface{}{}

	// table.Infoの取得
	var info *table.Info
	if self.UseTableInfo && tableInfos != nil {
		info = tableInfos[tableName]
	}
	if info == nil {
		return
	}

	// deleteカラム
	q := ""
	for _, field := range info.Fields {
		if field.AutoTime == table.AUTO_TIME_DELETE {
			q += ", " + field.Name + "=?"
			args = append(args, nil)
		}
	}
	query += " SET " + q[2:]

	// where
	whereQuery, whereArgs := self.QueryWhere()
	query += whereQuery
	args = append(args, whereArgs...)

	query = "UPDATE " + tableName + query

	return
}

func (self *query) QueryForceDelete(tableName string) (query string, args []interface{}) {
	query = "DELETE FROM " + tableName

	// where
	whereQuery, whereArgs := self.QueryWhere()
	query += whereQuery
	args = append(args, whereArgs...)

	return
}

func QueryShowCreateTable(tableName string) string {
	return "SHOW CREATE TABLE " + tableName
}

func QueryTruncate(tableName string) string {
	return "TRUNCATE TABLE " + tableName
}

func QueryDrop(tableName string) string {
	return "DROP TABLE " + tableName
}
