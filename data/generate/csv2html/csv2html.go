package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file_in := "base_fruit.csv"
	path := path.Base(file_in)

	file_chunks := strings.Split(path, ".")
	file_out := file_chunks[0] + ".html"
	fmt.Println(file_chunks)

	fmt.Println(path)
	fmt.Println(file_out)
	f, err := os.Open(file_in)
	check(err)
	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
