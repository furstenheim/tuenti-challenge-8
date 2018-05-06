package main

// Inspired in https://somemoreacademic.blogspot.com.es/2012/11/maximum-matching-in-bipartite-graph_10.html


import (
	"container/heap"
)

func (c * Case) hopcrof_karp () int {
	nMatch := 0
	for c.bfs() {
		for i := 1; i <= c.N; i++ {
			if c.matchLeft[i] == 0 && c.dfs(i){
				nMatch++
			}
		}
	}
	return nMatch
}

func (c * Case) bfs () bool {
	pq := make(PriorityQueue, 0, c.N)
	heap.Init(&pq)
	for m1, m2 := range(c.matchLeft) {
		if m2 == 0 {
			c.distances[m1] = 0
			heap.Push(&pq, m1)
		} else {
			c.distances[m1] = INF
		}
	}
	c.distances[0] = INF
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(int)
		if c.distances[item] < c.distances[0] {
			for v, _ := range(c.matrix[item]) {
				if c.distances[c.matchRight[v]] == INF {
					c.distances[c.matchRight[v]] = c.distances[item] + 1
					heap.Push(&pq, c.matchRight[v])
				}
			}
		}
	}
	return c.distances[0] != INF
}

func (c * Case) dfs (u int) bool {
	if u != 0 {
		for v, _ := range(c.matrix[u]) {
			if c.distances[c.matchRight[v]] == c.distances[u] + 1 {
				if c.dfs(c.matchRight[v]) {
					c.matchRight[v] = u
					c.matchLeft[u] = v
					return true
				}
			}
		}
		c.distances[u] = INF
		return false
	}
	return true
}
