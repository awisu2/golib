package db

import (
	_sql "database/sql"
)

// INFO:Null型をhtmlやtemplateで使いやすいように
type NullString _sql.NullString
type NullInt64 _sql.NullInt64
type NullBool _sql.NullBool
type NullFloat64 _sql.NullFloat64

// テキストに変換する
func (this NullString) Text() string {
	return NullStringToString(_sql.NullString(this))
}

// database/sqlの型に戻す
func (this NullString) Base() _sql.NullString {
	return _sql.NullString(this)
}

// テキストに変換する
func (this NullInt64) Text() string {
	return NullInt64ToString(_sql.NullInt64(this))
}

// database/sqlの型に戻す
func (this NullInt64) Base() _sql.NullInt64 {
	return _sql.NullInt64(this)
}

// テキストに変換する
func (this NullBool) Text() string {
	return NullBoolToString(_sql.NullBool(this))
}

// database/sqlの型に戻す
func (this NullBool) Base() _sql.NullBool {
	return _sql.NullBool(this)
}

// テキストに変換する
func (this NullFloat64) Text() string {
	return NullFloat64ToString(_sql.NullFloat64(this))
}

// database/sqlの型に戻す
func (this NullFloat64) Base() _sql.NullFloat64 {
	return _sql.NullFloat64(this)
}
