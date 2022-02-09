package main

import (
	"flag"
	"fmt"
	gla2 "golangAss2/pckg"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var flagGen int
	flag.IntVar(&flagGen, "gen", 0, "generates a list of n random float32's from 0.01 to 0.99 ")
	var flagSum string
	flag.StringVar(&flagSum, "sum", "", "input a string if you want to find the sum of a certain .txt file. f.ex. earnings.txt, leave empty for transactions")
	var flagComp bool
	flag.BoolVar(&flagComp, "comp", false, "compares the data from the transaction files, with fees and earnings")
	var flagMiln bool
	flag.BoolVar(&flagMiln, "mill", false, "Generates a million transactions")

	flag.Parse()

	if flagMiln {
		gla2.GenerateMillionTxs()
		gla2.GenerateFees()
		gla2.GenerateEarnings()
	}
	if flagGen > 0 && !flagMiln {
		gla2.GenerateRandomTxs(flagGen)
		gla2.GenerateFees()
		gla2.GenerateEarnings()
	}
	if len(flagSum) > 4 {
		gla2.Sum(gla2.OpenFile(flagSum))
		fmt.Println("sum of", flagSum, ":", gla2.R2Dec(gla2.Sum(gla2.OpenFile(flagSum))))
	} else {
		gla2.Sum()
	}
	if flagComp {
		Number1, Number2 := gla2.Compare()
		fmt.Println("comparing Number1 to Number2: ", Number1, ": ", Number2)
	}

}
