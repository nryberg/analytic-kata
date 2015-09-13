package main

import (
	"bufio"
	//"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	max := 5
	rand.Seed(int64(time.Now().Second()))
	path := "/Users/Nick/Develop/analytic-kata/toolsets/samples/threeletternames.txt"
	data, err := readLines(path)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < max; i++ {
		log.Println(data[rand.Intn(len(data))])
	}

}
