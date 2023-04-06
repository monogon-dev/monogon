package toposort

import "errors"

type outEdges[Node comparable] map[Node]bool

// Graph is a directed graph represented using an adjacency list.
// Its nodes have type T. An empty Graph object is ready to use.
type Graph[T comparable] struct {
	nodes map[T]outEdges[T]
	// Set of nodes added explicitly, used to check for references to nodes
	// which haven't been added explicitly.
	addedNodes map[T]bool
}

func (g *Graph[T]) ensureNodes() {
	if g.nodes == nil {
		g.nodes = make(map[T]outEdges[T])
	}
	if g.addedNodes == nil {
		g.addedNodes = make(map[T]bool)
	}
}

// AddNode adds the node n to the graph if it does not already exist, in which
// case it does nothing.
func (g *Graph[T]) AddNode(n T) {
	g.ensureNodes()
	g.addedNodes[n] = true
	g.addNodeImplicit(n)
}

// addNodeImplicit adds the node n to the graph if it does not already exist
// without marking it as being explicitly added and without ensuring that
// g.nodes is populated.
func (g *Graph[T]) addNodeImplicit(n T) {
	if _, ok := g.nodes[n]; !ok {
		g.nodes[n] = make(outEdges[T])
	}
}

// AddEdge adds the directed edge from the from node to the to node to the
// graph if it does not already exist, in which case it does nothing.
// If nodes from and/or to do not already exist, they are added implicitly.
func (g *Graph[T]) AddEdge(from, to T) {
	g.ensureNodes()
	g.addNodeImplicit(from)
	g.addNodeImplicit(to)
	g.nodes[from][to] = true
}

// ImplicitNodeReferences returns a map of nodes with the set of their incoming
// and outgoing references for all nodes which were referenced in an AddEdge
// call but not added via an explicit AddNode. If the length of the returned map
// is zero, all referenced nodes were explicitly added. This can be used to
// check for bad references in dependency trees.
func (g *Graph[T]) ImplicitNodeReferences() map[T]map[T]bool {
	out := make(map[T]map[T]bool)
	for n, outEdges := range g.nodes {
		if !g.addedNodes[n] {
			if out[n] == nil {
				out[n] = make(map[T]bool)
			}
			// Add all outgoing edges, the refernces from these must come from
			// reverse dependencies as this node was never added.
			for e := range outEdges {
				out[n][e] = true
			}
		}
		for e := range outEdges {
			// Add incoming edges to unreferenced nodes.
			if !g.addedNodes[e] {
				if out[e] == nil {
					out[e] = make(map[T]bool)
				}
				out[e][n] = true
			}
		}
	}
	return out
}

var ErrCycle = errors.New("graph has at least one cycle")

// TopologicalOrder sorts the nodes in the graph according to their topological
// order (aka. dependency order) and returns a valid order. Note that this
// implementation returns the reverse order of a "standard" topological order,
// i.e. if edge A -> B exists, it returns B, A. This is more convenient for
// solving dependency ordering problems, as a normal dependency graph will sort
// according to execution order.
// Further note that there can be multiple such orderings and any of them might
// be returned. Even multiple invocations of TopologicalOrder can return
// different orderings. If no such ordering exists (i.e. the graph contains at
// least one cycle), an error is returned.
func (g *Graph[T]) TopologicalOrder() ([]T, error) {
	g.ensureNodes()
	// Depth-first search-based implementation with O(|n| + |E|) complexity.
	var solution []T
	unmarkedNodes := make(map[T]bool)
	for n := range g.nodes {
		unmarkedNodes[n] = true
	}

	nodeVisitStack := make(map[T]bool)

	var visit func(n T) error
	visit = func(n T) error {
		if !unmarkedNodes[n] {
			return nil
		}
		// If we're revisiting a node already on the visit stack, we have a
		// cycle.
		if nodeVisitStack[n] {
			return ErrCycle
		}
		nodeVisitStack[n] = true
		// Visit every node m pointed to by an edge from the current node n.
		for m := range g.nodes[n] {
			if err := visit(m); err != nil {
				return err
			}
		}
		delete(nodeVisitStack, n)
		delete(unmarkedNodes, n)
		solution = append(solution, n)
		return nil
	}

	for n := range unmarkedNodes {
		if err := visit(n); err != nil {
			return nil, err
		}
	}
	return solution, nil
}
