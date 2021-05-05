package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: "+os.Args[0]+" <count>")
		os.Exit(3)
	}
	
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	streamread(i)
}



func streamread(c int) {
	f, err := os.Open("text1")
	if err != nil {
		panic(err)
	}
	
	b1 := make([]byte, c)
	
	n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}
