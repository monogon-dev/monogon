package toposort

import "testing"

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
		if err == ErrCycle {
			// Cycle found
			return
		}
		if err != nil {
			t.Error(err)
		}
		validateSolution(t, g, solution)
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
