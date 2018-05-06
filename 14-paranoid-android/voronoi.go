package main

import (
	"github.com/pradeep-pyro/triangle"
	"github.com/furstenheim/SimpleRTree"
	"math"
	"log"
	"container/heap"
)

const START_INDEX = -2
const END_INDEX = -1
type Algorithm struct {
	vertices, rayDirections [][2]float64
	edges [][2]int32
	rayOrigins []int32
	c Case
	centersTree, verticesTree *SimpleRTree.SimpleRTree
}

type AlgorithmNode struct {
	id int
	distance float64
}

type Graph struct {
	nPoints int
	distances map[int]map[int]float64 // start is -1 end is -2
}

func solveGraph (g Graph) (float64, bool) {
	currentBests := make(map[int]float64)
	for i := START_INDEX; i < g.nPoints; i++ {
		currentBests[i] = math.MaxFloat64
	}

	pq := make(PriorityQueue, 0, g.nPoints)
	heap.Init(&pq)

	pqi := &pqItem{
		distance: 0,
		id: START_INDEX,
	}
	heap.Push(&pq, pqi)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		if item.id == END_INDEX {
			return item.distance, true
		}
		if currentBest, _ := currentBests[item.id]; currentBest <= item.distance {
			continue
		}

		currentBests[item.id] = item.distance

		for v, d := range(g.distances[item.id]) {
			nextPq := &pqItem{
				distance: item.distance + d,
				id: v,
			}
			heap.Push(&pq, nextPq)
		}
	}
	return -1, false

}



func voronoi2Graph (a Algorithm) Graph {
	g := Graph{nPoints: len(a.vertices), distances: make(map[int]map[int]float64)}
	for i := START_INDEX; i < len(a.vertices); i++ {
		g.distances[i] = make(map[int]float64)
	}
	for _, e := range(a.edges) {
		// we need to remove those that are closer than the allowed distance
		// for that we need to find the closest point, we use the index
		v0 := a.vertices[int(e[0])]
		v1 := a.vertices[int(e[1])]
		isAllowed := a.isSideAvailable(v0[0], v0[1], v1[0], v1[1])
		// log.Println(d, allowedDistance, i, d >= allowedDistance, cx, cy, v0, v1)
		if isAllowed {
			g.distances[int(e[0])][int(e[1])] = SegmentDistance(v0[0], v0[1], v1[0], v1[1])
			g.distances[int(e[1])][int(e[0])] = SegmentDistance(v0[0], v0[1], v1[0], v1[1])
		}
	}

	xStart := a.c.xStart.approxFloat()
	yStart := a.c.yStart.approxFloat()
	// find where are start and end located
	sx, sy, _, ok := a.verticesTree.FindNearestPoint(xStart, yStart)
	if !ok {
		log.Fatal("Could not find closest point")
	}
	xEnd := a.c.xEnd.approxFloat()
	yEnd := a.c.yEnd.approxFloat()

	// very degenerate case
	if xStart == xEnd && yStart == yEnd {
		g.distances[START_INDEX][END_INDEX] = 0
	}

	ex, ey, _, ok := a.verticesTree.FindNearestPoint(xEnd, yEnd)
	if !ok {
		log.Fatal("Could not find closest point")
	}
	si := -1
	ei := -1

	for i, v := range(a.vertices) {
		if v[0] == sx && v[1] == sy {
			si = i
		}
		if v[0] == ex && v[1] == ey {
			ei = i
		}
	}

	// avoid degenerancy cases
	if ex == sx && ey == sy {
		log.Println("Sharing closest point")
		ei = si
	}
	if ei == -1 || si == -1 {
		log.Fatal("could not find initial point")
	}


	// Compute distance to vertices. We should actually find the edge where start and end are located. but that is harder

	var foundStart, foundEnd bool
	for _, e := range(a.edges) {

		v0 := a.vertices[int(e[0])]
		v1 := a.vertices[int(e[1])]
		if !(v0[0] != v1[0] || v0[1] != v1[1] /** edge is not degenrated */) {
			continue
		}
		// if ((sx == v0[0] && sy == v0[1]) || (sx == v1[0] && sy == v1[1])) &&
		if isBetween(v0[0], v0[1], v1[0], v1[1], xStart, yStart){
			foundStart = true
			if a.isSideAvailable(xStart, yStart, v1[0], v1[1]) {
				g.distances[START_INDEX][int(e[1])] = SegmentDistance(v1[0], v1[1], xStart, yStart)
			}
			if a.isSideAvailable(xStart, yStart, v0[0], v0[1]) {
				g.distances[START_INDEX][int(e[0])] = SegmentDistance(v0[0], v0[1], xStart, yStart)
			}
		}

		if isBetween(v0[0], v0[1], v1[0], v1[1], xEnd, yEnd) {
			if a.isSideAvailable(xEnd, yEnd, v1[0], v1[1]) {
				g.distances[int(e[1])][END_INDEX] = SegmentDistance(v1[0], v1[1], xEnd, yEnd)
			}
			if a.isSideAvailable(xEnd, yEnd, v0[0], v0[1]) {
				g.distances[int(e[0])][END_INDEX] = SegmentDistance(v0[0], v0[1], xEnd, yEnd)
			}
			foundEnd = true
		}


		if isBetween(v0[0], v0[1], v1[0], v1[1], xStart, yStart) && isBetween(v0[0], v0[1], v1[0], v1[1], xEnd, yEnd) {
			if a.isSideAvailable(xStart, yStart, xEnd, yEnd) {
				g.distances[START_INDEX][END_INDEX] = SegmentDistance(xStart, yStart, xEnd, yEnd)
			}
		}


	}


	// That means they are rays
	if !foundStart && !foundEnd && (isBetween(sx, sy, xStart, yStart, xEnd, yEnd) || isBetween(sx, sy, xEnd, yEnd, xStart, yStart)) {
		if a.isSideAvailable(xStart, yStart, xEnd, yEnd) {
			g.distances[START_INDEX][END_INDEX] = SegmentDistance(xStart, yStart, xEnd, yEnd)
		}
	}
	if !foundStart {
		for _, r := range (a.rayOrigins) {
			v0 := a.vertices[int(r)]
			if v0[0] == sx && v0[1] == sy {
				if a.isSideAvailable(xStart, yStart, v0[0], v0[1]) {
					g.distances[START_INDEX][int(r)] = SegmentDistance(v0[0], v0[1], xStart, yStart)
				}
				foundStart = true
			}
		}
	}
	if !foundEnd {
		for _, r := range (a.rayOrigins) {
			v0 := a.vertices[int(r)]
			if v0[0] == ex && v0[1] == ey {
				if a.isSideAvailable(xEnd, yEnd, v0[0], v0[1]) {
					g.distances[int(r)][END_INDEX] = SegmentDistance(v0[0], v0[1], xEnd, yEnd)
				}
				foundEnd = true
			}
		}
	}

	if !foundEnd || !foundStart {
		log.Fatal("Could not find nex", ei, si, foundStart, foundStart, a.edges, a.rayOrigins)
	}
	return g
}

