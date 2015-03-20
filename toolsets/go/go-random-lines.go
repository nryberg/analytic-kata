package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//func simple_sample(rdr *io.Reader) {
//}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	infile := "/users/nick/data/financial_insitition/all_2012/all_2012.csv"
	outfile := "../samples/banks.csv"
	//rnd := rand.Intn(10)
	//fmt.Println(rnd)
	fmt.Println(infile)
	fmt.Println(outfile)
	f, err := os.Open(infile)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	r := bufio.NewReader(f)

	count, err := lineCounter(r)
	fmt.Println("Lines : ", NumberToString(count, ','))
	fmt.Println("Rand: ", NumberToString(int(rand.Int63n(int64(count))), ','))
	/*
		s, e := r.ReadString('\n')
		for i := 1; i < 5; i++ {
			fmt.Println(s)
			s, e = r.ReadString('\n')
			if e != nil {
				fmt.Printf("error : %v\n", e)
				os.Exit(1)
			}
		}
	*/
}

func NumberToString(n int, sep rune) string {

	s := strconv.Itoa(n)

	startOffset := 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	l := len(s)

	commaIndex := 3 - ((l - startOffset) % 3)

	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < l; i++ {

		if commaIndex == 3 {
			buff.WriteRune(sep)
			commaIndex = 0
		}
		commaIndex++

		buff.WriteByte(s[i])
	}

	return buff.String()
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 8196)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return count, err
		}

		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func fetchLine(r io.Reader, line int) (string, error) {
	buf := make([]byte, 8196)

	lineSep := []byte{'\n'}

	for count := 0; count < line; count++ {
		str, err := r.ReadString(r, '\n')
		if err != nil && err != io.EOF {
			return count, err
		}

		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			break
		}
	}

	return count, nil
}
