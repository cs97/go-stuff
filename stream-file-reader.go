package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: "+os.Args[0]+" <file> <count>")
		os.Exit(3)
	}
	
	i, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	streamread(os.Args[1], i)
}



func streamread(s string, c int) {
	f, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	
	b1 := make([]byte, c)
	
	n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	
	n1, err = f.Read(b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))		
}
