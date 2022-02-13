# Golang Assignment 2 : simple currency processing
By Daniel Hao Huynh

## Backstory

I work for a company that developed a mobile App Store. My company sells apps, and takes a fee on each transaction going through the App Store.  

#### Flags
   * `-gen` generates a list of n random float32's from 0.01 to 0.99, generates the fees, and earnings thereafter saves them in their respective txt files
   * `-sum` get the sum of the default `transactions(txs.txt)`
   * `-getsum` input a string if you want to find the sum of a certain .txt file. f.ex. earnings.txt, leave empty for transactions
   * `-comp` compares the data from the transaction files, with fees and earnings
   * `-mill` Generates a million transactions, generates the fees, and earnings thereafter saves them in their respective txt files
   * `-perf` flag, indicates (print) performance (total time to do the given workflow).

#### Testing
   * Performed  pre-profiling unit testing using command `go test` using all of the functions from `functions.go` in the pckg folder on `main_test.go`
   * Performed post-profiling unit testing using command `go test` using all of the functions from `functions.go` in the pckg folder on `main_test.go`
   ###### (The time improved from around 14s to 2s after optimizing the code post profiling)

#### Profiling
   * Performed profiling on testing and main using theese commands:
   * - go test -cpuprofile cpu.prof -memprofile mem.prof -bench . 
   * - go run main.go -mill -sum -comp -cpuprofile cpu.prof  -memprofile mem.prof 
   
using the command `go tool pprof x.prof` (replace the x with either mem or cpu) to enter the editor concerning editing the export of profiling.<br>
Then the profiling of memory and cpu were exported as pdfs in their own folders.<br>
> saved in folders: `golang-two\profiling-tests --> \main or \unit-tests`  as  ..._CPU.pdf and ..._MEM.pdf
