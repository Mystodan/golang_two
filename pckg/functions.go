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

/**	OpenFile opens a file using filepath/name.
 *	@param filepath - a string
 */
func OpenFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	checkError(err)
	return file
}

/**	R2Dec rounds a float to 2 decimals
 *	@param n float64
 *  @return a float64 with 2 decimals
 */
func R2Dec(n float64) float64 {
	return math.Round(n*100) / 100
}

/**	createFile creates a file using filepath/name.
 *	@param filepath - a string
 */
func createFile(filepath string) *os.File {
	file, err := os.Create(filepath)
	checkError(err)
	return file
}

/**	checkError logs an error.
 *	@param inn - error value
 */
func checkError(inn error) {
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

/**	GenerateRandomTxs generates a list of random float64 numbers from 0.01 -> 99.99.
 *	and writes them in a new txs.txt file
 *	@param n - amount of random transactions
 */
func GenerateRandomTxs(n int) {

	file, err := os.Create("txs.txt")
	checkError(err)
	defer file.Close()

	for i := 0; i < n; i++ {
		file.Write([]byte(fmt.Sprint(R2Dec(generate(0.01, 99.99))) + "\n"))
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
	defer file[0].Close()

	if len(file) > 1 {
		fmt.Println("Sum() uses only the first parameter, any other wil be left unused")
	}
	if len(file) < 1 {
		file = append(file, OpenFile("txs.txt"))
	}

	// read the file line by line using scanner
	getLines := bufio.NewScanner(file[0])

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			sum += s
		}
	}
	checkError(getLines.Err())
	if len(file) < 1 {
		fmt.Println("sum: ", R2Dec(sum))
	}
	return sum
}

/**	GenerateFees generates a list of fees (30% of transactions).
 *	from transactions file txs.txt writes them in a new fees.txt file
 *	@param n - amount of random transactions
 */
func GenerateFees() {
	getFile := OpenFile("txs.txt")
	file := createFile("fees.txt")

	defer getFile.Close()
	defer file.Close()
	// read the file line by line using scanner
	getLines := bufio.NewScanner(getFile)

	for getLines.Scan() {
		s, _ := strconv.ParseFloat(getLines.Text(), 64)
		file.Write([]byte(fmt.Sprint(R2Dec(s*0.3)) + "\n"))
	}
	checkError(getLines.Err())
}

/**	GenerateEarnings generates a list of earnings.
 *	(70% of transactions)from transactions file.
 *  And writes them in a new earnings.txt file
 */
func GenerateEarnings() {
	OpenFile := OpenFile("txs.txt")
	file := createFile("earnings.txt")

	defer OpenFile.Close()
	defer file.Close()

	// read the file line by line using scanner
	getLines := bufio.NewScanner(OpenFile)

	for getLines.Scan() {
		s, _ := strconv.ParseFloat(getLines.Text(), 64)
		file.Write([]byte(fmt.Sprint(R2Dec(s*0.7)) + "\n"))

	}
	checkError(getLines.Err())
}

/**	Compare compares the data from the transaction files.
 *	Number1 = (Sum of fees.txt) minus (the fee of the total of transactions(txs.txt)).
 *	Number2 = (Sum of total.txt) minus (the fee of the total of transactions(txs.txt)).
 *  @return - both the numbers which should give 0,0
 */
func Compare() (float64, float64) {

	feesSum := Sum(OpenFile("fees.txt"))
	totalSum := Sum(OpenFile("txs.txt"))
	totalFees := totalSum * 0.3
	totalEarnings := Sum(OpenFile("earnings.txt"))

	return R2Dec(feesSum - totalFees), R2Dec(totalSum - (totalEarnings + totalFees))

}
