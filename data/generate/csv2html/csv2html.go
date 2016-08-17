package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
	// "path"
	// "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	file_in := args[0]
	// path := path.Base(file_in)

	// file_chunks := strings.Split(path, ".")
	// file_out := file_chunks[0] + ".html"
	// fmt.Println(file_chunks)

	f, err := os.Open(file_in)
	check(err)
	r := csv.NewReader(bufio.NewReader(f))
	row_num := 1
	working := "<table>\n"
	for {
		record, err := r.Read()
		if row_num == 1 {
			working += "<th>"
			for _, v := range record {
				working += "<td>" + v + "</td>\n"
			}
			working += "</th>"
		} else {
			working += "<tr>"
			for _, v := range record {
				working += "<td>" + v + "</td>\n"
			}
			working += "</tr>"
		}
		row_num += 1

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(record)
	}
	working += "</table>\n"
	fmt.Println(working)
}
