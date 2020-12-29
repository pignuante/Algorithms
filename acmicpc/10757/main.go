package main

import (
	"bufio"
	"math/big"
	"os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()

	var A, B big.Int
	sc.Scan()
	A.SetString(sc.Text(), 10)
	sc.Scan()
	B.SetString(sc.Text(), 10)

	wt.WriteString(A.Add(&A, &B).String())
}
