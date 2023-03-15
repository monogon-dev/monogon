package tinylb

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"

	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

func TestLoadbalancer(t *testing.T) {
	v := memory.Value[BackendSet]{}
	set := BackendSet{}
	v.Set(set.Clone())

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Listen failed: %v", err)
	}
	s := Server{
		Provider: &v,
		Listener: ln,
	}
	supervisor.TestHarness(t, s.Run)

	connect := func() net.Conn {
		conn, err := net.Dial("tcp", ln.Addr().String())
		if err != nil {
			t.Fatalf("Connection failed: %v", err)
		}
		return conn
	}

	c := connect()
	buf := make([]byte, 128)
	if _, err := c.Read(buf); err == nil {
		t.Fatalf("Expected error on read (no backends yet)")
	}

	// Now add a backend and expect it to be served.
	makeBackend := func(hello string) net.Listener {
		aln, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatalf("Failed to make backend listener: %v", err)
		}
		// Start backend.
		go func() {
			for {
				c, err := aln.Accept()
				if err != nil {
					return
				}
				// For each connection, keep writing 'hello' over and over, newline-separated.
				go func() {
					defer c.Close()
					for {
						if _, err := fmt.Fprintf(c, "%s\n", hello); err != nil {
							return
						}
						time.Sleep(100 * time.Millisecond)
					}
				}()
			}
		}()
		addr := aln.Addr().(*net.TCPAddr)
		set.Insert(hello, &SimpleTCPBackend{Remote: addr.AddrPort().String()})
		v.Set(set.Clone())
		return aln
	}

	as1 := makeBackend("a")
	defer as1.Close()

	for {
		c = connect()
		_, err := c.Read(buf)
		c.Close()
		if err == nil {
			break
		}
	}

	measure := func() map[string]int {
		res := make(map[string]int)
		for {
			count := 0
			for _, v := range res {
				count += v
			}
			if count >= 20 {
				return res
			}

			c := connect()
			b := bufio.NewScanner(c)
			if !b.Scan() {
				err := b.Err()
				if err == nil {
					err = io.EOF
				}
				t.Fatalf("Scan failed: %v", err)
			}
			v := b.Text()
			res[v]++
			c.Close()
		}
	}

	m := measure()
	if m["a"] < 20 {
		t.Errorf("Expected only one backend, got: %v", m)
	}

	as2 := makeBackend("b")
	defer as2.Close()

	as3 := makeBackend("c")
	defer as3.Close()

	as4 := makeBackend("d")
	defer as4.Close()

	m = measure()
	for _, id := range []string{"a", "b", "c", "d"} {
		if want, got := 4, m[id]; got < want {
			t.Errorf("Expected at least %d responses from %s, got %d", want, id, got)
		}
	}

	// Test killing backend connections on backend removal.
	// Open a bunch of connections to 'a'.
	var conns []*bufio.Scanner
	for len(conns) < 5 {
		c := connect()
		b := bufio.NewScanner(c)
		b.Scan()
		if b.Text() != "a" {
			c.Close()
		} else {
			conns = append(conns, b)
		}
	}

	// Now remove the 'a' backend.
	set.Delete("a")
	v.Set(set.Clone())
	// All open connections should now get killed.
	for _, b := range conns {
		start := time.Now().Add(time.Second)
		for b.Scan() {
			if time.Now().After(start) {
				t.Errorf("Connection still alive")
				break
			}
		}
	}
}

func BenchmarkLB(b *testing.B) {
	v := memory.Value[BackendSet]{}
	set := BackendSet{}
	v.Set(set.Clone())

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		b.Fatalf("Listen failed: %v", err)
	}
	s := Server{
		Provider: &v,
		Listener: ln,
	}
	supervisor.TestHarness(b, s.Run)

	makeBackend := func(hello string) net.Listener {
		aln, err := net.Listen("tcp", ":0")
		if err != nil {
			b.Fatalf("Failed to make backend listener: %v", err)
		}
		// Start backend.
		go func() {
			for {
				c, err := aln.Accept()
				if err != nil {
					return
				}
				go func() {
					fmt.Fprintf(c, "%s\n", hello)
					c.Close()
				}()
			}
		}()
		addr := aln.Addr().(*net.TCPAddr)
		set.Insert(hello, &SimpleTCPBackend{Remote: addr.AddrPort().String()})
		v.Set(set.Clone())
		return aln
	}
	var backends []net.Listener
	for i := 0; i < 10; i++ {
		b := makeBackend(fmt.Sprintf("backend%d", i))
		backends = append(backends, b)
	}

	defer func() {
		for _, b := range backends {
			b.Close()
		}
	}()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			conn, err := net.Dial("tcp", ln.Addr().String())
			if err != nil {
				b.Fatalf("Connection failed: %v", err)
			}
			buf := bufio.NewScanner(conn)
			buf.Scan()
			if !strings.HasPrefix(buf.Text(), "backend") {
				b.Fatalf("Invalid backend response: %q", buf.Text())
			}
			conn.Close()
		}
	})
	b.StopTimer()
}
