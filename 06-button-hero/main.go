package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
	"sort"
	"container/heap"
	"fmt"
	"math"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	firstLineFields := strings.Fields(line)
	numberOfCases, err := strconv.Atoi(firstLineFields[0])
	if (err != nil) {
		log.Fatal(err)
	}
	for i := 0; i < numberOfCases; i ++ {
		c := parseCase(reader, i)

		solveCase(c, i)
	}
}

type Case struct {
	//maxT uint32
	notes []Note
	matrix map[Note][]int
}

type Note struct {
	starT, endT, points uint32
}

func parseCase (reader *bufio.Reader, caseIndex int) Case {
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)
	nNotes, err := strconv.Atoi(fields[0])
	handleError(err)
	rawNotes := make([]Note, nNotes)
	for i := 0; i < nNotes; i++ {
		line, _ := reader.ReadString('\n')
		fieldsNote := strings.Fields(line)
		X, err:= strconv.Atoi(fieldsNote[0])
		handleError(err)
		L, err:= strconv.Atoi(fieldsNote[1])
		handleError(err)
		S, err:= strconv.Atoi(fieldsNote[2])
		handleError(err)
		P, err:= strconv.Atoi(fieldsNote[3])
		handleError(err)

		rawNotes[i] = Note{
			starT: uint32(X / S),
			endT: uint32(X / S + L/S),
			points: uint32(P),
		}
	}

	ns := NotesSorter(rawNotes)
	sort.Sort(ns)

	simplifiedNotes := make([]Note, 0, nNotes + 1)
	// append dummy initial value
	previousNote := Note{}

	for i := 0; i < nNotes; i++ {
		nextNote := rawNotes[i]
		// simplify
		if nextNote.starT == previousNote.starT && nextNote.endT == previousNote.endT {
			previousNote.points += nextNote.points
		} else {
			simplifiedNotes = append(simplifiedNotes, previousNote)
			previousNote = nextNote
		}
	}
	simplifiedNotes = append(simplifiedNotes, previousNote)

	log.Println("simplified notes has length ", len(simplifiedNotes))
  	matrix := make(map[Note][]int)
	// build matrix
	for i, n := range simplifiedNotes {
		possibleNextNotes := make([]int, 0)
		isCurrentMaxSet := false
		var currentMax uint32
		currentMax = math.MaxUint32

		for j := i + 1; j < len(simplifiedNotes); j++ {
			possibleNote := simplifiedNotes[j]

			if isCurrentMaxSet && possibleNote.starT > currentMax {
				break
			}

			if possibleNote.starT > n.endT {
				isCurrentMaxSet = true
				currentMax = min(currentMax, possibleNote.endT)
				possibleNextNotes = append(possibleNextNotes, j)
			}
		}
		matrix[n] = possibleNextNotes
	}

  	log.Println("matrix has size", len(matrix))
	c := Case{
		notes: simplifiedNotes,
		matrix: matrix,
	}
	return c
}


func solveCase (c Case, i int) {
	log.Println("-----------", i)

/*	log.Println(c.notes[0])
	log.Println(c.matrix[c.notes[0]])
	log.Println(c.notes[0: 2])
	log.Println(c.notes[689])*/
	currentBests := make(map[uint32]uint32)
	pq := PriorityQueue{arr: make([]*pqItem, 0), c: c}
	heap.Init(&pq)
	pqi := &pqItem{
		highestNoteIndex: 0,
		totalsPoints: 0,
	}
	heap.Push(&pq, pqi)
	var allTimeMax uint32
	allTimeMax = 0

  	for pq.Len() > 0 {
  		item := heap.Pop(&pq).(*pqItem)
		note := c.notes[item.highestNoteIndex]
  		if currentMax, ok := currentBests[note.endT]; ok && currentMax > item.totalsPoints {
			// we have done better before, the show must go on
  			continue
		}
		allTimeMax = max(allTimeMax, item.totalsPoints)
    		currentBests[note.endT] = item.totalsPoints

		for _, k := range(c.matrix[note]) {
			nextNote := c.notes[k]
			nextItem := &pqItem{
				highestNoteIndex: k,
				totalsPoints: item.totalsPoints + nextNote.points,
			}
			heap.Push(&pq, nextItem)
		}

	}
	fmt.Printf("Case #%d: %d\n", i +1, allTimeMax)
}


func handleError (err error){
	if err != nil {
		log.Fatal(err)
	}
}

func max (i, j uint32) uint32 {
	if i >= j {
		return i
	}
	return j
}
func min (i, j uint32) uint32 {
	if i <= j {
		return i
	}
	return j
}
