package scylla

import (
	"github.com/scylladb/gocqlx/v2/table"
)

type query struct {
	stmt  string
	names []string
}

type Table struct {
	del query
	get query
	ins query
	sel query
}

func (t Table) Del() (string, []string) {
	return t.del.stmt, t.del.names
}

func (t Table) Get() (string, []string) {
	return t.get.stmt, t.get.names
}

func (t Table) Ins() (string, []string) {
	return t.ins.stmt, t.ins.names
}

func (t Table) Sel() (string, []string) {
	return t.sel.stmt, t.sel.names
}

func NewTable(md table.Metadata) Table {
	var result Table

	t := table.New(md)

	stmt, names := t.Delete()
	result.del = query{stmt: stmt, names: names}

	stmt, names = t.Get()
	result.get = query{stmt: stmt, names: names}

	stmt, names = t.Insert()
	result.ins = query{stmt: stmt, names: names}

	stmt, names = t.Select()
	result.sel = query{stmt: stmt, names: names}

	return result
}
