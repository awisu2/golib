package db

import (
	_sql "database/sql"
	"strconv"
)

func NullInt64ToString(v _sql.NullInt64) (r string) {
	if !v.Valid {
		return
	}
	r = strconv.Itoa(int(v.Int64))
	return
}

func NullInt64ToInt(v _sql.NullInt64) (r int) {
	if !v.Valid {
		return
	}
	r = int(v.Int64)
	return
}

func NullStringToString(v _sql.NullString) (r string) {
	if !v.Valid {
		return
	}
	r = v.String

	return
}

func NullBoolToString(v _sql.NullBool) (r string) {
	if !v.Valid {
		return
	}
	r = "false"
	if v.Bool {
		r = "true"
	}
	return
}

func NullFloat64ToString(v _sql.NullFloat64) (r string) {
	if !v.Valid {
		return
	}
	r = strconv.FormatFloat(v.Float64, 'E', -1, 64)
	return
}
