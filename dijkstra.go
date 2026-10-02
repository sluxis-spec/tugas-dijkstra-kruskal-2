// Dijkstra's Algorithm
// Finds the shortest path from a source vertex to all other vertices.
// Graph: map where graph[u] = []Edge{{To: v, Weight: w}}
// Time: O((V + E) log V)
package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a weighted edge from a node to To.
type Edge struct {
	To     string
	Weight int
}

// item is a priority-queue entry.
type item struct {
	node string
	dist int
}

// priorityQueue implements heap.Interface (min-heap by dist).
type priorityQueue []item

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x any)         { *pq = append(*pq, x.(item)) }
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	return it
}

// Dijkstra returns the shortest distances from start and the predecessor map.
func Dijkstra(graph map[string][]Edge, start string) (map[string]int, map[string]string) {
	dist := make(map[string]int, len(graph))
	prev := make(map[string]string, len(graph))
	for n := range graph {
		dist[n] = math.MaxInt
	}
	dist[start] = 0

	pq := &priorityQueue{{node: start, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		it := heap.Pop(pq).(item)
		if it.dist > dist[it.node] { // stale entry
			continue
		}
		for _, e := range graph[it.node] {
			if nd := it.dist + e.Weight; nd < dist[e.To] {
				dist[e.To] = nd
				prev[e.To] = it.node
				heap.Push(pq, item{node: e.To, dist: nd})
			}
		}
	}
	return dist, prev
}

// ReconstructPath builds the shortest path from start to end using prev.
// Returns nil if end is unreachable from start.
func ReconstructPath(prev map[string]string, start, end string) []string {
	var path []string
	for cur := end; cur != ""; cur = prev[cur] {
		path = append([]string{cur}, path...)
		if cur == start {
			return path
		}
	}
	return nil
}

func main() {
	graph := map[string][]Edge{
		"A": {{"B", 4}, {"C", 2}},
		"B": {{"A", 4}, {"C", 1}, {"D", 5}},
		"C": {{"A", 2}, {"B", 1}, {"D", 8}},
		"D": {{"B", 5}, {"C", 8}, {"E", 2}},
		"E": {{"D", 2}},
	}

	dist, prev := Dijkstra(graph, "A")
	fmt.Println("Distances from A:", dist)
	fmt.Println("Shortest path A -> E:", ReconstructPath(prev, "A", "E"))
}