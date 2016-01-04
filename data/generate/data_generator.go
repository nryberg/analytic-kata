package main

import (
	"fmt"
	_ "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Fruit struct {
}

func main() {
	dat, err := ioutil.ReadFile("./data_template.yaml")
	if err != nil {
		log.Println("File Read error: ", err)
	}
	fmt.Print(string(dat))
}
