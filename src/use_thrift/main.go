package main

import (
	"fmt"
	"use_thrift/gen-go/search"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	host = "10.108.35.17"
	port = 8991
)

func main()  {
	fmt.Println("=========")

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Println("Addr: ", addr)
	// create thrift client

}

func getClient() (search.FastVideoSearchServiceClient, error){
	//var transport  *thrift.T
	//client, err := search.NewFastVideoSearchServiceClientFactory()
	return nil, nil
}