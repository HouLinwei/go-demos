package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func Compress() {
	// 短文本的时候并没有什么优势。
	fmt.Println("Compress Data.")
	rawMsg := "raw msg"
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	gzipWriter.Name = "test gzip compression"
	gzipWriter.ModTime = time.Now()

	idx, err := gzipWriter.Write([]byte(rawMsg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(idx)

	if err := gzipWriter.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Before:", len([]byte(rawMsg)))
	fmt.Println("After:", len(buf.Bytes()))

	// write data to file.
	err = ioutil.WriteFile("/tmp/gzip.txt", buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func DeCompress() {
	fmt.Println("Decompress Data.")
	data, err := ioutil.ReadFile("/tmp/gzip.txt")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gzipReader.Name)
	fmt.Println(gzipReader.ModTime)

	data, err = ioutil.ReadAll(gzipReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

}

// How to use gzip with Golang?
func main() {
	fmt.Print("Gzip Demo.\n")
	Compress()
	DeCompress()
}
