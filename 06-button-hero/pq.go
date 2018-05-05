package main

type pqItem struct {
	highestNoteIndex int
	totalsPoints uint32
}

type PriorityQueue struct {
	arr []*pqItem
	c Case
}

func (pq PriorityQueue) Len() int { return len(pq.arr) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.c.notes[pq.arr[i].highestNoteIndex].starT == pq.c.notes[pq.arr[j].highestNoteIndex].starT {
		return pq.c.notes[pq.arr[i].highestNoteIndex].points > pq.c.notes[pq.arr[j].highestNoteIndex].points
	}
	return pq.c.notes[pq.arr[i].highestNoteIndex].starT < pq.c.notes[pq.arr[j].highestNoteIndex].starT
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*pqItem)
	(*pq).arr = append((*pq).arr, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := (*pq).arr
	n := len(old)
	item := old[n-1]
	(*pq).arr = old[0 : n-1]
	return item
}

