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

/**	openFile opens a file using filepath/name.
 *	@param filepath - a string
 */
func openFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	checkErr(err)
	return file
}

/**	createFile creates a file using filepath/name.
 *	@param filepath - a string
 */
func createFile(filepath string) *os.File {
	file, err := os.Create(filepath)
	checkErr(err)
	return file
}

/**	checkErr logs an error.
 *	@param inn - error value
 */
func checkErr(inn error) {
	if inn != nil {
		log.Fatal(inn)
	}
}

/**	generate generates a random float64 number
 *	@param min - minimum float64 value
 *	@param max - maximum float64 value
 */
func generate(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

/**	GenerateRandomTxs generates a list of random float64 numbers.
 *	and writes them in a new txs.txt file
 *	@param n - amount of random transactions
 */
func GenerateRandomTxs(n int) {

	file, err := os.Create("txs.txt")
	checkErr(err)
	defer file.Close()

	for i := 0; i < n; i++ {
		file.Write([]byte(fmt.Sprint(math.Round(generate(0.0, 99.99)*100)/100) + "\n"))
	}
}

/**	GenerateMillionTxs generates a list of random 1 million float64 numbers.
 *	and writes them in a new txs.txt file
 */
func GenerateMillionTxs() {
	GenerateRandomTxs(1000000)
}

/**	Sum sums a list of transaction float64 numbers.
 *	and prints the sum
 *	@param	file - takes in a os.File.
 */
func Sum(file ...*os.File) float64 {
	var sum float64
	noPar := false
	if len(file) > 1 {
		fmt.Println("Sum() uses only the first parameter, any other wil be left unused")
	}
	if len(file) < 1 {
		file = append(file, openFile("txs.txt"))
		noPar = true
	}
	defer file[0].Close()
	// read the file line by line using scanner
	getLines := bufio.NewScanner(file[0])

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			sum += s
		}
	}
	checkErr(getLines.Err())
	if noPar {
		fmt.Println("sum: ", math.Round(sum*100)/100)
	}
	return sum
}

/**	GenerateFees generates a list of fees (30% of transactions).
 *	from transactions file txs.txt writes them in a new fees.txt file
 *	@param n - amount of random transactions
 */
func GenerateFees() {
	openFile := openFile("txs.txt")
	file := createFile("fees.txt")

	defer openFile.Close()
	defer file.Close()
	// read the file line by line using scanner
	getLines := bufio.NewScanner(openFile)

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			normalfee := math.Round((s*0.3)*100) / 100
			file.Write([]byte(fmt.Sprint(normalfee) + "\n"))
		}
	}
	checkErr(getLines.Err())
}

/**	GenerateEarnings generates a list of earnings.
 *	(70% of transactions)from transactions file.
 *  And writes them in a new earnings.txt file
 */
func GenerateEarnings() {
	openFile := openFile("txs.txt")
	file := createFile("earnings.txt")

	defer openFile.Close()
	defer file.Close()

	// read the file line by line using scanner
	getLines := bufio.NewScanner(openFile)

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			normalfee := math.Round((s*0.7)*100) / 100
			file.Write([]byte(fmt.Sprint(normalfee) + "\n"))
		}
	}
	checkErr(getLines.Err())
}

/**	Compare compares the data from the transaction files.
 *	Number1 = (Sum of fees.txt) minus (the fee of the total of transactions(txs.txt)).
 *	Number2 = (Sum of total.txt) minus (the fee of the total of transactions(txs.txt)).
 *  @return - both the numbers which should give 0,0
 */
func Compare() (float64, float64) {

	feesSum := Sum(openFile("fees.txt"))
	totalSum := Sum(openFile("txs.txt"))
	totalFees := totalSum * 0.3
	totalEarnings := Sum(openFile("earnings.txt"))

	return math.Round((feesSum-totalFees)*100) / 100, math.Round((totalSum-(totalEarnings+totalFees))*100) / 100

}
