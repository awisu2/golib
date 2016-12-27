package table

import ()

const (
	TYPE_INT = iota
	TYPE_MIDDLEINT
	TYPE_STRING
	TYPE_DATE
	TYPE_TIME
	TYPE_BOOL
	TYPE_TEXT
)

const (
	AUTO_TIME_NON = iota
	AUTO_TIME_CREATE
	AUTO_TIME_UPDATE
	AUTO_TIME_DELETE
)

// フィールドタイプがintであるかチェック
func IsDataTypeInt(data_type int) bool {
	return data_type == int(TYPE_INT)
}

// フィールドタイプがStringであるかチェック
func IsDataTypeString(data_type int) bool {
	return data_type == int(TYPE_STRING)
}

// フィールドタイプがDateであるかチェック
func IsDataTypeDate(data_type int) bool {
	return data_type == int(TYPE_DATE)
}

// フィールドタイプがDateであるかチェック
func IsDataTypeTime(data_type int) bool {
	return data_type == int(TYPE_TIME)
}

// フィールドタイプがBoolであるかチェック
func IsDataTypeBool(data_type int) bool {
	return data_type == int(TYPE_BOOL)
}

func IsDataTypeText(date_type int) bool {
	return date_type == int(TYPE_TEXT)
}

// 時間タイプがNoneであるかチェック
func IsDateNone(date_type int) bool {
	return date_type == int(AUTO_TIME_NON)
}
