package gla2

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
func DivRound(a int64, b int64) (int64, int64) {
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

	return div, ((a*10)/b - div*10)
}

/**	OpenFile opens a file using filepath/name.
 *	@param filepath - a string
 */
func OpenFile(filepath string) *os.File {
	a, b := DivRound(25, 2)
	fmt.Println(a, "+", b, "=", a+b)
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
 *
 *  @return - transaction as a string
 */
func generateTx() []byte {
	digits := make([]int64, 4)
	max := int64(9)
	min := int64(0)
	for i := range digits {
		min = 0
		if digits[1] == 0 && digits[0] == 0 && i == 3 {
			min = 1
		}
		digits[i] = rand.Int63()%(max-min+1) + min
	}
	return []byte(strconv.FormatInt(digits[0], 10) + strconv.FormatInt(digits[1], 10) + "." + strconv.FormatInt(digits[2], 10) + strconv.FormatInt(digits[3], 10) + "\n")
}

/**	GenerateRandomTxs generates a list of random numbers from 0.01 -> 99.99.
 *	and writes them in a new txs.txt file
 *	@param n - amount of random transactions
 */
func GenerateRandomTxs(n int) {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.Write((generateTx()))
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
	ret, _ := DivRound(sum, 100)
	if !hasParam {
		fmt.Println("sum(£): ", (ret))
	}
	return sum
}

func write2file(inn int64) []byte {
	innString := strconv.FormatInt(inn, 10)
	zFill := ""
	nFill := ""

	if len(innString) == 1 {
		nFill = innString
		zFill = "00.0"
	} else if len(innString) == 2 {
		nFill = innString
		zFill = "00."
	} else if len(innString) == 3 {
		zFill = "0"
		nFill = innString[:1] + "." + innString[1:]
	} else {
		nFill = innString[:2] + "." + innString[2:]
	}

	return []byte((zFill + nFill) + "\n")
}

func readString2int(inn string) int {
	returnVal, err := strconv.Atoi(strings.Join(strings.Split(inn, "."), ""))
	checkError(err)
	return returnVal
}

/**	createSubFile generates a list of values (val% of transactions).
 *	from transactions(main) file txs.txt writes them in a new (sub) file
 *	@param main - sample file
 *	@param sub - new file
 *	@param val - percentage of transactions from txs
 */
func createSubFile(main, sub *os.File, val int) { // created for ease of use, when editing a reoccuring function.
	defer main.Close()
	defer sub.Close()
	// read the file line by line using scanner
	getLines := bufio.NewScanner(main)
	var buf bytes.Buffer
	for getLines.Scan() {

		ret, _ := DivRound(int64((readString2int(getLines.Text()))), 100)
		buf.Write(write2file(ret * int64(val)))
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
		returnVal = append(returnVal, int64(readString2int(getLines.Text())))
	}
	return returnVal
}

/**	GenerateFees generates a list of fees (30% of transactions).
 *	from transactions file txs.txt writes them in a new fees.txt file
 *	@param n - amount of random transactions
 */
func GenerateFees() {
	createSubFile(OpenFile(txs), createFile(fees), 30)
}

/**	GenerateEarnings generates a list of earnings.
 *	(70% of transactions)from transactions file.
 *  And writes them in a new earnings.txt file
 */
func GenerateEarnings() {
	createSubFile(OpenFile(txs), createFile(earn), 70)
}

/**	Compare compares the data from the transaction files.
 *	Number1 = (Sum of fees.txt) minus (the fee of the total of transactions(txs.txt)).
 *	Number2 = (Sum of total.txt) minus (the fee of the total of transactions(txs.txt)).
 *  @return - both the numbers which should give 0,0
 */
func Compare() (int64, int64) {

	totalSum := Sum(OpenFile(txs))
	feesSum := Sum(OpenFile(fees))
	getFees, _ := DivRound(totalSum, 10)
	totalFees := getFees * 3
	totalEarnings := Sum(OpenFile(earn))

	fmt.Println(feesSum)
	fmt.Println(totalSum)
	fmt.Println(totalFees + totalEarnings)
	fmt.Println(totalEarnings)

	return (feesSum - totalFees), (totalSum - (totalEarnings + totalFees))

}
