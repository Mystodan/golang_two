package main

import (
	"flag"
	"fmt"
	gla2 "golang-two/pckg"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// Setting up for profiling for memory and cpu
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	var flagGen int
	flag.IntVar(&flagGen, "gen", 0, "generates a list of n random float32's from 0.01 to 0.99 ")
	var flagSum bool
	flag.BoolVar(&flagSum, "sum", false, "get the sum of transactions(txs.txt)")
	var flagGetSum string
	flag.StringVar(&flagGetSum, "getsum", "", "input a string if you want to find the sum of a certain .txt file. f.ex. earnings.txt, leave empty for transactions")
	var flagComp bool
	flag.BoolVar(&flagComp, "comp", false, "compares the data from the transaction files, with fees and earnings")
	var flagMiln bool
	flag.BoolVar(&flagMiln, "mill", false, "Generates a million transactions")
	var flagPerf bool
	flag.BoolVar(&flagPerf, "perf", false, "indicates (print) performance (total time to do the given workflow).")

	flag.Parse()

	if *cpuprofile != "" { // flag for cpu profiling
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	if flagMiln { // flag for generate million

		gla2.GenerateMillionTxs()
		gla2.GenerateFees()
		gla2.GenerateEarnings()
	}
	if flagGen > 0 && !flagMiln { // flag for generate

		gla2.GenerateRandomTxs(flagGen)
		gla2.GenerateFees()
		gla2.GenerateEarnings()

	}
	if flagComp { // flag for compare

		Number1, Number2 := gla2.Compare()
		fmt.Println("comparing Number1 to Number2: ", Number1, ": ", Number2)
	}
	if flagSum { // flag for sum
		gla2.Sum()

	}
	if len(flagGetSum) > 0 { // flag for sumgetter
		gla2.Sum(gla2.OpenFile(flagGetSum))
		fmt.Println("sum of", flagGetSum, ":", (int64(gla2.Sum(gla2.OpenFile(flagGetSum)))))
	}
	if !(flagMiln || flagComp || flagGen > 0 || flagSum || len(flagGetSum) > 0) { // if no flags
		gla2.GenerateMillionTxs()
		gla2.GenerateFees()
		gla2.GenerateEarnings()
		num1, num2 := gla2.Compare()
		fmt.Println("Result:", num1, ", ", num2)
	}

	if *memprofile != "" { // flag for profiling
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	if flagPerf { // flag for time after flag workflow
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println("Elapsed time after given workflow: ", elapsed)
	}
}
