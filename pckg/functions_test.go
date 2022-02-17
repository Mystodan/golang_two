package gla2_test

import (
	"bufio"
	"fmt"
	gla2 "golang-two/pckg"
	"math/rand"
	"strconv"
	"testing"
)

/** compareFloat(a, b []float64) - compares two float arrays.
 *  @param a - float array to compare
 *  @param b - float array to compare
 *  @return bool - returns true if they are equal
 */
func compareFloat(a, b []float64) bool {
	if !(len(a) == len(b)) {
		return false
	} else {
		for i, v := range b {
			if !(v == b[i]) {
				return false
			}
		}
	}
	return true
}

/** TestGenerate(t *testing.T)
 *	tests the following functions:
 *	 @see GenerateRandomTxs()
 *	 @see GenerateFees()
 *	 @see GenerateEarnings()
 */
func TestGenerate(t *testing.T) {
	var rTests = []struct {
		seed int64
		size int
		want []float64
	}{ // Ideal values for tests
		{1, 3, []float64{ // test with seed(1) and txs amount(3)
			60.46, 94.04, 66.45}},
		{2, 10, []float64{ // test with seed(2) and txs amount(10)
			16.74, 26.51, 5.16, 11.91, 61.42, 92.76, 42.5, 20.62, 21.2, 12.7}},
		{3, 15, []float64{ // test with seed(3) and txs amount(15)
			71.99, 65.26, 94.19, 76.81, 89.35, 21.91, 42.76, 50.77, 32.51, 46.84, 35.03, 79.56, 81.7, 34.94, 96.03}},
	}

	for i, tString := range rTests {
		rand.Seed(tString.seed) // sets the seed
		feesWant := []float64{} //values to read file
		earnWant := []float64{}

		solutionTxs := []float64{}
		solutionEarn := []float64{}
		solutionFees := []float64{}

		gla2.GenerateRandomTxs(tString.size) //generates amount(size) of transactions
		gla2.GenerateFees()                  //generates fees based on generated transactions(30%)
		gla2.GenerateEarnings()              //generates earnings based on generated transactions(70%)

		txsFile := gla2.OpenFile("txs.txt") // reads generated files
		earnFile := gla2.OpenFile("earnings.txt")
		feesFile := gla2.OpenFile("fees.txt")
		defer txsFile.Close() // closes files at the very end of function
		defer earnFile.Close()
		defer feesFile.Close()

		///
		///	GETTERS FOR VALUES FROM FILE
		///

		getLines := bufio.NewScanner(txsFile) // creates a new scanner on transactions file

		for getLines.Scan() {
			s, _ := strconv.ParseFloat(getLines.Text(), 64)
			solutionTxs = append(solutionTxs, s)           // TRANSACTIONS
			solutionEarn = append(solutionEarn, (s * 0.7)) // EARNING RATE
			solutionFees = append(solutionFees, (s * 0.3)) // FEES RATE

		}

		getLines = bufio.NewScanner(feesFile) // sets new scanner on fees file

		for getLines.Scan() {
			s, _ := strconv.ParseFloat(getLines.Text(), 64)
			feesWant = append(feesWant, s)

		}
		getLines = bufio.NewScanner(earnFile) // sets new scanner on earn file

		for getLines.Scan() {
			s, _ := strconv.ParseFloat(getLines.Text(), 64)
			earnWant = append(earnWant, s)

		}

		///
		/// COMPARES VALUES FROM FILES WITH IDEAL VALUES
		///

		fmt.Println("Testing GenerateRandomTxs()x", i+1, " ... ")
		if !compareFloat(solutionTxs, tString.want) {
			t.Fatal("(", i+1, ")Failed at comparing Transactions!")
		} else {
			fmt.Println("(", i+1, ")Passed!")
		}
		fmt.Println("Testing GenerateFees()x", i+1, " ... ")
		if !compareFloat(solutionFees, feesWant) {
			t.Fatal("(", i+1, ")Failed at comparing Fees!")
		} else {
			fmt.Println("(", i+1, ")Passed!")
		}
		fmt.Println("Testing GenerateEarnings()x", i+1, "  ... ")
		if !compareFloat(solutionEarn, earnWant) {
			t.Fatal("(", i+1, ")Failed at comparing Earnings!")
		} else {
			fmt.Println("(", i+1, ")Passed!")
		}
	}
}

