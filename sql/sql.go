package sql

import (
	"database/sql"
	"fmt"
	"github.com/awisu2/golib/log"

	"strconv"
)

type DB struct {
	*sql.DB
}

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

//　whereのoperator
type WhereOperator int

// 文字列として取得
func (self WhereOperator) String() string { return whereOperatorString[self-1] }

// where operatorの実値
const (
	WHERE_OPERATOR_EQ WhereOperator = iota + 1
	WHERE_OPERATOR_NE               // !=
	WHERE_OPERATOR_LT               // <
	WHERE_OPERATOR_LE               // <=
	WHERE_OPERATOR_GT               // >
	WHERE_OPERATOR_GE               // >=
	WHERE_OPERATOR_IN
)

// where operator用文字
var whereOperatorString = [...]string{
	"=",
	"!=",
	"<",
	"<=",
	">",
	">=",
	"IN",
}

type query struct {
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
	Name     string
	Operator WhereOperator
	Value    interface{}
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

// get new query
func NewQuery() *query {
	return &query{column: "*", UseTableInfo: true, Security: &security{CanNoWhereUpdateOrDelete: false}}
}

// select時の取得columnを設定
// セットした値は特に変換されずそのまま使用される(例: a as a, b)
func (self *query) Column(column string) *query {
	self.column = column
	return self
}

// whereの設定
func (self *query) Where(name string, value interface{}) *query {
	if self.Wheres == nil {
		self.Wheres = []*where{}
	}
	self.Wheres = append(self.Wheres, &where{name, WHERE_OPERATOR_EQ, value})
	return self
}

// whereの設定
func (self *query) WhereIn(name string, value interface{}) *query {
	if self.Wheres == nil {
		self.Wheres = []*where{}
	}
	self.Wheres = append(self.Wheres, &where{name, WHERE_OPERATOR_IN, value})
	return self
}

// whereの設定
func (self *query) WhereO(name string, value interface{}, operator WhereOperator) *query {
	if self.Wheres == nil {
		self.Wheres = []*where{}
	}
	self.Wheres = append(self.Wheres, &where{name, operator, value})
	return self
}

// where条件をクリア
func (self *query) ClearWhere() *query {
	self.Wheres = nil
	return self
}

// order asc を設定
func (self *query) OrderAsc(name string) *query {
	if self.Orders == nil {
		self.Orders = []*order{}
	}
	self.Orders = append(self.Orders, &order{name, ORDER_TYPE_ASC})
	return self
}

// order desc を設定
func (self *query) OrderDesc(name string) *query {
	if self.Orders == nil {
		self.Orders = []*order{}
	}
	self.Orders = append(self.Orders, &order{name, ORDER_TYPE_DESC})
	return self
}

func (self *query) Limit(offset int, rowcount int) *query {
	if self.limit == nil {
		self.limit = &limit{offset, rowcount}
	} else {
		self.limit.offset = offset
		self.limit.rowcount = rowcount
	}
	return self
}

func (self *query) LimitRowcount(rowcount int) *query {
	return self.Limit(LIMIT_OFFSET_NON, rowcount)
}

func (self *query) Group(group string) *query {
	self.GroupBy = group
	return self
}

// insertまたはupdate用の値をセット
func (self *query) Set(column string, value interface{}) *query {
	if self.Sets == nil {
		self.Sets = map[string]interface{}{}
	}
	self.Sets[column] = value
	return self
}

// insertまたはupdate用の値を複数セット
func (self *query) SetValues(vals map[string]interface{}) *query {
	for column, value := range vals {
		self.Set(column, value)
	}
	return self
}

// join分用
func (self *query) Join(join *Join) *query {
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
func (self *query) Select(table string, db *DB) (datas []RowData, err error) {
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
func (self *query) SelectToMap(table string, db *DB) (datas map[string]RowData, err error) {
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
func (self *query) SelectToNullString(table string, db *DB) (datas []RowStrData, err error) {
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
func (self *query) SelectRow(table string, db *DB) (data RowData, err error) {
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

func (self *query) Count(table string, db *DB, useLimit bool) (count int, err error) {
	column, limit := self.column, self.limit
	self.column = "count(*) AS count"
	if !useLimit {
		self.limit = nil
	}
	data, err := self.SelectRow(table, db)
	self.column, self.limit = column, limit
	if err != nil {
		return
	}
	if data == nil {
		return
	}

	_count := data["count"]
	count, _ = strconv.Atoi(_count)

	return
}

// Update実行
func (self *query) Update(table string, db *DB) (result sql.Result, err error) {
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
func (self *query) Insert(tableName string, db *DB) (result sql.Result, err error) {
	query, args := self.QueryInsert(tableName)
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

// Insert実行
func (self *query) Inserts(tableName string, db *DB, vals []map[string]interface{}) (result sql.Result, err error) {
	query, args := self.QueryInserts(tableName, vals)
	log.Println(query)
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

// Delete実行
func (self *query) Delete(tableName string, db *DB, isForce bool) (result sql.Result, err error) {
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
func (self *query) UnDelete(tableName string, db *DB) (result sql.Result, err error) {
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
