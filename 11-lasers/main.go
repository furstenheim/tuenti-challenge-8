package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
	"fmt"
)


const INF = 1000
func main () {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	handleError(err)
	firstLineFields := strings.Fields(line)
	numberOfCases, err := strconv.Atoi(firstLineFields[0])
	if (err != nil) {
		log.Fatal(err)
	}
	for i := 0; i < numberOfCases; i ++ {
		c := parseCase(reader)
		solveCase(c, i)
	}
}

type Case struct {
	N, M, I int
	matrix map[int]map[int]bool
	matchLeft, matchRight, distances map[int]int
}

func parseCase (reader *bufio.Reader) Case {
	line, err := reader.ReadString('\n')
	handleError(err)
	fields := strings.Fields(line)
	N, err := strconv.Atoi(fields[0])
	handleError(err)
	M, err := strconv.Atoi(fields[1])
	handleError(err)
	I, err := strconv.Atoi(fields[2])
	handleError(err)

	c := Case{M: M, N: N, I: I, matrix: make(map[int]map[int]bool), matchLeft: make(map[int]int), matchRight: make(map[int]int), distances: make(map[int]int)}
	for i := 0; i < N; i++ {
		// we follow hopcroft convention 1 --- n
		c.matrix[i + 1] = make(map[int]bool)
		c.matchLeft[i + 1] = 0
		c.matchRight[i + 1] = 0
	}
	for i := 0; i < I; i++ {
		laserLine, err := reader.ReadString('\n')
		handleError(err)
		fields := strings.Fields(laserLine)
		A, err := strconv.Atoi(fields[0])
		handleError(err)
		B, err := strconv.Atoi(fields[1])
		c.matrix[A + 1][B + 1] = true
	}
	return c
}

func solveCase (c Case, i int) {
	flow := c.hopcrof_karp()
	log.Println(i+1, flow, c.I)
	fmt.Printf("Case #%d: %d\n", i + 1, c.M + c.N - flow)
}


func handleError (err error){
	if err != nil {
		log.Fatal(err)
	}
}