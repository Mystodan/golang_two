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

	for _, tString := range rTests {
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

			fmt.Println("Testing GenerateRandomTxs() ... ")
			if !compareFloat(solutionTxs, tString.want) {
				t.Fatal("Failed at comparing Transactions!")
			} else {
				fmt.Println("Passed! ")
			}
			fmt.Println("Testing GenerateFees() ... ")
			if !compareFloat(solutionFees, feesWant) {
				t.Fatal("Failed at comparing Fees!")
			} else {
				fmt.Println("Passed! ")
			}
			fmt.Println("Testing GenerateEarnings() ... ")
			if !compareFloat(solutionEarn, earnWant) {
				t.Fatal("Failed at comparing Earnings!")
			} else {
				fmt.Println("Passed! ")
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

	for _, tString := range rTests {
		rand.Seed(tString.seed) // sets the seed
		testN := fmt.Sprintf("%f", tString.want)
		t.Run(testN, func(t *testing.T) {
			gla2.GenerateRandomTxs(tString.size) //generates amount(size) of transactions

			solution := gla2.R2Dec(gla2.Sum(gla2.OpenFile("txs.txt")))

			fmt.Println("HERES THE SOL", solution)
			/// COMPARES VALUES FROM FILES WITH IDEAL VALUES
			fmt.Println("Testing GenerateRandomTxs() ... ")
			if solution != tString.want {
				t.Fatal("Failed at comparing Sum!")
			} else {
				fmt.Println("Passed! ")
			}
		})
	}
}
