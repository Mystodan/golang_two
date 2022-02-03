package gla2

import (
	"log"
	"math/rand"
	"os"
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
