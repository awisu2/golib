package sql

import (
	"errors"
	"github.com/awisu2/golib/sql/table"
)

// tableInfos
// auto query custom
// example: auto set Updated_at, Deleted_at, Created_at
var tableInfos = map[string]*table.Info{}

// get tableInfos
func GetTableInfos() map[string]*table.Info {
	return tableInfos
}

// get tableInfo
func GetTableInfo(name string) (*table.Info, error) {
	v, ok := tableInfos[name]
	if !ok {
		return nil, errors.New("no exists table.Info. name:" + name)
	}
	return v, nil
}

// set tableInfos
func SetTableInfos(infos map[string]*table.Info) {
	tableInfos = infos
}

// add tableInfo
func AddTableInfo(name string, tableInfo *table.Info) {
	tableInfos[name] = tableInfo
}
