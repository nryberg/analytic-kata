package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
	entry := Entry{}
	dat, err := ioutil.ReadFile("./data_template.yaml")
	if err != nil {
		log.Println("File Read error: ", err)
	}
	err = yaml.Unmarshal(dat, &entry)
	if err != nil {
		log.Println("File Read error: ", err)
	}
	fmt.Println(entry.Name)

	fruit := Fruit{}
	fruit.Name = "Fred"
	fruit.Weight = 32
	fruit.Measure = "Weight"

	entries := []Entry{}
	
	entries = append(entries, 
	entry.Fruit = fruit

	log.Println("Entry:", entry)
	d, err := yaml.Marshal(&entry)
	if err != nil {
		log.Println("Error: ", err)
	}

	fmt.Println(string(d))
}
