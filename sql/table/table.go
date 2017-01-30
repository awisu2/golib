package table

import (
	"html/template"
)

// フィールドチェック関数
type checkValue func(s string) (err error)

// 値表示時のデコレーション
type htmlDecoration func(s string) (html template.HTML)

// レコードのカラムの情報
type Field struct {
	Name       string
	Type       ColumnType
	Length     int
	AutoTime   AutoTime
	IsNull     bool
	Default    string
	Comment    string
	Checker    checkValue
	Decoration htmlDecoration
}

// 情報
type Info struct {
	TableName  string
	ConfigName string
	Fields     Fields
	Indexs     [][]string
	Uniqs      [][]string
}

// カラムをまとめたもの(keyをレコード名として使用)
type Fields []Field
