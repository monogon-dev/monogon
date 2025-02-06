// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package clitable implements tabular display for command line tools.
//
// The generated tables are vaguely reminiscent of the output of 'kubectl'. For
// example:
//
//	NAME ADDRESS  STATUS
//	foo  1.2.3.4  Healthy
//	bar  1.2.3.12 Timeout
//
// The tables are sparse by design, with each Entry having a number of Key/Value
// entries. Then, all keys for all entries get unified into columns, and entries
// without a given key are rendered spare (i.e. the rendered cell is empty).
package clitable

import (
	"fmt"
	"io"
	"strings"
)

// Table is a list of entries that form a table with sparse columns. Each entry
// defines its own columns.
type Table struct {
	entries []Entry
}

// Add an entry to the table.
func (t *Table) Add(e Entry) {
	t.entries = append(t.entries, e)
}

// An Entry is made up of column key -> value pairs.
type Entry struct {
	columns []entryColumn
}

// Add a key/value pair to the entry.
func (e *Entry) Add(key, value string) {
	e.columns = append(e.columns, entryColumn{
		key:   key,
		value: value,
	})
}

// Get a value from a given key, returning zero string if not set.
func (e *Entry) Get(key string) string {
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

// Columns returns the keys and widths of columns that are present in the table's
// entries.
func (t *Table) Columns() Columns {
	var res Columns
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

type Columns []*Column

// A Column in a table, not containing all entries that make up this table, but
// containing their maximum width (for layout purposes).
type Column struct {
	// key is the column key.
	key string
	// width is the maximum width (in runes) of all the entries' data in this column.
	width int
}

// upsert a key into a list of columns, returning the upserted column.
func (c *Columns) upsert(key string) *Column {
	for _, col := range *c {
		if col.key == key {
			return col
		}
	}
	col := &Column{
		key:   key,
		width: len(key),
	}
	*c = append(*c, col)
	return col
}

// filter returns a copy of columns where the only columns present are the ones
// whose onlyColumns values are true. If only columns is nil, no filtering takes
// place (all columns are returned).
func (c Columns) filter(onlyColumns map[string]bool) Columns {
	var res []*Column
	for _, cc := range c {
		if onlyColumns != nil && !onlyColumns[cc.key] {
			continue
		}
		res = append(res, cc)
	}
	return res
}

// PrintHeader writes a table-like header to the given file, keeping margin
// spaces between columns.
func (c Columns) printHeader(f io.Writer, margin int) {
	for _, cc := range c {
		fmt.Fprintf(f, "%-*s", cc.width+margin, strings.ToUpper(cc.key))
	}
	fmt.Fprintf(f, "\n")
}

// Print writes a table-like representation of this table to the given file,
// first filtering the columns by onlyColumns (if not set, no filtering takes
// place).
func (t *Table) Print(f io.Writer, onlyColumns map[string]bool) {
	margin := 3
	cols := t.Columns().filter(onlyColumns)
	cols.printHeader(f, margin)

	for _, e := range t.entries {
		for _, c := range cols {
			v := e.Get(c.key)
			fmt.Fprintf(f, "%-*s", c.width+margin, v)
		}
		fmt.Fprintf(f, "\n")
	}
}
