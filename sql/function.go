package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// open db connection
func Open(host string, database string) (db *DB, err error) {
	//TODO:mysqlオンリーなのは改善したい
	_db, err := sql.Open("mysql", host+"/"+database)
	if err != nil {
		return
	}
	db = &DB{_db}
	return
}

// open db connection by config
func OpenByConfig(config *Config) (db *DB, err error) {
	return Open(config.Host, config.Database)
}

// execute select query
func Query(query string, db *DB, args ...interface{}) (datas []RowStrData, err error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	datas, err = RowsToDatas(rows)
	return
}

// execute query
func Exec(query string, db *DB, args ...interface{}) (result sql.Result, err error) {
	result, err = db.Exec(query, args...)
	if err != nil {
		return
	}
	return
}

// check table exists
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
func TruncateTable(tableName string, db *DB) (result sql.Result, err error) {
	return Exec(QueryTruncate(tableName), db)
}

// テーブル削除
func DropTable(tableName string, db *DB) (result sql.Result, err error) {
	return Exec(QueryDrop(tableName), db)
}

// レコードの存在チェック
func IsExist(tableName string, name string, value interface{}, db *DB) bool {
	query, args := NewQuery().Column("id").Where(name, value).LimitRowcount(1).QuerySelect(tableName)
	rows, err := db.Query(query, args...)
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
	sql := NewQuery().Column("count(*) as count")
	if isDeleteFlag {
		sql.Where("deleted_at", nil)
	}
	data, err := sql.SelectRow(table, db)
	if err != nil {
		return 0, err
	}
	return data.Int("count"), nil
}
