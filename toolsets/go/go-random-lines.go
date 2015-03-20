package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	infile := "/users/nick/data/financial_insitition/all_2012/all_2012.csv"
	outfile := "../samples/banks.csv"
	// r := rand.New(time.Now().UnixNano())
	rand.Seed(241234)
	rnd := rand.Intn(10)
	fmt.Println(rnd)
	fmt.Println(infile)
	fmt.Println(outfile)
	f, err := os.Open(infile)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(f.Name)
	r := bufio.NewReader(f)
	s, e := bufio.Readln(r)
	for e == nil {
		fmt.Println(s)
		s, e = Readln(r)
	}
}
