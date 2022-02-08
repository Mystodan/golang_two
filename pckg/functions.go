package gla2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

/**	openFile using filepath/name.
 *	@param filepath - a string
 */
func openFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	checkErr(err)
	return file
}
func createFile(filepath string) *os.File {
	file, err := os.Create(filepath)
	checkErr(err)
	return file
}

func checkErr(inn error) {
	if inn != nil {
		log.Fatal(inn)
	}
}

func generate(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GenerateRandomTxs(n int) {

	file, err := os.Create("tsx.txt")
	checkErr(err)
	defer file.Close()

	for i := 0; i < n; i++ {
		file.Write([]byte(fmt.Sprint(math.Round(generate(0.0, 99.99)*100)/100) + "\n"))
	}
}
func GenerateMillionTxs() {
	GenerateRandomTxs(1000000)
}

func Sum() {
	fmt.Println("sum: ", math.Round(getSum(openFile("tsx.txt"))*100)/100)
}
func getSum(file *os.File) float64 {
	var sum float64
	defer file.Close()

	// read the file line by line using scanner
	getLines := bufio.NewScanner(file)

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			sum += s
		}
	}
	checkErr(getLines.Err())
	return sum
}

func GenerateFees() {
	openFile := openFile("tsx.txt")
	file := createFile("normal-fees.txt")

	defer openFile.Close()
	defer file.Close()
	// read the file line by line using scanner
	getLines := bufio.NewScanner(openFile)

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			normalfee := s * 0.3
			file.Write([]byte(fmt.Sprint(normalfee) + "\n"))
		}
	}
	checkErr(getLines.Err())
}

func GenerateEarnings() {
	openFile := openFile("tsx.txt")
	file := createFile("earnings.txt")

	defer openFile.Close()
	defer file.Close()
	// read the file line by line using scanner
	getLines := bufio.NewScanner(openFile)

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			normalfee := s * 0.7
			file.Write([]byte(fmt.Sprint(normalfee) + "\n"))
		}
	}
	checkErr(getLines.Err())
}
func Compare() (float64, float64) {

	feesSum := getSum(openFile("normal-fees.txt"))
	totalSum := getSum(openFile("tsx.txt"))
	totalFees := totalSum * 0.3
	totalEarnings := getSum(openFile("earnings.txt"))

	return math.Round((feesSum-totalFees)*100) / 100, math.Round((totalSum-(totalEarnings+totalFees))*100) / 100

}
