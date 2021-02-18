package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
    "bufio"
)


func getKey() string {
	keyreader := bufio.NewReader(os.Stdin)
	fmt.Println("key:")
	key, _ := keyreader.ReadString('\n')
	
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}


func encrypt(data []byte) []byte {
	block, _ := aes.NewCipher([]byte(getKey()))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte) []byte {
	block, err := aes.NewCipher([]byte(getKey()))
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return plaintext
}

func main() {
	
	fi, _ := os.Open(os.Args[2]) 
	read := bufio.NewReader(fi)
	data, _ := ioutil.ReadAll(read) 
	
	switch os.Args[1] {
	case "-enc":
		fo, _ := os.Create(os.Args[2]+".enc") 
		fo.Write(encrypt(data))
		fmt.Println("done")
	case "-dec":
		fo, _ := os.Create(os.Args[2]+".dec") 
		fo.Write(decrypt(data))
		fmt.Println("done")
	default:
		fmt.Println("o no no no !!!")
	}	
}


