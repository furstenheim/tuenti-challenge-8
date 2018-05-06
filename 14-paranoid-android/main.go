package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
	"fmt"
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
	nPoints int
	points [][2]float64
	flatPoints []float64
	xStart, xEnd, yStart, yEnd, rAvoidance exactFloat
}

type exactFloat struct {
	dividend, divisor int
}

const I_TO_LOG = -1


func parseCase (reader *bufio.Reader, n int) Case {
	line, err := reader.ReadString('\n')
	if n == I_TO_LOG {
		fmt.Print(line)
	}
	handleError(err)
	fields := strings.Fields(line)
	N, err := strconv.Atoi(fields[0])
	handleError(err)
	c := Case{nPoints: N, points: make([][2]float64, 0, N), flatPoints: make([]float64, 0, 2 * N)}
	for i:= 0; i < N; i++ {
		line, err := reader.ReadString('\n')
		if n == I_TO_LOG {
			fmt.Print(line)
		}
		handleError(err)
		fields := strings.Fields(line)
		x, err := strconv.Atoi(fields[0])
		handleError(err)
		y, err := strconv.Atoi(fields[1])
		handleError(err)
		c.points = append(c.points, [2]float64{float64(x), float64(y)})
		c.flatPoints = append(c.flatPoints, float64(x), float64(y))
	}
	line, err = reader.ReadString('\n')
	if n == I_TO_LOG {
		fmt.Print(line)
	}
	handleError(err)
	radius := toExactFloat(line)
	lineStart, err := reader.ReadString('\n')
	if n == I_TO_LOG {
		fmt.Print(lineStart)
	}
	handleError(err)
	coordinatesStart := strings.Fields(lineStart)
	lineEnd, err := reader.ReadString('\n')
	if n == I_TO_LOG {
		fmt.Print(lineEnd)
	}
	handleError(err)
	coordinatesEnd := strings.Fields(lineEnd)
	c.rAvoidance = radius
	c.xStart = toExactFloat(coordinatesStart[0])
	c.yStart = toExactFloat(coordinatesStart[1])
	c.xEnd = toExactFloat(coordinatesEnd[0])
	c.yEnd = toExactFloat(coordinatesEnd[1])
	return c
}

func solveCase (c Case, i int) {
	log.Println("-----------", i)
	a := computeVoronoi(c)
	// log.Println(a.vertices, "edges", a.edges, "rayorigins", a.rayOrigins, a.rayDirections)
	g := voronoi2Graph(a)
	// log.Println(g)
	distance, ok := solveGraph(g)
	if !ok {
		fmt.Printf("Case #%d: IMPOSSIBLE\n", i + 1)
	} else {
		fmt.Printf("Case #%d: %.3f\n", i + 1, distance)
	}
}


func handleError (err error){
	if err != nil {
		log.Fatal(err)
	}
}

func toExactFloat (s string) exactFloat {
	fields := strings.Split(strings.Fields(s)[0], "/")
	div, err := strconv.Atoi(fields[0])
	handleError(err)
	divisor, err := strconv.Atoi(fields[1])
	handleError(err)
	return exactFloat{dividend: div, divisor: divisor}
}