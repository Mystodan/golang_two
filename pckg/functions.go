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
	returnVal := []byte{}
	for i := 0; i < n; i++ {
		returnVal = append(returnVal, ([]byte(fmt.Sprint(R2Dec(generate(0.01, 99.99))) + "\n"))...)
	}
	_, _ = createFile("txs.txt").Write(returnVal)
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
	hasParam := true
	if len(file) > 1 {
		fmt.Println("Sum() uses only the first parameter, any other wil be left unused")
	}
	if len(file) == 0 {
		hasParam = false
		file = append(file, OpenFile("txs.txt"))
	}
	defer file[0].Close()

	// read the file line by line using scanner
	getLines := bufio.NewScanner(file[0])

	for getLines.Scan() {
		if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
			sum = sum + s
		}
	}
	checkError(getLines.Err())
	if !hasParam {
		fmt.Println("sum: ", R2Dec(sum))
	}
	return sum
}

/**	createSubFile generates a list of values (val% of transactions).
 *	from transactions(main) file txs.txt writes them in a new (sub) file
 *	@param main - sample file
 *	@param sub - new file
 *	@param val - percentage of transactions from txs
 */
func createSubFile(main, sub *os.File, val float64) { // created for ease of use, when editing a reoccuring function.
	defer main.Close()
	defer sub.Close()
	// read the file line by line using scanner
	getLines := bufio.NewScanner(main)

	returnVal := []byte{}
	for getLines.Scan() {
		s, _ := strconv.ParseFloat(getLines.Text(), 64)
		returnVal = append(returnVal, []byte(fmt.Sprint(R2Dec(s*val))+"\n")...)
		/* improved by profiling
		 * instead of using the write function multiple times in a for loop, calling out a heavy load function
		 * multiple times, it now calls it only once after appending all values into a byte array.
		 * time reduced from 20.x seconds to 1.x seconds
		 */
	}
	_, _ = sub.Write(returnVal)
	checkError(getLines.Err())
}

/**	GenerateFees generates a list of fees (30% of transactions).
 *	from transactions file txs.txt writes them in a new fees.txt file
 *	@param n - amount of random transactions
 */
func GenerateFees() {
	createSubFile(OpenFile("txs.txt"), createFile("fees.txt"), 0.3)
}

/**	GenerateEarnings generates a list of earnings.
 *	(70% of transactions)from transactions file.
 *  And writes them in a new earnings.txt file
 */
func GenerateEarnings() {
	createSubFile(OpenFile("txs.txt"), createFile("earnings.txt"), 0.7)
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
