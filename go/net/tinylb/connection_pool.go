package tinylb

import (
	"net"
	"sort"
	"sync"
)

// connectionPool maintains information about open connections to backends, and
// allows for closing either arbitrary connections (by ID) or all connections to
// a given backend.
//
// This structure exists to allow tinylb to kill all connections of a backend
// that has just been removed from the BackendSet.
//
// Any time a connection is inserted into the pool, a unique ID for that
// connection is returned.
//
// Backends are identified by 'target name' which is an opaque string.
//
// This structure is likely the performance bottleneck of the implementation, as
// it takes a non-RW lock for every incoming connection.
type connectionPool struct {
	mu sync.Mutex
	// detailsById maps connection ids to details about that connection.
	detailsById map[int64]*connectionDetails
	// idsByTarget maps a target name to all connection IDs that opened to it.
	idsByTarget map[string][]int64

	// cid is the connection id counter, increased any time a connection ID is
	// allocated.
	cid int64
}

// connectionDetails for each open connection. These are held in
// connectionPool.details
type connectionDetails struct {
	// conn is the active net.Conn backing this connection.
	conn net.Conn
	// target is the target name to which this connection was initiated.
	target string
}

func newConnectionPool() *connectionPool {
	return &connectionPool{
		detailsById: make(map[int64]*connectionDetails),
		idsByTarget: make(map[string][]int64),
	}
}

// Insert a connection that's handled by the given target name, and return the
// connection ID used to remove this connection later.
func (c *connectionPool) Insert(target string, conn net.Conn) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	id := c.cid
	c.cid++

	c.detailsById[id] = &connectionDetails{
		conn:   conn,
		target: target,
	}
	c.idsByTarget[target] = append(c.idsByTarget[target], id)
	return id
}

// CloseConn closes the underlying connection for the given connection ID, and
// removes that connection ID from internal tracking.
func (c *connectionPool) CloseConn(id int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cd, ok := c.detailsById[id]
	if !ok {
		return
	}

	ids := c.idsByTarget[cd.target]
	// ids is technically sorted because 'id' is always monotonically increasing, so
	// we could be smarter and do a binary search here.
	ix := -1
	for i, id2 := range ids {
		if id2 == id {
			ix = i
			break
		}
	}
	if ix == -1 {
		panic("Programming error: connection present in detailsById but not in idsByTarget")
	}
	c.idsByTarget[cd.target] = append(ids[:ix], ids[ix+1:]...)
	cd.conn.Close()
	delete(c.detailsById, id)
}

// CloseTarget closes all connections to a given backend target name, and removes
// them from internal tracking.
func (c *connectionPool) CloseTarget(target string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, id := range c.idsByTarget[target] {
		c.detailsById[id].conn.Close()
		delete(c.detailsById, id)
	}
	delete(c.idsByTarget, target)
}

// Targets removes all currently active backend target names.
func (c *connectionPool) Targets() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	res := make([]string, 0, len(c.idsByTarget))
	for target, _ := range c.idsByTarget {
		res = append(res, target)
	}
	sort.Strings(res)
	return res
}
