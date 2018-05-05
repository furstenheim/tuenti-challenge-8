package main

import (
	"container/heap"
	"log"
	"strings"
	"strconv"
)

type growingMatching struct {
	startString, endString string
	// includedParts map[int]bool
	remainingParts map[int]bool
}

type Case struct {
	parts []string
}

func solveCase (c Case) string {
	pq := make(PriorityQueue, 0, len(c.parts))
	heap.Init(&pq)
	for i, p := range(c.parts) {
		for j := 0; j < len(p) + 1; j++ {
			gm := growingMatching{
				endString: p[0:j],
				startString:p[j:],
				remainingParts: map[int]bool{},
			}
			// everything except this part is remaining
			for j := 0; j < len(c.parts); j++ {
				if j != i {
					gm.remainingParts[j] = true
				}
			}
			pqi := &pqItem{
				match: gm,
				t: len(gm.startString) + len(gm.endString),
			}
			heap.Push(&pq, pqi)
			// log.Println(gm)
		}
	}


	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		match := item.match

		if len(match.startString) == 0 && len(match.endString) == 0 {
			success := make([]string, 0)
			for i, _ := range (c.parts) {
				if _, ok := match.remainingParts[i]; !ok {
					success = append(success, strconv.Itoa(i + 1))
				}
			}
			log.Println("Successss", success)
			return strings.Join(success, ",")
		}
		for i, ok := range(match.remainingParts) {
			if !ok {
				log.Fatal("There should only be trues in the map", i, match)
			}
			p := c.parts[i]
			if len(match.startString) != 0 && strings.HasPrefix(p, match.startString[0: min(len(p), len(match.startString))]) {
				var newStartString string
				if len(p) >= len(match.startString) {
					newStartString = p[len(match.startString):]
				} else {
					newStartString = match.startString[len(p):]
				}
				increasedMatching := growingMatching{
					endString: match.endString,
					// we need to match the end of it
					startString: newStartString,
					remainingParts: map[int]bool{},
				}
				for j, _ := range(match.remainingParts) {
					if j != i {
						increasedMatching.remainingParts[j] = true
					}
				}
				pqi := &pqItem{
					match: increasedMatching,
					t: len(newStartString) + len(match.endString),
				}
				heap.Push(&pq, pqi)
			}
			if len(match.endString) != 0 && strings.HasSuffix(p, match.endString[len(match.endString) - min(len(p), len(match.endString)):]) {
				var newEndString string
				if len(p) >= len(match.endString) {
					newEndString = p[0: len(p) - len(match.endString)]
				} else {
					newEndString = match.endString[0: len(match.endString) - len(p)]
				}
				increasedMatching := growingMatching{
					startString: match.startString,
					// we need to match the end of it
					endString: newEndString,
					remainingParts: map[int]bool{},
				}
				for j, _ := range(match.remainingParts) {
					if j != i {
						increasedMatching.remainingParts[j] = true
					}
				}
				pqi := &pqItem{
					match: increasedMatching,
					t: len(newEndString) + len(match.startString),
				}
				heap.Push(&pq, pqi)
			}

		}
	}
	log.Fatal("Could not find match")
	return ""
}

func min (i, j int) int {
	if i < j {
		return i
	}
	return j
}