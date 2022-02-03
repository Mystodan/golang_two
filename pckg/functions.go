package gla2

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func checkErr(inn error) {
	if inn != nil {
		log.Fatal(inn)
	}
}
func generate(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func GenerateRandomTxs(n int) {
	file, err := os.Create("tsx.txt")
	checkErr(err)
	defer file.Close()
	var rString string

	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			randInt := generate(48, 57)
			if !(randInt == 48 && j == 0) {
				rString += string(rune(randInt))
			}
			if j == 1 {
				rString += "."
			}
		}
		rString += "\n"
	}
	file.WriteString(rString)
}
func Sum() {
	var sum float64
	file, err := os.Open("tsx.txt")
	checkErr(err)
	defer file.Close()

	// read the file line by line using scanner
	getLines := bufio.NewScanner(file)

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			sum += s
		}
	}
	checkErr(getLines.Err())
	fmt.Println(sum)

}
