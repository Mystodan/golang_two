package gla2

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

const (
	fPath = "./"
	txs   = fPath + "txs.txt"
	fees  = fPath + "fees.txt"
	earn  = fPath + "earnings.txt"
)

/**
 *    Divides two integers, and rounds using bankers rounds for more precision.
 *
 *		- a / b, approximate decimal-error.
 *
 *    @param a - Divident
 *    @param b - Divisor
 *
 */
func Round(a int64, b int64) (int64, int64) {
	div := a / b
	rest := a - div*b

	// Bankers round
	if rest*10 >= 5*b {
		if rest*10 == 5*b {
			if div%2 == 1 {
				div += 1
			}
		} else {
			div += 1
		}
	}

	return div, (a*10)/b - div*10
}

/**	OpenFile opens a file using filepath/name.
 *	@param filepath - a string
 */
func OpenFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	checkError(err)
	return file
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

/**	generate generates a random transaction from 0.01 -> 99.99.
 *	@param min - minimum int value
 *	@param max - maximum int value
 *  @return - transaction as a string
 */
func generateTx(min int64, max int64) string {
	var num1, num2 int64
	num1 = rand.Int63()%(max-min) + min
	if min == 0 {
		min = 1
	}
	num2 = rand.Int63()%(max-min) + min
	return fmt.Sprint(num1, ".", num2)

}

/**	GenerateRandomTxs generates a list of random numbers from 0.01 -> 99.99.
 *	and writes them in a new txs.txt file
 *	@param n - amount of random transactions
 */
func GenerateRandomTxs(n int) {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.Write([]byte((generateTx(0, 99)) + "\n"))
	}
	createFile(txs).Write(buf.Bytes())
}

/**	GenerateMillionTxs generates a list of random 1 million numbers.
 *	and writes them in a new txs.txt file
 */
func GenerateMillionTxs() {
	GenerateRandomTxs(1000000)
}

/**	Sum sums a list of transaction int numbers.
 *	and prints the sum
 *	@param	file - takes in a os.File.
 */
func Sum(file ...*os.File) int64 {
	sum := int64(0)
	hasParam := true
	if len(file) == 0 {
		hasParam = false
		file = append(file, OpenFile(txs))
	}
	defer file[0].Close()
	getData := readSubFiles(file[0])

	for _, val := range getData {
		sum += val
	}
	if !hasParam {
		fmt.Println("sum: ", (sum))
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
	var buf bytes.Buffer
	for getLines.Scan() {
		s, _ := strconv.ParseFloat(getLines.Text(), 64)
		val, _ := Round(int64((s)), int64(val*100))
		buf.Write([]byte(fmt.Sprint(val) + "\n"))

		/* improved by profiling
		 * instead of using the write function multiple times in a for loop, calling out a heavy load function
		 * multiple times, it now calls it only once after appending all values into a byte array.
		 * time reduced from 20.x seconds to 1.x seconds
		 */
	}
	_, _ = sub.Write(buf.Bytes())
	checkError(getLines.Err())
}

func readSubFiles(inn *os.File) []int64 {
	defer inn.Close()

	getLines := bufio.NewScanner(inn)
	returnVal := []int64{}
	for getLines.Scan() {
		s, _ := strconv.ParseInt(getLines.Text(), 10, 64)

		returnVal = append(returnVal, s)
	}
	return returnVal
}

/**	GenerateFees generates a list of fees (30% of transactions).
 *	from transactions file txs.txt writes them in a new fees.txt file
 *	@param n - amount of random transactions
 */
func GenerateFees() {
	createSubFile(OpenFile(txs), createFile(fees), 0.3)
}

/**	GenerateEarnings generates a list of earnings.
 *	(70% of transactions)from transactions file.
 *  And writes them in a new earnings.txt file
 */
func GenerateEarnings() {
	createSubFile(OpenFile(txs), createFile(earn), 0.7)
}

/**	Compare compares the data from the transaction files.
 *	Number1 = (Sum of fees.txt) minus (the fee of the total of transactions(txs.txt)).
 *	Number2 = (Sum of total.txt) minus (the fee of the total of transactions(txs.txt)).
 *  @return - both the numbers which should give 0,0
 */
func Compare() (int64, int64) {

	feesSum := Sum(OpenFile(fees))
	totalSum := Sum(OpenFile(txs))
	totalFees := int64((int64(totalSum)))
	totalEarnings := Sum(OpenFile(earn))
	/* 	fmt.Println(feesSum)
	   	fmt.Println(totalSum)
	   	fmt.Println(totalFees)
	   	fmt.Println(totalEarnings)
	   	fmt.Println("FEE DIFF: ", Round(Conv2dec(totalSum)*0.3-Conv2dec(feesSum), 2))
	   	fmt.Println("EARN DIFF: ", Round(Conv2dec(totalSum)*0.7-Conv2dec(totalEarnings), 2)) */

	return (feesSum - totalFees), (totalSum - (totalEarnings + totalFees))

}
