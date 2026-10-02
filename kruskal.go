// Kruskal's Algorithm
// Finds a Minimum Spanning Tree (MST) of a weighted undirected graph.
// Time: O(E log E)
package main

import (
	"fmt"
	"sort"
)

// Edge represents a weighted undirected edge.
type Edge struct {
	U, V   string
	Weight int
}

// DSU is a disjoint-set (union-find) with path compression + union by rank.
type DSU struct {
	parent map[string]string
	rank   map[string]int
}

func NewDSU(nodes []string) *DSU {
	d := &DSU{parent: make(map[string]string, len(nodes)), rank: make(map[string]int, len(nodes))}
	for _, n := range nodes {
		d.parent[n] = n
	}
	return d
}

func (d *DSU) Find(x string) string {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]] // path compression
		x = d.parent[x]
	}
	return x
}

func (d *DSU) Union(x, y string) bool {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return false
	}
	if d.rank[px] < d.rank[py] { // union by rank
		px, py = py, px
	}
	d.parent[py] = px
	if d.rank[px] == d.rank[py] {
		d.rank[px]++
	}
	return true
}

// Kruskal returns the MST edges and the total weight.
func Kruskal(nodes []string, edges []Edge) ([]Edge, int) {
	sort.Slice(edges, func(i, j int) bool { return edges[i].Weight < edges[j].Weight })

	d := NewDSU(nodes)
	var mst []Edge
	total := 0
	for _, e := range edges {
		if d.Union(e.U, e.V) { // adding it won't create a cycle
			mst = append(mst, e)
			total += e.Weight
			if len(mst) == len(nodes)-1 { // MST complete
				break
			}
		}
	}
	return mst, total
}

func main() {
	nodes := []string{"A", "B", "C", "D", "E"}
	edges := []Edge{
		{"A", "B", 4}, {"A", "C", 2}, {"B", "C", 1},
		{"B", "D", 5}, {"C", "D", 8}, {"D", "E", 2},
	}

	mst, total := Kruskal(nodes, edges)
	fmt.Println("MST edges:", mst)
	fmt.Println("Total weight:", total)
}