package toposort

import (
	"cmp"
	"errors"
	"slices"
	"testing"
)

func TestBasic(t *testing.T) {
	// Test the example from Wikipedia's Topological sorting page
	var g Graph[int]
	g.AddEdge(5, 11)
	g.AddEdge(11, 2)
	g.AddEdge(7, 11)
	g.AddEdge(7, 8)
	g.AddEdge(11, 9)
	g.AddEdge(8, 9)
	g.AddEdge(11, 10)
	g.AddEdge(3, 8)
	g.AddEdge(3, 10)

	solution, err := g.TopologicalOrder()
	if err != nil {
		t.Fatal(err)
	}
	validateSolution(t, g, solution)

	detSolution, err := g.DetTopologicalOrder(cmp.Compare)
	if err != nil {
		t.Fatal(err)
	}
	validateSolution(t, g, detSolution)
}

func TestImplicitNodesEmpty(t *testing.T) {
	var g Graph[int]
	g.AddNode(5)
	g.AddNode(10)
	g.AddEdge(5, 10)

	out := g.ImplicitNodeReferences()
	if len(out) != 0 {
		t.Errorf("expected no implicit nodes, got %d", len(out))
	}
}

func TestImplicitNodesForward(t *testing.T) {
	var g Graph[int]
	g.AddNode(5)
	g.AddEdge(5, 10)

	out := g.ImplicitNodeReferences()
	if len(out) != 1 {
		t.Errorf("expected 1 implicit node, got %d", len(out))
	}
	if len(out[10]) != 1 {
		t.Error("expected node 10 to be implicit")
	}
	if !out[10][5] {
		t.Errorf("expected node 10 to be referenced by node 5")
	}
}

func TestImplicitNodesBackwards(t *testing.T) {
	var g Graph[int]
	g.AddNode(10)
	g.AddEdge(5, 10)

	out := g.ImplicitNodeReferences()
	if len(out) != 1 {
		t.Errorf("expected 1 implicit node, got %d", len(out))
	}
	if len(out[5]) != 1 {
		t.Error("expected node 5 to be implicit")
	}
	if !out[5][10] {
		t.Errorf("expected node 5 to be referenced by node 10")
	}
}

func TestTopoSortDet(t *testing.T) {
	var g Graph[int]
	g.AddEdge(5, 11)
	g.AddEdge(11, 2)
	g.AddEdge(7, 11)
	g.AddEdge(7, 8)
	g.AddEdge(11, 9)
	g.AddEdge(8, 9)
	g.AddEdge(11, 10)
	g.AddEdge(3, 8)
	g.AddEdge(3, 10)

	firstSolution, err := g.DetTopologicalOrder(cmp.Compare)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		sol, err := g.DetTopologicalOrder(cmp.Compare)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(firstSolution, sol) {
			t.Fatalf("solution in iteration %v (%v) is not equal to first solution (%v)", i, firstSolution, sol)
		}
	}
}

// Fuzzer can be run with
// bazel test //go/algorithm/toposort:toposort_test
//   --test_arg=-test.fuzz=FuzzTopoSort
//   --test_arg=-test.fuzzcachedir=/tmp/fuzz
//   --test_arg=-test.fuzztime=60s

func FuzzTopoSort(f *testing.F) {
	// Add starting example from above
	f.Add([]byte{5, 11, 11, 2, 7, 11, 7, 8, 11, 9, 8, 9, 11, 10, 3, 8, 3, 10})
	f.Fuzz(func(t *testing.T, a []byte) {
		var g Graph[int]
		for i := 0; i < len(a)-1; i += 2 {
			g.AddEdge(int(a[i]), int(a[i+1]))
		}
		solution, err := g.TopologicalOrder()
		if errors.Is(err, ErrCycle) {
			// Cycle found
			return
		}
		if err != nil {
			t.Error(err)
		}
		validateSolution(t, g, solution)

		detSolution, err := g.DetTopologicalOrder(cmp.Compare)
		if errors.Is(err, ErrCycle) {
			// Cycle found
			return
		}
		if err != nil {
			t.Error(err)
		}
		validateSolution(t, g, detSolution)
	})
}

func validateSolution(t *testing.T, g Graph[int], solution []int) {
	t.Helper()
	visited := make(map[int]bool)
	for _, n := range solution {
		visited[n] = true
		for m := range g.nodes[n] {
			if !visited[m] {
				t.Errorf("node %d has an edge to node %d, but not yet visited", n, m)
			}
		}
	}
}
