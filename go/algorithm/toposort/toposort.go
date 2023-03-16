package toposort

import "errors"

type outEdges[Node comparable] map[Node]bool

// Graph is a directed graph represented using an adjacency list.
// Its nodes have type T. An empty Graph object is ready to use.
type Graph[T comparable] struct {
	nodes map[T]outEdges[T]
}

func (g *Graph[T]) ensureNodes() {
	if g.nodes == nil {
		g.nodes = make(map[T]outEdges[T])
	}
}

// AddNode adds the node n to the graph if it does not already exist, in which
// case it does nothing.
func (g *Graph[T]) AddNode(n T) {
	g.ensureNodes()
	if _, ok := g.nodes[n]; !ok {
		g.nodes[n] = make(outEdges[T])
	}
}

// AddEdge adds the directed edge from the from node to the to node to the
// graph if it does not already exist, in which case it does nothing.
func (g *Graph[T]) AddEdge(from, to T) {
	g.ensureNodes()
	g.AddNode(from)
	g.AddNode(to)
	g.nodes[from][to] = true
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
