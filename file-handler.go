package main

import (
    "fmt"
    "bufio"
    "os"
    "io/ioutil"
)

func writedata(data []byte, name string) {
	fo, _ := os.Create(name) 
	fo.Write(data)
	fmt.Println("done")
}

func readdata(name string) []byte {
	fi, _ := os.Open(name) 
	read := bufio.NewReader(fi)
	data, _ := ioutil.ReadAll(read) 
	return data
}

func main() {
	writedata(readdata("text1"),"text2")
}
