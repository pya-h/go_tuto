package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHashCode(filename string) (uint32, error) {
	fbs, err := ioutil.ReadFile(filename)  // read file into bytes
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(fbs)
	return h.Sum32(), nil
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	} ()

	/*hsh := crc32.NewIEEE()
	hsh.Write([]byte("test"))
	sum := hsh.Sum32()
	fmt.Println(sum)*/
	c1, err := getHashCode("non-crypt-hash.go")
	if err != nil {
		panic(err)
	}
	c2, err2 := getHashCode("crypted-hash.go")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(c1, c2, c1 == c2)
}
