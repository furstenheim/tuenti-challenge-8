package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func main () {
	conn, err := net.Dial("tcp", "52.49.91.111:3241")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(conn)
	status, err := reader.ReadString('\n')
	handleError(err)
	nCases, err := strconv.Atoi(strings.Fields(status)[1])
	handleError(err)
	log.Println(status)
	status, err = reader.ReadString('\n')
	handleError(err)
	log.Println(status)
	fmt.Fprint(conn, "SUBMIT\n")
	status, err = reader.ReadString('\n')
	handleError(err)
	log.Println(status)

	for i:= 0; i < nCases; i++ {
		// Problem
		status, err = reader.ReadString('\n')
		handleError(err)
		log.Println(status)

		c := Case{parts: strings.Fields(status)}
		success := solveCase(c)

		fmt.Fprint(conn, success + "\n")

		status, err = reader.ReadString('\n')
		handleError(err)
		log.Println(status)
	}


}

func handleError (err error){
	if err != nil {
		log.Fatal(err)
	}
}
