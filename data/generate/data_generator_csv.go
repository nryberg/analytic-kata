package main

import (
	//"fmt"
	// "gopkg.in/yaml.v2"
	"log"
	"encoding/csv"
	"os"
)

type Entry struct {
	Fruit
}
type Fruit struct {
	Name        string
	Weight      float32
	Measure     string
	Price       float32
	GrossWeight float32 `yaml:"gross_weight"`
}

func main() {
	//entries := []Fruit{}
	f, err := os.Open("./base_fruit.csv")

	//	dat, err := ioutil.ReadFile("./base_fruit.csv")
	if err != nil {
		log.Println("File Read error: ", err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	 if err != nil {
		 	log.Println("CSV Read Error: ", err)
	 }
	log.Println("Entry count: ", len(lines))

}
