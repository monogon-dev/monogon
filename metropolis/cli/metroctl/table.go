package main

import (
	"fmt"
	"io"
	"strings"
)

// table is a list of entries that form a table with sparse columns. Each entry
// defines its own columns.
type table struct {
	entries []entry
}

// add an entry to the table.
func (t *table) add(e entry) {
	t.entries = append(t.entries, e)
}

// An entry is made up of column key -> value pairs.
type entry struct {
	columns []entryColumn
}

// add a key/value pair to the entry.
func (e *entry) add(key, value string) {
	e.columns = append(e.columns, entryColumn{
		key:   key,
		value: value,
	})
}

// get a value from a given key, returning zero string if not set.
func (e *entry) get(key string) string {
	for _, col := range e.columns {
		if col.key == key {
			return col.value
		}
	}
	return ""
}

// An entryColumn is a pair for table column key and entry column value.
type entryColumn struct {
	key   string
	value string
}

// columns returns the keys and widths of columns that are present in the table's
// entries.
func (t *table) columns() columns {
	var res columns
	for _, e := range t.entries {
		for _, c := range e.columns {
			tc := res.upsert(c.key)
			if len(c.value) > tc.width {
				tc.width = len(c.value)
			}
		}
	}
	return res
}

type columns []*column

// A column in a table, not containing all entries that make up this table, but
// containing their maximum width (for layout purposes).
type column struct {
	// key is the column key.
	key string
	// width is the maximum width (in runes) of all the entries' data in this column.
	width int
}

// upsert a key into a list of columns, returning the upserted column.
func (c *columns) upsert(key string) *column {
	for _, col := range *c {
		if col.key == key {
			return col
		}
	}
	col := &column{
		key:   key,
		width: len(key),
	}
	*c = append(*c, col)
	return col
}

// filter returns a copy of columns where the only columns present are the ones
// whose onlyColumns values are true. If only columns is nil, no filtering takes
// place (all columns are returned).
func (c columns) filter(onlyColumns map[string]bool) columns {
	var res []*column
	for _, cc := range c {
		if onlyColumns != nil && !onlyColumns[cc.key] {
			continue
		}
		res = append(res, cc)
	}
	return res
}

// printHeader writes a table-like header to the given file, keeping margin
// spaces between columns.
func (c columns) printHeader(f io.Writer, margin int) {
	for _, cc := range c {
		fmt.Fprintf(f, "%-*s", cc.width+margin, strings.ToUpper(cc.key))
	}
	fmt.Fprintf(f, "\n")
}

// print writes a table-like representation of this table to the given file,
// first filtering the columns by onlyColumns (if not set, no filtering takes
// place).
func (t *table) print(f io.Writer, onlyColumns map[string]bool) {
	margin := 3
	cols := t.columns().filter(onlyColumns)
	cols.printHeader(f, margin)

	for _, e := range t.entries {
		for _, c := range cols {
			v := e.get(c.key)
			fmt.Fprintf(f, "%-*s", c.width+margin, v)
		}
		fmt.Fprintf(f, "\n")
	}
}