func (a * Algorithm) isSideAvailable (x0, y0, x1, y1 float64) bool {
	allowedDistance := a.c.rAvoidance.approxFloat()
	mx := (x0 + x1) / 2
	my := (y0 + y1) / 2
	cx, cy, _, ok := a.centersTree.FindNearestPoint(mx, my)
	if !ok {
		log.Fatal("Could not find closest point")
	}
	d := DistanceFromSegmentToPoint(x0, y0, x1, y1, cx, cy)
	// for rays closest point might be the vertex
	// d1 := SegmentDistance(cx, cy, x0, y0)
	// d2 := SegmentDistance(cx, cy, x1, y1)
	// log.Println(d, allowedDistance, i, d >= allowedDistance, cx, cy, v0, v1)
	if d >= allowedDistance {
		return true
	}
	return false
}

func computeVoronoi (c Case) Algorithm {
	vertices, edges, rayOrigins, rayDirections := triangle.Voronoi(c.points)
	centersTree := SimpleRTree.New().Load(SimpleRTree.FlatPoints(c.flatPoints))
	flatVertices := make([]float64, len(vertices) * 2)
	for i, v := range(vertices){
		flatVertices[2 * i] = v[0]
		flatVertices[2 * i + 1] = v[1]
	}

	verticesTree := SimpleRTree.New().Load(SimpleRTree.FlatPoints(flatVertices))

	return Algorithm{c: c, vertices: vertices, edges: edges, rayOrigins: rayOrigins, rayDirections: rayDirections, verticesTree: verticesTree, centersTree:centersTree}
}

func min (x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func max (x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func (f exactFloat) approxFloat () float64 {
	return float64(f.dividend) / float64(f.divisor)
}

func SegmentDistance (x1, y1, x2, y2 float64) float64 {
	x := x1
	y := y1
	dx := x2 - x
	dy := y2 - y

	return math.Sqrt(dx*dx + dy*dy)
}

func DistanceFromSegmentToPoint(x1, y1, x2, y2, px, py float64) float64 {
	// code from go geo
	x := x1
	y := y1
	dx := x2 - x
	dy := y2 - y

	if dx != 0 || dy != 0 {
		t := ((px-x)*dx + (py-y)*dy) / (dx*dx + dy*dy)

		if t > 1 {
			x = x2
			y = y2
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}

	dx = px - x
	dy = py - y

	return math.Sqrt(dx*dx + dy*dy)
}

func isBetween (x0, y0, x1, y1, px, py float64) bool {
	crossproduct := (py - y0) * (x1 - x0) - (px - x0) * (y1 - y0)
	if math.Abs(crossproduct) > 0.00001 {
		return false
	}
	dotproduct := (px - x0) * (x1 - x0) + (py - y0)*(y1 - y0)
	if dotproduct < 0 {
		return false
	}
	squaredLength := (x1 - x0)*(x1 - x0) + (y1 - y0)*(y1 - y0)
	if dotproduct > squaredLength {
		return false
	}
	return true
}
