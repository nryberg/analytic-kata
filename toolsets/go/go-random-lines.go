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

func fetchLine(f *os.File, line int) string {
	count := 0
	s := bufio.NewScanner(f)
	var output string
	for s.Scan() {

		if count == line {
			output = s.Text()
		}
		count += 1
	}
	return output
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	infile := "/users/nick/data/financial_insitition/all_2012/all_2012.csv"
	outfile := "../samples/banks.csv"
	sample_size := 10
	//rnd := rand.Intn(10)
	//fmt.Println(rnd)
	fmt.Println(infile)
	fmt.Println(outfile)
	f, err := os.Open(infile)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	f_out, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("error opening out file: %v\n", err)
		os.Exit(1)
	}

	defer f_out.Close()

	count, err := lineCounter(f)
	f.Seek(0, 0)

	// TODO: Create array of line numbers, sort the array and then
	// simply pull those.  No need to seek.

	for counter := 0; counter < sample_size; counter++ {
		sample_line := int(rand.Int63n(int64(count)))
		fmt.Println(counter, sample_line)
		outline := fetchLine(f, sample_line)
		f.Seek(0, 0)
		_, err := f_out.WriteString(outline + "\n")
		if err != nil {
			fmt.Printf("error writing to file: %v\n", err)
			os.Exit(1)
		}
	}

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

func lineCounter(f *os.File) (int, error) {
	r := bufio.NewReader(f)
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
