package main

import (
    "fmt"
    "bufio"
    "os"
    "io/ioutil"
    "compress/gzip"
    "bytes"
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

func zip(data []byte) []byte {
	var b bytes.Buffer
    w := gzip.NewWriter(&b)
    w.Write([]byte(data))
    w.Close()
    return b.Bytes()
}

func unzip(data []byte) []byte {
	buf := bytes.NewBuffer(data)
    reader, _ := gzip.NewReader(buf)
	out, err := ioutil.ReadAll(reader)
	if err != nil {
        panic(err)
    }
    return out
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: "+os.Args[0]+" [-zip|-unzip] <file>")
		os.Exit(3)
	}
	switch os.Args[1] {
	case "-zip":
		writedata(zip(readdata(os.Args[2])),os.Args[2]+".zip")
		fmt.Println("done")
	case "-unzip":
		writedata(unzip(readdata(os.Args[2])),os.Args[2]+".unzip")
		fmt.Println("done")
	default:
		fmt.Println("usage: "+os.Args[0]+" [-zip|-unzip] <file>")
	}	
}

