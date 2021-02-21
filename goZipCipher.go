package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"bufio"
	"os"
	"io"
	"io/ioutil"
	"compress/gzip"
	"bytes"
    
)

func writedata(data []byte, name string) {
	fo, _ := os.Create(name) 
	fo.Write(data)
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

func getKey() string {
	keyreader := bufio.NewReader(os.Stdin)
	fmt.Println("key:")
	key, _ := keyreader.ReadString('\n')
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, keyhash string) []byte {
	block, _ := aes.NewCipher([]byte(keyhash))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		//fmt.Println(err)
		panic(err)
	}
	dataout := gcm.Seal(nonce, nonce, data, nil)
	return dataout
}

func decrypt(data []byte, keyhash string) []byte {
	block, err := aes.NewCipher([]byte(keyhash))
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	dataout, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	return dataout
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage: "+os.Args[0]+" [-enc|-dec] <file>")
		os.Exit(3)
	}
	
	switch os.Args[1] {
	case "-enc":
		writedata(encrypt(zip(readdata(os.Args[2])), getKey()),os.Args[2]+".enc")
		fmt.Println("done")
	case "-dec":
		writedata(unzip(decrypt(readdata(os.Args[2]), getKey())),os.Args[2]+".dec")
		fmt.Println("done")
	default:
		fmt.Println("usage: "+os.Args[0]+" [-enc|-dec] <file>")
	}	
}


