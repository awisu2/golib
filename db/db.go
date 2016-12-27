package db

import (
	_sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Host     string
	Database string
	DB       *_sql.DB
}

// 設定値に従ってコネクションを作成
func Open(host string, database string) (*DB, error) {
	//TODO:mysqlオンリーなのは改善したい
	_db, err := _sql.Open("mysql", host+"/"+database)
	if err != nil {
		return nil, err
	}

	return &DB{host, database, _db}, nil
}

// dbコネクションをクローズ
func (self *DB) Close() (err error) {
	return self.DB.Close()
}

// データの取得
func (self *DB) Query(query string, args ...interface{}) (*_sql.Rows, error) {
	return self.DB.Query(query, args...)
}

// insert,updateの実行
func (self *DB) Exec(query string, args ...interface{}) (_sql.Result, error) {
	return self.DB.Exec(query, args...)
}
