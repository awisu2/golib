package db

import (
	_sql "database/sql"
	"fmt"
	"github.com/awisu2/golib/db/table"
	"github.com/awisu2/golib/log"
)

const (
	ORDER_TYPE_ASC = iota
	ORDER_TYPE_DESC
)

const (
	JOIN_TYPE_INNER = iota + 1
	JOIN_TYPE_LEFT
	JOIN_TYPE_RIGHT
	JOIN_TYPE_FULL
)

const (
	LIMIT_OFFSET_NON = -1
)

const (
	WHERE_TYPE_EQUAL = iota
	WHERE_TYPE_IN
)

type sql struct {
	column       string
	Wheres       []*where
	Orders       []*order
	limit        *limit
	Sets         map[string]interface{}
	UseTableInfo bool
	Joins        []*Join
	GroupBy      string
	Security     *security
}

type where struct {
	Name  string
	Type  int
	Value interface{}
}

type order struct {
	Name      string
	OrderType int
}

type limit struct {
	offset   int
	rowcount int
}

type Join struct {
	Table string
	Alias string
	Type  int
	On    string
}

// 通常まずそうな処理を防止するフラグ
type security struct {
	CanNoWhereUpdateOrDelete bool // update,deleteをするときにwhereがないことを許容(default: false)
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
	return &sql{column: "*", UseTableInfo: true, Security: &security{CanNoWhereUpdateOrDelete: false}}
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
	self.Wheres = append(self.Wheres, &where{name, int(WHERE_TYPE_EQUAL), value})
	return self
}

// whereの設定
func (self *sql) WhereIn(name string, value interface{}) *sql {
	if self.Wheres == nil {
		self.Wheres = []*where{}
	}
	self.Wheres = append(self.Wheres, &where{name, int(WHERE_TYPE_IN), value})
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

func (self *sql) Group(group string) *sql {
	self.GroupBy = group
	return self
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
func (self *sql) SetValues(vals map[string]interface{}) *sql {
	for column, value := range vals {
		self.Set(column, value)
	}
	return self
}

// join分用
func (self *sql) Join(join *Join) *sql {
	if join == nil {
		return self
	}
	if self.Joins == nil {
		self.Joins = []*Join{}
	}
	self.Joins = append(self.Joins, join)
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

	datas, err = RowsToString(rows)
	return
}

// map形式にして取得
func (self *sql) SelectToMap(table string, db *DB) (datas map[string]RowData, err error) {
	query, args := self.QuerySelect(table)
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	datas, err = RowsToMap(rows)
	return
}

// select実行
func (self *sql) SelectToNullString(table string, db *DB) (datas []RowStrData, err error) {
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

	datas, err := RowsToString(rows)
	if err == nil && len(datas) > 0 {
		data = datas[0]
	}
	return
}

// Update実行
func (self *sql) Update(table string, db *DB) (result _sql.Result, err error) {
	// check where exist
	if self.Security.CanNoWhereUpdateOrDelete == false {
		if self.Wheres == nil || len(self.Wheres) == 0 {
			err = fmt.Errorf("no where by Update.")
			return
		}
	}

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
	// check where exist
	if self.Security.CanNoWhereUpdateOrDelete == false {
		if self.Wheres == nil || len(self.Wheres) == 0 {
			err = fmt.Errorf("no where by Delete.")
			return
		}
	}

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

// deleted_atを設定している場合に有効
func (self *sql) UnDelete(tableName string, db *DB) (result _sql.Result, err error) {
	query, args := self.QueryUnDelete(tableName)
	if query == "" {
		err = fmt.Errorf("no deleted_at column.")
		return
	}
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

func Query(query string, db *DB, args ...interface{}) (datas []RowStrData, err error) {
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

// レコードの存在チェック
func IsExist(tableName string, name string, value interface{}, db *DB) bool {
	rows, err := db.Query(NewSql().Column("id").Where(name, value).QuerySelect(tableName))
	if err != nil {
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	}

	return true
}

// テーブルのレコード数取得
func Count(table string, db *DB, isDeleteFlag bool) (int, error) {
	sql := NewSql().Column("count(*) as count")
	if isDeleteFlag {
		sql.Where("deleted_at", nil)
	}
	data, err := sql.SelectRow(table, db)
	if err != nil {
		return 0, err
	}
	return data.Int("count"), nil
}
