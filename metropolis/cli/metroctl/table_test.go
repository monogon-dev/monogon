package main

import (
	"bytes"
	"strings"
	"testing"
)

// TestTableLayout performs a smoke test of the table layout functionality.
func TestTableLayout(t *testing.T) {
	tab := table{}

	e := entry{}
	e.add("id", "short")
	e.add("labels", "")
	tab.add(e)

	e = entry{}
	e.add("whoops", "only in second")
	e.add("labels", "bar")
	e.add("id", "this one is a very long one")
	tab.add(e)

	e = entry{}
	e.add("id", "normal length")
	e.add("labels", "foo")
	tab.add(e)

	buf := bytes.NewBuffer(nil)
	tab.print(buf, nil)

	golden := `
ID                            LABELS   WHOOPS           
short                                                   
this one is a very long one   bar      only in second   
normal length                 foo                       
`
	golden = strings.TrimSpace(golden)
	got := strings.TrimSpace(buf.String())
	if got != golden {
		t.Logf("wanted: \n" + golden)
		t.Logf("got: \n" + got)
		t.Errorf("mismatch")
	}
}
