package main_test

import (
	"bufio"
	"fmt"
	gla2 "golangAss2/pckg"
	"math/rand"
	"strconv"
	"testing"
)

func compareFloat(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

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
		testN := fmt.Sprintf("%f", tString.want)
		t.Run(testN, func(t *testing.T) {
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
				if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
					solutionTxs = append(solutionTxs, s)                   // TRANSACTIONS
					solutionEarn = append(solutionEarn, gla2.R2Dec(s*0.7)) // EARNING RATE
					solutionFees = append(solutionFees, gla2.R2Dec(s*0.3)) // FEES RATE
				}
			}

			getLines = bufio.NewScanner(feesFile) // sets new scanner on fees file

			for getLines.Scan() {
				if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
					feesWant = append(feesWant, s)
				}
			}
			getLines = bufio.NewScanner(earnFile) // sets new scanner on earn file

			for getLines.Scan() {
				if s, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
					earnWant = append(earnWant, s)
				}
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
		})
	}
}

func TestSum(t *testing.T) {
	var rTests = []struct {
		seed int64
		size int
		want float64
	}{ // Ideal values for tests
		{1, 3, 220.95},   // test with seed(1) and txs amount(3) and the sum
		{99, 10, 457.93}, // test with seed(1) and txs amount(3) and the sum
		{13, 8, 362.21},  // test with seed(1) and txs amount(3) and the sum
	}

	for i, tString := range rTests {
		rand.Seed(tString.seed) // sets the seed
		testN := fmt.Sprintf("%f", tString.want)
		t.Run(testN, func(t *testing.T) {
			gla2.GenerateRandomTxs(tString.size) //generates amount(size) of transactions

			solution := gla2.R2Dec(gla2.Sum(gla2.OpenFile("txs.txt")))

			/// COMPARES VALUES FROM FILES WITH IDEAL VALUES
			fmt.Println("Testing Sum()x", i+1, " ... ")
			if solution != tString.want {
				t.Fatal("(", i+1, ")Failed at comparing Sum!")
			} else {
				fmt.Println("(", i+1, ")Passed!")
			}
		})
	}
}

func TestMillion(t *testing.T) {
	var rTests = []struct {
		seed int64
		want int64
	}{ // Ideal values for tests
		{1, 1000000}, // test with seed(1) and txs amount(3) and the sum
	}
	for i, tString := range rTests {
		rand.Seed(tString.seed) // sets the seed
		testN := fmt.Sprintf("%d", tString.want)
		t.Run(testN, func(t *testing.T) {
			gla2.GenerateMillionTxs()

			var solution int64
			getLines := bufio.NewScanner(gla2.OpenFile("txs.txt")) // sets new scanner on fees file

			for getLines.Scan() {
				if _, err := strconv.ParseFloat(getLines.Text(), 64); err == nil {
					solution++
				}
			}
			/// COMPARES VALUES FROM FILES WITH IDEAL VALUES
			fmt.Println("Testing GenerateMillionTxs()x", i+1, " ... ")
			if solution != tString.want {
				t.Fatal("(", i+1, ")Failed at counting amount of transactions!")
			} else {
				fmt.Println("(", i+1, ")Passed!")
			}
		})
	}
}

func TestCompare(t *testing.T) {
	var rTests = []struct {
		seed   int64
		want   []float64
		amount int
	}{ // Ideal values for tests
		// test with seed(1) and wanted return value amount and the amount of transactions{ 0 = a million transactions}
		{1, []float64{255.68, -103.63}, 10}, // 10 transactions
		{1, []float64{-0.01, -0.01}, 0},     // 1 million
	}
	for i, tString := range rTests {
		rand.Seed(tString.seed) // sets the seed
		testN := fmt.Sprintf("%f", tString.want)
		t.Run(testN, func(t *testing.T) {
			if tString.amount == 0 {
				gla2.GenerateMillionTxs()
			} else {
				gla2.GenerateRandomTxs(tString.amount)
			}
			gla2.GenerateFees()
			gla2.GenerateEarnings()

			Number1, Number2 := gla2.Compare()

			solution := []float64{Number1, Number2}

			/// COMPARES VALUES FROM FILES WITH IDEAL VALUES

			fmt.Println(Number1, Number2)
			fmt.Println("Testing Compare() x", i+1, ")  ... ")
			if solution[0] != tString.want[0] {
				t.Fatal("(", i+1, ")Failed at comparing transactions with other values!")
			} else {
				fmt.Println("(", i+1, ")Passed!")
			}
		})
	}
}
