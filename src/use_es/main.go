package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	c, err := elastic.NewClient(elastic.SetURL("http://10.99.184.74:9200"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
