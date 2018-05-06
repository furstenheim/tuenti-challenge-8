package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
	"fmt"
)

const MODULO = 1000000007
var cachedValues = make(map[int]int)
var cachedCouples = make(map[Couple]int)
var cachedMagicConstants = make(map[int]int)

type Couple [2]int


func main () {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
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
	n int
}

func parseCase (reader *bufio.Reader) Case {
	line, err := reader.ReadString('\n')
	handleError(err)
	fields := strings.Fields(line)
	N, err := strconv.Atoi(fields[0])
	handleError(err)
	c := Case{n: N}
	return c
}

func solveCase (c Case, i int) {
	v := solveI(c.n)
	log.Println("---------", i)
	fmt.Printf("Case #%d: %d\n", i + 1, v)

}

func init () {
	cachedMagicConstants[1] = 0
}
func computeMagicConstant (n int) int {
	if v, ok := cachedMagicConstants[n]; ok {
		return v
	}
	total := computeMagicConstant(n - 1)
	total = add(total, solveI(n- 1))

	for i := 1; i < n - 1; i++ {
		coeff := mul(n - 1 - i, computeMagicConstant(i))
		total = add(total, 2 * solveI(i))
		total = add(total, coeff)
	}
	cachedMagicConstants[n] = total
	return  total

}

func solveI (n int) int {
	if v, ok := cachedValues[n]; ok {
		return v
	}
	total := 1
	for i:= 1; i < n; i++ {
		coeff := mul(add(mul((n - i), (n - i)), (n - i - 1)), solveI(i))
		total = add(total, coeff)
		for j := i + 1; j < n; j++ {
			coeff1 := add((n - j), (n - j - 1))
			coeff2 := add((n - j), (n - j - 1))
			asym := add(mul(j - i, computeMagicConstant(i)), solveI(i))
                        coeff3 := mul(add(coeff1, coeff2), asym)
			total = add(total, coeff3)
		}
	}

	cachedValues[n] = total
	return total
}






func solveAsymmetricCouple (couple Couple) int {
	if v, ok := cachedCouples[couple]; ok {
		return v
	}
	total := 1
	for i:= 1; i < couple[0]; i++ {
		coeff := mul(add(mul((couple[0] - i), (couple[1] - i)), (couple[0] - i - 1)), solveAsymmetricCouple(Couple{i, i}))
		total = add(total, coeff)
		for j := i + 1; j < couple[0]; j++ {
			coeff1 := add((couple[0] - j), (couple[1] - j - 1))
			coeff2 := add((couple[0] - j), (couple[0] - j - 1))
                        coeff3 := mul(add(coeff1, coeff2), solveAsymmetricCouple(Couple{i, j}))
			total = add(total, coeff3)
		}
	}

	cachedCouples[couple] = total
	return total
}




func handleError (err error){
	if err != nil {
		log.Fatal(err)
	}
}

func mul (a, b int) int {
	return (a * b) % MODULO
}

func add (a, b int) int {
	return (a + b) % MODULO
}
