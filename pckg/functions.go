package gla2

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Setting up path for transaction files
const (
	fPath = "./"
	txs   = fPath + "txs.txt"
	fees  = fPath + "fees.txt"
	earn  = fPath + "earnings.txt"
)

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
	getData := readFiles(file[0])
	for _, val := range getData {
		sum += val
	}
	if !hasParam {
		fmt.Println("sum(¢): ", (sum))
	}
	return sum
}

/**	write2file formats an int into a floatvalue as string and returns it.
 *	by essentially getting the length of the int and converting it into string.
 *	@param inn - int64
 *	@return []byte for working with buffers
 */
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

/**	readString2int reads a float(as string) and returns it as an int.
 *	by essentially splitting a string by its ".", and joining it.
 *	@param inn - read string
 *	@return int - string without "."
 */
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

		ret := math.RoundToEven((float64(readString2int(getLines.Text())) * float64(val) / 100))
		buf.Write(write2file(int64(ret)))
		/* improved by profiling
		 * instead of using the write function multiple times in a for loop, calling out a heavy load function
		 * multiple times, it now calls it only once after appending all values into a byte array.
		 * time reduced from 20.x seconds to 1.x seconds
		 */
	}
	_, _ = sub.Write(buf.Bytes())
	checkError(getLines.Err())
}

/**	readFiles reads a list of values of either transactions or a subfile.
 *	from transactions(main) file txs.txt returns them in an int array.
 *	@param inn - sample file
 *	@return returnVal - []int64 with values from file
 */
func readFiles(inn *os.File) []int64 {
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
	totalFees := int64(math.RoundToEven(float64(totalSum) * 0.3))
	totalEarnings := Sum(OpenFile(earn))

	return (feesSum - totalFees), (totalSum - (totalEarnings + totalFees))
}
