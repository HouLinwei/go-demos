package main

import "fmt"
import (
	"gopkg.in/redis.v4"
	"reflect"
	"encoding/json"
)
func main()  {
	fmt.Println("test redis command.")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	key := "test-set2"

	// SADD
	vids := []string{"int","float","uint","string",}
	for _ , v := range vids{
		err = client.SAdd(key, v).Err()
		if err != nil{
			fmt.Println(err)
		}
	}

	// SMEMBER
	res, err := client.SMembers(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(res))

	p := Person{Name:"testdwdwdw",Age:113,}
	err = client.SAdd("test-set4", p).Err()
	if err != nil {
		fmt.Println(err)
	}

	p2, err := client.SMembers("test-set4").Result()
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(p2)
		for _, raw := range p2{
			var tmpP Person
			json.Unmarshal([]byte(raw), &tmpP)
			fmt.Println(tmpP.Name)
			fmt.Println(tmpP.Age)
		}
	}
}

type Person struct {
	Name string
	Age int
}

func (p Person)MarshalBinary()([]byte, error){
	return json.Marshal(p)
	//data, _:= json.Marshal(p)
	//var buf bytes.Buffer
	//gzipWriter := gzip.NewWriter(&buf)
	//
	//idx, err := gzipWriter.Write(data)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(idx)
	//
	//if err := gzipWriter.Close(); err != nil {
	//	log.Fatal(err)
	//}
	//return buf.Bytes(), nil
	//return []byte(fmt.Sprintf("%s%d",p.Name,p.Age)), nil
}

func (p Person)UnmarshalBinary(data []byte)error{
	fmt.Println(data)
	p.Name = "res"
	p.Age =  11
	return  nil
}