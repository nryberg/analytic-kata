package main

import (
	//"fmt"
	// "gopkg.in/yaml.v2"
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v1"
)

type Config struct {
	facts struct {
		description  string
		name         string
		dateScale    string
		quantityLow  int
		quantityHigh int
	}
	dimensions []string
}

type entry struct {
	fruit
}

type fruit struct {
	Name        string
	Weight      float32
	Measure     string
	Price       float32
	GrossWeight float32 `yaml:"gross_weight"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("./Apples_and_Oranges/fruit.csv")
	if err != nil {
		panic(err)
	}

	var config Config

	configFile, err := ioutil.ReadFile("./Apples_and_Oranges/config.yaml")
	check(err)

	err = yaml.Unmarshal(configFile, &config)
	check(err)

	//	dat, err := ioutil.ReadFile("./base_fruit.csv")

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println("CSV Read Error: ", err)
	}
	log.Println("Entry count: ", len(lines))
	// fmt.Printf(lines[0].Name)
}
