package table

import ()

// column type
type ColumnType int

// column type string
func (self ColumnType) String() string { return columnTypeString[self] }

// column type const
const (
	TYPE_INT ColumnType = iota
	TYPE_MIDDLEINT
	TYPE_STRING
	TYPE_DATE
	TYPE_TIME
	TYPE_BOOL
	TYPE_TEXT
	TYPE_MEDIUMTEXT
)

// column type const string
var columnTypeString = [...]string{
	"INT",
	"MIDDLEINT",
	"VARCHAR",
	"DATETIME",  // '1000-01-01 00:00:00' ～ '9999-12-31 23:59:59'
	"TIMESTAMP", // '1970-01-01 00:00:01' UTC ～ '2038-01-19 03:14:07' UTC
	"BOOL",
	"TEXT",
	"MEDIUMTEXT",
}

type AutoTime int

const (
	AUTO_TIME_NON AutoTime = iota
	AUTO_TIME_CREATE
	AUTO_TIME_UPDATE
	AUTO_TIME_DELETE
)

func (self ColumnType) IsInt() bool {
	return self == TYPE_INT
}

// フィールドタイプがStringであるかチェック
func (self ColumnType) IsString() bool {
	return self == TYPE_STRING
}

// フィールドタイプがDateであるかチェック
func (self ColumnType) IsDate() bool {
	return self == TYPE_DATE
}

// フィールドタイプがDateであるかチェック
func (self ColumnType) IsTime() bool {
	return self == TYPE_TIME
}

// フィールドタイプがBoolであるかチェック
func (self ColumnType) IsBool() bool {
	return self == TYPE_BOOL
}

func (self ColumnType) IsText() bool {
	return self == TYPE_TEXT
}

// 時間タイプがNoneであるかチェック
func (self AutoTime) IsNone() bool {
	return self == AUTO_TIME_NON
}
