package db

import (
	_sql "database/sql"
	"github.com/awisu2/golib/db/table"
	"github.com/awisu2/golib/log"
)

const (
	ORDER_TYPE_ASC = iota
	ORDER_TYPE_DESC
)

const (
	LIMIT_OFFSET_NON = -1
)

type sql struct {
	column       string
	Wheres       []*where
	Orders       []*order
	limit        *limit
	Sets         map[string]interface{}
	Updated_at   []string
	Deleted_at   []string
	Created_at   []string
	UseTableInfo bool
}

type where struct {
	Name      string
	WhereType string
	Value     interface{}
}

type order struct {
	Name      string
	OrderType int
}

type limit struct {
	offset   int
	rowcount int
}

// table操作Interface
type tableInterface interface {
	// テーブル名を指定しTable.Infoを返却
	GetInfo(string) *table.Info
}

// インスタンス
var _table tableInterface

// セットしておくとテーブル名をベースに操作を自動化
// Updated_at, Deleted_at, Created_atを自動セット
func SetTable(t tableInterface) {
	_table = t
}

func NewSql() *sql {
	return &sql{column: "*", UseTableInfo: true}
}

// select時の取得columnを設定
// セットした値は特に変換されずそのまま使用される(例: a as a, b)
func (self *sql) Column(column string) *sql {
	self.column = column
	return self
}

// whereの設定
func (self *sql) Where(name string, value interface{}) *sql {
	if self.Wheres == nil {
		self.Wheres = []*where{}
	}
	self.Wheres = append(self.Wheres, &where{name, "", value})
	return self
}

// where条件をクリア
func (self *sql) ClearWhere() *sql {
	self.Wheres = nil
	return self
}

// order asc を設定
func (self *sql) OrderAsc(name string) *sql {
	if self.Orders == nil {
		self.Orders = []*order{}
	}
	self.Orders = append(self.Orders, &order{name, ORDER_TYPE_ASC})
	return self
}

// order desc を設定
func (self *sql) OrderDesc(name string) *sql {
	if self.Orders == nil {
		self.Orders = []*order{}
	}
	self.Orders = append(self.Orders, &order{name, ORDER_TYPE_DESC})
	return self
}

func (self *sql) Limit(offset int, rowcount int) *sql {
	if self.limit == nil {
		self.limit = &limit{offset, rowcount}
	} else {
		self.limit.offset = offset
		self.limit.rowcount = rowcount
	}
	return self
}

func (self *sql) LimitRowcount(rowcount int) *sql {
	return self.Limit(LIMIT_OFFSET_NON, rowcount)
}

// insertまたはupdate用の値をセット
func (self *sql) Set(column string, value interface{}) *sql {
	if self.Sets == nil {
		self.Sets = map[string]interface{}{}
	}
	self.Sets[column] = value
	return self
}

// insertまたはupdate用の値を複数セット
func (self *sql) SetValues(sets map[string]interface{}) *sql {
	for column, value := range sets {
		self.Set(column, value)
	}
	return self
}

// select実行
func (self *sql) Select(table string, db *DB) (datas []RowData, err error) {
	query, args := self.QuerySelect(table)
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	datas, err = RowsToDatas(rows)
	return
}

// select実行
func (self *sql) SelectRow(table string, db *DB) (data RowData, err error) {
	query, args := self.QuerySelect(table)
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	datas, err := RowsToDatas(rows)
	if err == nil && len(datas) > 0 {
		data = datas[0]
	}
	return
}

// Update実行
func (self *sql) Update(table string, db *DB) (result _sql.Result, err error) {
	query, args := self.QueryUpdate(table)
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

// Insert実行
func (self *sql) Insert(tableName string, db *DB) (result _sql.Result, err error) {
	query, args := self.QueryInsert(tableName)
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

// Insert実行
func (self *sql) Inserts(tableName string, db *DB, vals []map[string]interface{}) (result _sql.Result, err error) {
	query, args := self.QueryInserts(tableName, vals)
	log.Println(query)
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

// Delete実行
func (self *sql) Delete(tableName string, db *DB, isForce bool) (result _sql.Result, err error) {
	var query string
	var args []interface{}
	if isForce {
		query, args = self.QueryForceDelete(tableName)
	} else {
		query, args = self.QueryDelete(tableName)
	}
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

func Query(query string, db *DB, args ...interface{}) (datas []RowData, err error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	datas, err = RowsToDatas(rows)
	return
}

func Exec(query string, db *DB, args ...interface{}) (result _sql.Result, err error) {
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

// テーブルの存在確認
func IsExistTable(tableName string, db *DB) bool {
	rows, err := db.Query("SHOW TABLES like '" + tableName + "'")
	if err != nil {
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	}

	return true
}

// truncate
func TruncateTable(tableName string, db *DB) (result _sql.Result, err error) {
	return Exec(QueryTruncate(tableName), db)
}

// テーブル削除
func DropTable(tableName string, db *DB) (result _sql.Result, err error) {
	return Exec(QueryDrop(tableName), db)
}