/** TestSum(t *testing.T)
 *	tests the following functions:
 *	 @see GenerateRandomTxs()
 *	 @see Sum()
 */
func TestSum(t *testing.T) {
	var rTests = []struct {
		seed int64
		size int
		want int64
	}{ // Ideal values for tests
		{1, 3, 14146},   // test with seed(1) and txs amount(3) and the sum
		{99, 10, 36872}, // test with seed(1) and txs amount(3) and the sum
		{13, 8, 29853},  // test with seed(1) and txs amount(3) and the sum
	}

	for i, tString := range rTests {
		rand.Seed(tString.seed)                        // sets the seed
		gla2.GenerateRandomTxs(tString.size)           //generates amount(size) of transactions
		solution := gla2.Sum(gla2.OpenFile("txs.txt")) // Saves the sum from function

		/// COMPARES VALUES FROM FILES WITH IDEAL VALUES
		fmt.Println("Testing Sum()x", i+1, " ... ")
		if solution != tString.want {
			t.Fatal("(", i+1, ")Failed at comparing Sum!")
		} else {
			fmt.Println("(", i+1, ")Passed!")
		}

	}
}

/** TestMillion()
 *	tests the following functions:
 *	 @see GenerateMillionTxs()
 */
func TestMillion(t *testing.T) {
	var rTests = []struct {
		seed int64
		want int64
	}{ // Ideal values for tests
		{1, 1000000}, // test with seed(1) and txs amount(3) and the sum
	}
	for i, tString := range rTests {
		rand.Seed(tString.seed)   // sets the seed
		gla2.GenerateMillionTxs() // generates transactions

		var solution int64
		getLines := bufio.NewScanner(gla2.OpenFile("txs.txt")) // sets new scanner on fees file

		for getLines.Scan() { //counts the amouont of values
			if _, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
				solution++ //saves amount in solution
			}
		}
		/// COMPARES VALUES FROM FILES WITH IDEAL VALUES
		fmt.Println("Testing GenerateMillionTxs()x", i+1, " ... ")
		if solution != tString.want {
			t.Fatal("(", i+1, ")Failed at counting amount of transactions!")
		} else {
			fmt.Println("(", i+1, ")Passed!")
		}
	}
}

/** TestCompare(t *testing.T)
 *	tests the following functions:
 *	 @see the GenerateRandomTxs()
 *	 @see the GenerateFees()
 *	 @see the GenerateEarnings()
 *	 @see GenerateMillionTxs()
 *	 @see Compare()
 */
func TestCompare(t *testing.T) {
	var rTests = []struct {
		seed   int64
		want   []int64
		amount int
	}{ // Ideal values for tests
		// test with seed(1) and wanted return value amount and the amount of transactions{ 0 = a million transactions}
		{5, []int64{0, 0}, 10},    // 10 transactions, seed = 5
		{10, []int64{0, 0}, 30},   // 30 transactions, seed = 10
		{1, []int64{282, -74}, 0}, // 0 in amount = 1 million, seed = 1
	}
	for i, tString := range rTests {
		rand.Seed(tString.seed) // sets the seed
		if tString.amount < 1 { // generates a million transaction if amount is less than 1
			gla2.GenerateMillionTxs()
		} else {
			gla2.GenerateRandomTxs(tString.amount)
		}
		gla2.GenerateFees()     // necessary for compare funtion
		gla2.GenerateEarnings() // necessary for compare funtion

		Number1, Number2 := gla2.Compare() // runs compare funtion
		solution := []int64{int64(Number1), int64(Number2)}
		/// COMPARES VALUES FROM FILES WITH IDEAL VALUES
		fmt.Println("Testing Compare() x", i+1, ")  ... ")
		if (solution[0] != tString.want[0]) && (solution[1] != tString.want[1]) {
			t.Fatal("(", i+1, ")Failed at comparing transactions with other values!")
		} else {
			fmt.Println("(", i+1, ")Passed!")
		}
	}
}
