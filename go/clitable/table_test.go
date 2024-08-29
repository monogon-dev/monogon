package clitable

import (
	"bytes"
	"strings"
	"testing"
)

// TestTableLayout performs a smoke test of the table layout functionality.
func TestTableLayout(t *testing.T) {
	var tab Table

	var e Entry
	e.Add("id", "short")
	e.Add("labels", "")
	tab.Add(e)

	e = Entry{}
	e.Add("whoops", "only in second")
	e.Add("labels", "bar")
	e.Add("id", "this one is a very long one")
	tab.Add(e)

	e = Entry{}
	e.Add("id", "normal length")
	e.Add("labels", "foo")
	tab.Add(e)

	buf := bytes.NewBuffer(nil)
	tab.Print(buf, nil)

	golden := `
ID                            LABELS   WHOOPS           
short                                                   
this one is a very long one   bar      only in second   
normal length                 foo                       
`
	golden = strings.TrimSpace(golden)
	got := strings.TrimSpace(buf.String())
	if got != golden {
		t.Logf("wanted: \n%s", golden)
		t.Logf("got: \n%s", got)
		t.Errorf("mismatch")
	}
}
