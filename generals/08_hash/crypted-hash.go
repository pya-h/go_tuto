package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	} ()

	nh := crc32.NewIEEE()
	nh.Write([]byte("this is a test"))
	_32bit_hash := nh.Sum32() // 32 bit hash can be represented as an uint
	// ******** these hash codes are nearly impossible to recover or reverse ****** ////
	ch := sha1.New()
	ch.Write([]byte("this is a test"))
	_160_bit_hash := ch.Sum([]byte{}) // but there is no 160-bit native type 
	fmt.Println("Crypted hash by SHA-1 is: ", _160_bit_hash, ", and non-crypted hash by CRC32 is: ", _32bit_hash)
}
